# Clickstream Analytics System

## Table of Contents
1. [Overview](#overview)
2. [Features](#features)
3. [Tech Stack](#tech-stack)
4. [Getting Started](#getting-started)
5. [Usage](#usage)
6. [Monitoring](#monitoring)
7. [GraphQL API](#graphql-api)
8. [Scaling](#scaling)
9. [Next Steps](#next-steps)
10. [Contributing](#contributing)
11. [License](#license)

## Overview
Clickstream Analytics System is a high-throughput, real-time analytics platform that tracks user interactions on a website, such as clicks, page views, and scrolls. It uses **Golang** for backend processing and **Apache Cassandra** as a distributed NoSQL database for efficient time-series data storage. The system also integrates **Kafka** for real-time event streaming, **Prometheus & Grafana** for monitoring, and **GraphQL** for advanced querying.

## Features
- **Real-time event tracking** (clicks, page views, etc.)
- **Kafka-based event ingestion** for scalable processing
- **Cassandra DB for high-write throughput**
- **Prometheus & Grafana for monitoring API and Kafka metrics**
- **GraphQL API for flexible querying**
- **Horizontal scaling with Cassandra multi-node setup**
- **Single `docker-compose.yml` file integrating all services**

## Tech Stack
- **Backend:** Golang
- **Database:** Apache Cassandra
- **Streaming:** Apache Kafka
- **Monitoring:** Prometheus & Grafana
- **API:** REST & GraphQL
- **Deployment:** Docker & Kubernetes

---

## Getting Started
### Prerequisites
Ensure you have the following installed:
- **Docker & Docker Compose**
- **Go 1.22+**
- **Cassandra 4.0+**
- **Kafka 3.0+**
- **Prometheus & Grafana**

### Installation
#### **Clone Repository**
```sh
git clone https://github.com/yourname/clickstream-analytics.git
cd clickstream-analytics
```

#### **Start All Services with Docker Compose**
```sh
docker-compose up -d
```

#### **Run Backend Service (if not using Docker Compose)**
```sh
go run cmd/main.go
```

## Usage
### **Track Event (POST)**
```sh
curl -X POST "http://localhost:8080/track" -H "Content-Type: application/json" -d '{
    "user_id": "550e8400-e29b-41d4-a716-446655440000",
    "page_url": "/home",
    "event_type": "click",
    "user_agent": "Mozilla/5.0",
    "ip_address": "192.168.1.1"
}'
```

### **Get Page Views (GET)**
```sh
curl -X GET "http://localhost:8080/analytics/page-views/home"
```

## GraphQL API
### **Query analytics using GraphQL**
```sh
curl -X POST "http://localhost:8080/graphql" -H "Content-Type: application/json" -d '{ "query": "{ pageViews(pageURL: \"/home\") }" }'
```

## Next Steps
- 🚀 **Deploy to Kubernetes with Helm**
- 🌎 **Add WebSockets for real-time dashboards**
- 📊 **Enhance analytics with Spark or Flink**

## Contributing
PRs are welcome! Fork the repo and submit a pull request.

## License
This project is licensed under the Apache-2.0 License.

