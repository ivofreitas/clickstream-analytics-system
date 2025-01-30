package app

import (
	"github.com/gocql/gocql"
	"github.com/ivofreitas/clickstream-analytics-system/internal/domain"
	"time"
)

type EventService struct {
	repo EventRepository
}

func NewEventService(repo EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) TrackEvent(event domain.Event) error {
	event.EventID, _ = gocql.RandomUUID()
	event.EventTime = time.Now()

	err := s.repo.SaveEvent(event)
	if err != nil {
		return err
	}

	return s.repo.UpdateViewCount(event.PageURL)
}

func (s *EventService) GetPageViews(pageURL string) (int, error) {
	return s.repo.GetPageViews(pageURL)
}
