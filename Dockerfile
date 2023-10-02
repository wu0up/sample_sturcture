FROM golang:1.20-buster

COPY app /src
WORKDIR /src

RUN go mod download
RUN go build -o /go/main

#RUN go run /src/app/pkg/migrate/migrate.go
#WORKDIR /src/app/pkg/migrate/migrations
#RUN go run goose.go

#FROM debian
#FROM gcr.io/distroless/base-debian10
#WORKDIR /go/
#COPY --from=build /go/main .
#COPY --from=build /src/.env .
# COPY images/ ./images/

EXPOSE 8000
CMD ["/go/main"]
