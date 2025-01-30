package db

import (
	"log"
	"time"

	"github.com/gocql/gocql"
)

func InitCassandra() *gocql.Session {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "clickstream"
	cluster.Consistency = gocql.Quorum
	cluster.ConnectTimeout = 10 * time.Second

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to Cassandra: %v", err)
	}

	log.Println("Cassandra connected successfully")
	return session
}
