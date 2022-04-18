package query

import (
	"database/sql"
	"time"
)

func ViewsChart(db *sql.DB, repoID int) (c TimeSeriesChart, err error) {
	var (
		chartLabelDates []Date
		viewCounts      []any
		uniqueCounts    []any
	)

	viewRows, err := db.Query(`select date, count, uniques
												from repotrafficviews
												where repo_id=$1
												ORDER BY date asc`, repoID)
	if err != nil {
		return
	}
	defer viewRows.Close()

	for viewRows.Next() {
		var (
			date           time.Time
			count, uniques int
		)
		err = viewRows.Scan(&date, &count, &uniques)
		if err != nil {
			return
		}

		chartLabelDates = append(chartLabelDates, Date(date))
		viewCounts = append(viewCounts, count)
		uniqueCounts = append(uniqueCounts, uniques)
	}
	if err = viewRows.Err(); err != nil {
		return
	}

	return TimeSeriesChart{
		Labels: chartLabelDates,
		Datasets: []Dataset{
			{
				Label:           "Views",
				Data:            viewCounts,
				BackgroundColor: "#238636",
				BorderColor:     "#238636",
			},
			{
				Label:           "Unique visitors",
				BackgroundColor: "#1F6FEB",
				BorderColor:     "#1F6FEB",
				Data:            uniqueCounts,
			},
		},
	}, nil
}
