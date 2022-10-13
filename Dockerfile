
# syntax=docker/dockerfile:1
FROM  golang:1.18 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN true
COPY *.go ./
COPY radexone/ ./radexone
COPY cmd/ ./cmd

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /goradex

## Deploy
FROM scratch
COPY --from=build /goradex /goradex

EXPOSE 9090

ENTRYPOINT ["/goradex"]
