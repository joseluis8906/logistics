package sales

type (
	Dimension struct {
		Width  float64
		Height float64
		Length float64
	}
)

func (d Dimension) Volume() float64 {
	return d.Height * d.Width * d.Length
}
