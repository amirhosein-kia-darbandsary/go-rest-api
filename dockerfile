FROM golang:1.22.4 AS builder

WORKDIR /app


COPY go.mod /app/

# RUN  go mod download


COPY . /app/

RUN go build -o capp /app/cmd/server/main.go


FROM alpine:latest AS production

COPY --from=builder /app .

CMD [ "./capp" ]