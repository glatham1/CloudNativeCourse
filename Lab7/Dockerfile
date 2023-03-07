FROM golang:1.15-alpine AS build

WORKDIR /src/
COPY server.go /src/
RUN CGO_ENABLED=0 go build -o /bin/server

FROM scratch
COPY --from=build /bin/webserver /bin/server
ENTRYPOINT ["/bin/server"]