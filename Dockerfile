FROM golang:alpine3.20 AS base

# ------------------------------------------------------------------------------
# Development image
# ------------------------------------------------------------------------------

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

# ------------------------------------------------------------------------------
# Application image
# ------------------------------------------------------------------------------

FROM base AS app

WORKDIR /transactionServer

COPY . .

RUN go mod download
RUN go build -o dist/ ./cmd/...

CMD ["/dist/transactionServer"]
