FROM golang:alpine3.20 AS base

FROM base AS dev

RUN set -eux; \
	apk add -U --no-cache \
		make \
        curl \
		jq \
        postgresql \
	;

WORKDIR /transactionServer

COPY . .

RUN go get -d -v ./...
