FROM golang:1.22.5

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd

# Кафка не успевает запуститься, не смотря на depends, 
# поэтому app должен запускаться с задержкой
# Придумаю другое решение - исправлю
CMD ["sh", "-c", "sleep 10 && ./main"]