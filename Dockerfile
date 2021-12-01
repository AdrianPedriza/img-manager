#  -- Base imagen where generates the bin executable file --
FROM golang:1.17-buster as builder

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/cmd/server ./...
## -- Production image --
FROM alpine:3 as prod_img

COPY  --from=builder /app/cmd/server/server /img-manager

EXPOSE 8080

ENTRYPOINT ["/img-manager"]
