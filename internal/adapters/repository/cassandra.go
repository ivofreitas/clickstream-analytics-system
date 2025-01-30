package repository

import (
	"github.com/gocql/gocql"
	"github.com/ivofreitas/clickstream-analytics-system/internal/app"
	"github.com/ivofreitas/clickstream-analytics-system/internal/domain"
	"time"
)

var day = time.Now().Truncate(24 * time.Hour)

type CassandraEventRepository struct {
	session *gocql.Session
}

func NewCassandraEventRepository(session *gocql.Session) app.EventRepository {
	return &CassandraEventRepository{session: session}
}

func (r *CassandraEventRepository) SaveEvent(event domain.Event) error {
	query := "INSERT INTO events (event_id, user_id, page_url, event_type, event_time, user_agent, ip_address) VALUES (?, ?, ?, ?, ?, ?, ?)"
	return r.session.Query(query, event.EventID, event.UserID, event.PageURL, event.EventType, event.EventTime, event.UserAgent, event.IPAddress).Exec()
}

func (r *CassandraEventRepository) GetPageViews(pageURL string) (int, error) {
	var viewCount int
	query := "SELECT view_count FROM page_views WHERE page_url = ? AND day = ?"
	err := r.session.Query(query, pageURL, day).Consistency(gocql.One).Scan(&viewCount)
	return viewCount, err
}

func (r *CassandraEventRepository) UpdateViewCount(pageURL string) error {
	updateQuery := "UPDATE page_views SET view_count = view_count + 1 WHERE page_url = ? AND day = ?"
	err := r.session.Query(updateQuery, pageURL, day).Exec()
	return err
}
