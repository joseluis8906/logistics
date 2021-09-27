## Usage

**Setup all**

`$ make all`

**Run it**

`$ ./build/logistics`

**Check logs**

`$ tail -n25 -f /var/log/system.log | grep logistics`

**Use the endpoint**

`$ curl localhost:3000/sales/quote-shipping\?from_lat=0\&from_lng=0\&to_lat=1000\&to_lng=1\&weigth=7.5\&width=5\&height=5\&length=5`

**Run tests**

`$ go test ./... -cover`
