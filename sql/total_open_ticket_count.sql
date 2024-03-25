SELECT countMerge(count) AS total_count
FROM analytics.total_open_ticket_count
WHERE guild_id = ?;