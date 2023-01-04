ARG GOLANG=1.19-alpine

# Build stage
FROM golang:${GOLANG} AS BUILD-STAGE
RUN apk add --no-cache curl git
WORKDIR /app
COPY . /app
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o web_crawler ./cmd/web_crawler/


# Final stage
FROM --platform=linux/amd64 alpine:3.16.2

WORKDIR /app
COPY --from=BUILD-STAGE /app/web_crawler .
COPY --from=BUILD-STAGE /app/wait-for.sh .

RUN apk update && \
    apk add --no-cache ca-certificates && \
    apk add --no-cache curl && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/* && \
    chmod +x /app/wait-for.sh

EXPOSE 8080
CMD ["./web_crawler"]

