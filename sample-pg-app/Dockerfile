FROM golang:1.19 as build-stage

WORKDIR /build

COPY go.mod go.sum /build/
RUN go mod download

COPY main.go /build

RUN CGO_ENABLED=0 go build -o /build/sample-pg-app
RUN chmod 755 /build/sample-pg-app

FROM scratch AS release

WORKDIR /app

COPY --from=build-stage /build/sample-pg-app /app/sample-pg-app

CMD ["/app/sample-pg-app"]