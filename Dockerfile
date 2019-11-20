FROM golang:alpine AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build cmd/server/server.go

FROM scratch
COPY --from=builder /app/server /app/server
ENV PORT 5001
EXPOSE 5001
ENTRYPOINT [ "/app/server" ]