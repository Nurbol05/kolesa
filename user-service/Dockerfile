FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY shared ./shared
COPY user-service ./user-service

COPY .env .env

RUN cd user-service && \
    go build -o /app/user-service-app main.go && \
    chmod +x /app/user-service-app

EXPOSE 8081

CMD ["/app/user-service-app"]