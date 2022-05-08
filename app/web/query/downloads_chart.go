package query

import (
	"database/sql"
	"time"
)

func DownloadsChart(db *sql.DB, repoID int) (c TimeSeriesChart, err error) {
	var (
		chartLabelDates = []Date{}
		downloadCounts  = []any{}
	)

	downloadRows, err := db.Query(`SELECT max(ras.date_time), sum(ras.download_count)
		FROM repositories r
		join releases rel on r.id = rel.repo_id
		join releaseassets ras on ras.release_id = rel.id
		WHERE r.id=$1
		GROUP BY DATE(ras.date_time)
		ORDER BY DATE(ras.date_time) asc`, repoID)
	if err != nil {
		return
	}
	defer downloadRows.Close()

	for downloadRows.Next() {
		var (
			date  time.Time
			count int
		)
		err = downloadRows.Scan(&date, &count)
		if err != nil {
			return
		}

		chartLabelDates = append(chartLabelDates, Date(date))
		downloadCounts = append(downloadCounts, count)
	}
	if err = downloadRows.Err(); err != nil {
		return
	}

	return TimeSeriesChart{
		Labels: chartLabelDates,
		Datasets: []Dataset{
			{
				Label:           "Downloads",
				Data:            downloadCounts,
				BackgroundColor: "#238636",
				BorderColor:     "#238636",
			},
		},
	}, nil
}
