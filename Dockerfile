FROM golang:1.7

WORKDIR /app

COPY . .

RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y

CMD ["go", "run", "main.go"]