FROM golang:1.26

# Устанавливаем gcc и библиотеки для работы CGO и SQLite
RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc \
    libc6-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Сначала копируем только файлы зависимостей для кэширования
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальной исходный код
COPY . .

# Открываем нужный порт
EXPOSE 7540

# Собираем приложение С ВКЛЮЧЕННЫМ CGO
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /go-final-project

CMD ["/go-final-project"]
