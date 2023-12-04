FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/dl7850949/blob-storage
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/blob-storage /go/src/gitlab.com/dl7850949/blob-storage


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/blob-storage /usr/local/bin/blob-storage
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["blob-storage"]
