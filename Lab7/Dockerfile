FROM golang:1.15-alpine AS build

WORKDIR /src/
COPY webserver.go /src/
RUN CGO_ENABLED=0 go build -o /bin/webserver

FROM golang:1.19-alpine AS build

WORKDIR /src/
COPY server.go /src/
COPY go.mod /src/
COPY go.sum /src/
RUN go mod download
RUN go build -o /src/server .

COPY --from=build /src/server /usr/local/bin .
EXPOSE 50051
CMD ["/usr/local/bin/server"]