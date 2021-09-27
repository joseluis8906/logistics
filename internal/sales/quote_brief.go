package sales

type (
	QuoteBrief struct {
		Distance Currency
		Weight   Currency
		Size     Currency
	}
)

func (q QuoteBrief) Total() Currency {
	return q.Distance + q.Weight + q.Size
}
