FROM golang:1.22.4 AS builder

WORKDIR /app


COPY go.mod go.sum /app/

RUN  go mod download


COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o App /app/cmd/server/main.go


FROM alpine:latest AS production

COPY --from=builder /app .


CMD [ "./App" ]
