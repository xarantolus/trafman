package query

import (
	"database/sql"
	"time"
)

func RepoStatsChart(db *sql.DB, repoID int) (c TimeSeriesChart, err error) {
	statRows, err := db.Query(`SELECT max(date_time), stars, forks, watchers
						FROM repostats
						WHERE repo_id=$1
						GROUP BY DATE(date_time), stars, repo_id, forks, watchers
						order by DATE(date_time) asc`, repoID)
	if err != nil {
		return
	}
	defer statRows.Close()

	var (
		chartLabelDates = []Date{}
		chartStars      = []any{}
		chartForks      = []any{}
		chartWatchers   = []any{}
	)

	for statRows.Next() {
		var (
			date                   time.Time
			stars, forks, watchers int
		)
		err = statRows.Scan(&date, &stars, &forks, &watchers)
		if err != nil {
			return
		}

		chartLabelDates = append(chartLabelDates, Date(date))
		chartStars = append(chartStars, stars)
		chartForks = append(chartForks, forks)
		chartWatchers = append(chartWatchers, watchers)
	}
	if err = statRows.Err(); err != nil {
		return
	}
	return TimeSeriesChart{
		Labels: chartLabelDates,
		Datasets: []Dataset{
			{
				Label:           "Stars",
				Data:            chartStars,
				BackgroundColor: "#b86347",
				BorderColor:     "#b86347",
			},
			{
				Label:           "Forks",
				BackgroundColor: "#47B863",
				BorderColor:     "#47B863",
				Data:            chartForks,
			},
			{
				Label:           "Watchers",
				BackgroundColor: "#6347B8",
				BorderColor:     "#6347B8",
				Data:            chartWatchers,
			},
		},
	}, nil
}
