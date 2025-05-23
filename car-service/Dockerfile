FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY shared ./shared
COPY car-service ./car-service

COPY .env .env

RUN cd car-service && \
    go build -o /app/car-service-app main.go && \
    chmod +x /app/car-service-app

EXPOSE 8082

CMD ["/app/car-service-app"]