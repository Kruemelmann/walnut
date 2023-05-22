FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /app-build .

FROM alpine
WORKDIR /
COPY --from=builder /app-build /app
EXPOSE 80
ENTRYPOINT ["/app"]
