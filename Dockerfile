# Step 1: Modules caching
FROM golang:1.18beta2-alpine3.15 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.18beta2-alpine3.15 as builder
RUN apk add git
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./cmd/app

# Step 3: Final
FROM scratch
COPY --from=builder /app/configs /configs
COPY --from=builder /app/log /log
COPY --from=builder /bin/app /app
CMD ["/app"]