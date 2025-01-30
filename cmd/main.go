package main

import (
	"github.com/ivofreitas/clickstream-analytics-system/internal/adapters/http"
	"github.com/ivofreitas/clickstream-analytics-system/internal/adapters/repository"
	"github.com/ivofreitas/clickstream-analytics-system/internal/app"
	"github.com/ivofreitas/clickstream-analytics-system/pkg/db"
	"github.com/julienschmidt/httprouter"
	"log"
	nethttp "net/http"
)

func main() {
	// Initialize Cassandra
	session := db.InitCassandra()
	defer session.Close()

	// Setup Repository
	repo := repository.NewCassandraEventRepository(session)

	// Setup Service
	service := app.NewEventService(repo)

	// Setup HTTP Handlers
	handler := http.NewEventHandler(service)

	// Setup Router
	router := httprouter.New()
	router.POST("/track", handler.TrackEvent)
	router.GET("/analytics/page-views/:page_url", handler.GetPageViews)

	log.Println("Server running on :8080")
	log.Fatal(nethttp.ListenAndServe(":8080", router))
}
