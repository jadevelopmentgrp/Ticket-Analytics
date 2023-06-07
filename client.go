package analytics

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2"
	"time"
)

type Client struct {
	client clickhouse.Conn
}

func NewClient(client clickhouse.Conn) *Client {
	return &Client{
		client,
	}
}

func Connect(address string, connections int, database, username, password string, readTimeout time.Duration) (*Client, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{address},
		Auth: clickhouse.Auth{
			Database: database,
			Username: username,
			Password: password,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		MaxOpenConns: connections,
		MaxIdleConns: connections,
		ReadTimeout:  readTimeout,
		ClientInfo: clickhouse.ClientInfo{
			Products: []struct {
				Name    string
				Version string
			}{
				{
					Name:    "tickets-analytics",
					Version: "0.1",
				},
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &Client{
		client: conn,
	}, nil
}