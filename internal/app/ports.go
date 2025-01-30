package app

import "github.com/ivofreitas/clickstream-analytics-system/internal/domain"

type EventRepository interface {
	SaveEvent(event domain.Event) error
	GetPageViews(pageURL string) (int, error)
	UpdateViewCount(pageURL string) error
}
