# syntax=docker/dockerfile:1
FROM golang:1.22.0 as base
WORKDIR /src
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app main.go

FROM scratch
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /app /app
CMD [ "./app" ]