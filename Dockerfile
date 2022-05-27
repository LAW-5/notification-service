FROM golang:1.18 AS Production
WORKDIR /app
COPY go.mod .env ./
RUN go mod tidy
COPY . .
RUN go build -o notification-srv
EXPOSE 5010
CMD /app/notification-srv