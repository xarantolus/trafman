package query

/*
SELECT repo_id, max(date_time), referrer, count, uniques
FROM repotrafficreferrers
WHERE repo_id=405108802
GROUP BY DATE(date_time), repo_id, referrer, count, uniques
order by DATE(date_time) asc
*/
