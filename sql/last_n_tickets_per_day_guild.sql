SELECT date, countMerge(count) AS count
FROM analytics.tickets_per_day
WHERE guild_id = ?
GROUP BY date
ORDER BY date desc
WITH FILL STEP toIntervalDay(-1)
LIMIT ?;