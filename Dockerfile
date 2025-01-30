FROM golang:1.22

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o clickstream-api ./cmd/main.go

EXPOSE 8080
CMD ["./clickstream-api"]
