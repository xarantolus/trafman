package query

import (
	"database/sql"
	"time"
)

func ClonesChart(db *sql.DB, repoID int) (c TimeSeriesChart, err error) {
	var (
		chartLabelDates []Date
		cloneCounts     []any
		cloneUniques    []any
	)

	cloneRows, err := db.Query(`select date, count, uniques
												from repotrafficclones
												where repo_id=$1
												ORDER BY date asc`, repoID)
	if err != nil {
		return
	}
	defer cloneRows.Close()

	for cloneRows.Next() {
		var (
			date           time.Time
			count, uniques int
		)
		err = cloneRows.Scan(&date, &count, &uniques)
		if err != nil {
			return
		}

		chartLabelDates = append(chartLabelDates, Date(date))
		cloneCounts = append(cloneCounts, count)
		cloneUniques = append(cloneUniques, uniques)
	}
	if err = cloneRows.Err(); err != nil {
		return
	}

	return TimeSeriesChart{
		Labels: chartLabelDates,
		Datasets: []Dataset{
			{
				Label:           "Clones",
				BackgroundColor: "#1F6FEB",
				BorderColor:     "#1F6FEB",
				Data:            cloneCounts,
			},
			{
				Label:           "Unique",
				BackgroundColor: "#238636",
				BorderColor:     "#238636",
				Data:            cloneUniques,
			},
		},
	}, nil
}
