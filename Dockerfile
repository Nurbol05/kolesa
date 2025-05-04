# Go бейнесін қолданамыз
FROM golang:1.23

# Жұмыс папкасы
WORKDIR /app

# Модуль файлдарын көшіру
COPY go.mod go.sum ./
RUN go mod download

# .env файлын көшіру
COPY .env .env

# Қалған барлық файлдарды көшіру
COPY . .

# Құрастыру
RUN go build -o main .

# Порт көрсету
EXPOSE 8080

# Контейнерді іске қосу
CMD ["./main"]
