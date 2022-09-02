FROM golang:1.18-alpine AS build
WORKDIR /app

RUN apk add --update make

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download
COPY ./src ./
COPY ./Makefile ./
RUN make build

FROM alpine AS release

WORKDIR /app
COPY --from=build /app/todo ./
COPY --from=build /app/migrations/ ./migrations/
COPY --from=build /app/config.yaml ./

EXPOSE 8080/tcp
ENTRYPOINT ["/app/todo"]
