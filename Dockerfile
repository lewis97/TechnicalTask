FROM golang:alpine3.20 AS base

FROM base AS dev

RUN set -eux; \
	apk add -U --no-cache \
		make \
        curl \
		jq \
        postgresql \
		git \
	;

RUN go install github.com/vektra/mockery/v2@v2.43.2

WORKDIR /transactionServer

COPY . .

RUN go get -d -v ./...

