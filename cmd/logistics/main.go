package main

import (
	"context"
	"errors"
	"fmt"
	"log/syslog"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"bitbucket.org/josecaceresatencora/logistics/internal/sales"
	"bitbucket.org/josecaceresatencora/logistics/pkg/bus"
	myhttp "bitbucket.org/josecaceresatencora/logistics/pkg/http"
	mykafka "bitbucket.org/josecaceresatencora/logistics/pkg/kafka"
	mylog "bitbucket.org/josecaceresatencora/logistics/pkg/log"
	"bitbucket.org/josecaceresatencora/logistics/pkg/logistics"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	programName := filepath.Base(os.Args[0])
	rand.Seed(time.Now().UnixNano())
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "dev"
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	//###########
	//### log ###
	//###########
	{
		sysLog, err := syslog.New(syslog.LOG_NOTICE|syslog.LOG_USER, programName)
		if err != nil {
			log.Fatal(err)
		}

		log.SetFormatter(mylog.CustomLogger{
			Env:       env,
			Formatter: log.StandardLogger().Formatter,
		})
		log.SetOutput(sysLog)
		log.SetReportCaller(true)
	}

	//#############
	//### viper ###
	//#############
	{
		viper.SetConfigName(env)
		viper.SetConfigType("json")
		viper.AddConfigPath("./configs")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}

		mylog.SetLogLevel(viper.GetString("log.level"))
	}

	wg := sync.WaitGroup{}

	//#############################
	//### etcd(runtime configs) ###
	//#############################
	{
		runtimeViper := viper.New()
		err := runtimeViper.AddRemoteProvider("etcd", fmt.Sprintf("http://%s:%d", viper.GetString("etcd.host"), viper.GetInt("etcd.port")), "/config/logistics.json")
		if err != nil {
			log.Fatal(err)
		}
		runtimeViper.SetConfigType("json")

		err = runtimeViper.ReadRemoteConfig()
		if err != nil {
			log.Fatal(err)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					log.Info("remote viper watcher gracefully shutdown!")
					return
				default:
					time.Sleep(time.Second * 5)
					err := runtimeViper.WatchRemoteConfig()
					if err != nil {
						log.Errorf("unable to read remote config: %v", err)
						continue
					}

					logLevel := runtimeViper.GetString("log.level")
					mylog.SetLogLevel(logLevel)
				}
			}
		}()
	}

	//#############
	//### kafka ###
	//#############
	{
		host := viper.GetString("kafka.host")
		topics := viper.GetStringSlice("kafka.groups.default.topics")
		consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": host,
			"group.id":          viper.GetString("kafka.groups.default.id"),
			"auto.offset.reset": "earliest",
		})
		if err != nil {
			log.Fatal(err)
		}

		err = consumer.SubscribeTopics(topics, nil)
		if err != nil {
			log.Fatal(err)
		}

		log.Debug("listening for events...")

		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					consumer.Close()
					log.Info("kafka consumer gracefully shutdown!")
					return
				default:
					msg, err := consumer.ReadMessage(time.Millisecond * 100)
					if err != nil && err.(kafka.Error).Code() != kafka.ErrTimedOut {
						log.Error(err)
					}

					if msg != nil {
						log.Debug("event arrive: ", string(msg.Value))
					}
				}
			}
		}()

		producer, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": host,
		})
		if err != nil {
			log.Fatal(err)
		}
		bus.SetOutput(&mykafka.Emiter{Producer: producer, Topic: viper.GetString("kafka.groups.default.topics.0")})

		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					producer.Close()
					log.Info("kafka producer events gracefully shutdown!")
					return
				case e := <-producer.Events():
					switch ev := e.(type) {
					case *kafka.Message:
						if ev.TopicPartition.Error != nil {
							log.Errorf("Delivery failed: %v\n", ev.TopicPartition)
							continue
						}

						log.Debugf("Delivered message to %v\n", ev.TopicPartition)
					}
				}
			}
		}()
	}

	//#############
	//### mongo ###
	//#############
	{
		user := viper.GetString("mongodb.logistics.user")
		passwd := viper.GetString("mongodb.logistics.passwd")
		host := viper.GetString("mongodb.logistics.host")
		port := viper.GetInt("mongodb.logistics.port")
		db := viper.GetString("mongodb.logistics.db")
		uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", user, passwd, host, port, db)
		_, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatal(err)
		}
	}

	//#######################
	//### chi http router ###
	//#######################
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/heath", myhttp.Health)
	r.Route("/sales", func(r chi.Router) {
		r.Get("/quote-shipping", func(w http.ResponseWriter, r *http.Request) { quoteShipping(ctx, w, r) })
	})
	r.Route("/transport", func(r chi.Router) {
		r.Get("/city-location", func(w http.ResponseWriter, r *http.Request) { cityLocation(ctx, w, r) })
	})

	//###################
	//### http server ###
	//###################
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("http.server.port")),
		Handler: r,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}

		log.Info("http server gracefully shutdown!")
	}()

	logistics.SetApp(logistics.Application{
		Commands: logistics.Commands{},
		Queries: logistics.Queries{
			QuoteShipping: sales.QuoteShipping,
		},
	})

	log.Infof("service is running on :%d port", viper.GetInt("http.server.port"))

	osCall := <-done
	log.Infof("system call: %+v, shutting down, please wait", osCall)

	err := httpServer.Shutdown(ctx)
	if err != nil {
		log.Error(err)
	}

	cancel()

	wg.Wait()
	log.Info("service shutdown completed")
}
