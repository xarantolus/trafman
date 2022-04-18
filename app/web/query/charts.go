package query

import "time"

type Dataset struct {
	Label           string `json:"label"`
	BackgroundColor string `json:"backgroundColor"`
	BorderColor     string `json:"borderColor"`
	Fill            bool   `json:"fill"`
	Data            []any  `json:"data"`
}

type TimeSeriesChart struct {
	Labels   []Date    `json:"labels"`
	Datasets []Dataset `json:"datasets"`
}

type Date time.Time

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte("\"" + time.Time(d).Format("2006-01-02") + "\""), nil
}
