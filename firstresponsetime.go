package analytics

import (
	"context"
	"time"
)

func (c *Client) GetFirstResponseTimeStats(context context.Context, guildId uint64) (TripleWindow, error) {
	query := `
SELECT
    avgMerge(all_time),
    avgOrNullMerge(monthly),
    avgOrNullMerge(weekly)
FROM analytics.first_response_time_guild
WHERE guild_id = ?
GROUP BY guild_id`

	// Values in seconds
	var allTime int64
	var monthly, weekly *int64
	if err := c.client.QueryRow(context, query, guildId).Scan(&allTime, &monthly, &weekly); err != nil {
		return TripleWindow{}, err
	}

	return TripleWindow{
		AllTime: time.Duration(allTime) * time.Second,
		Monthly: mapNullableSecondsToDuration(monthly),
		Weekly:  mapNullableSecondsToDuration(weekly),
	}, nil
}
