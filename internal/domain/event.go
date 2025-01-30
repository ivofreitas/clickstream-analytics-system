package domain

import (
	"time"

	"github.com/gocql/gocql"
)

type Event struct {
	EventID   gocql.UUID `json:"event_id"`
	UserID    gocql.UUID `json:"user_id"`
	PageURL   string     `json:"page_url"`
	EventType string     `json:"event_type"`
	EventTime time.Time  `json:"event_time"`
	UserAgent string     `json:"user_agent"`
	IPAddress string     `json:"ip_address"`
}
