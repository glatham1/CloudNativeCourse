
FROM golang:1.17.5-alpine AS build

WORKDIR /Lab7/
COPY go.mod .
COPY go.sum .
RUN go mod download
 
COPY ./movieapi/ ./movieapi
COPY movieserver/server.go /Lab7
RUN CGO_ENABLED=0 go build -o /bin/movieserver

FROM scratch
COPY --from=build /bin/movieserver /bin/movieserver
ENTRYPOINT ["/bin/movieserver"]
