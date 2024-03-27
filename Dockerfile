# syntax=docker/dockerfile:1
FROM golang:1.22.1 as build
WORKDIR /src
COPY main.go .
RUN go build -o /bin/hello ./main.go

FROM scratch
COPY --from=build /bin/hello /bin/hello
CMD [ "/bin/hello" ]