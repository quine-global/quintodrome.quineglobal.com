FROM debian:stable AS cssbuilder
WORKDIR /app

RUN set -x && apt-get update && \
  DEBIAN_FRONTEND=noninteractive apt-get install -y curl

COPY internal ./internal/
COPY sites ./sites/

FROM golang AS gobuilder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ARG TARGETARCH
ARG SITE=quintodrome
RUN GOOS=linux GOARCH=${TARGETARCH} go build -buildvcs=false -ldflags="-s -w" -o ./app ./sites/${SITE}/cmd

FROM debian:stable-slim AS runner
WORKDIR /app

RUN set -x && apt-get update && \
  DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates && \
  rm -rf /var/lib/apt/lists/*

COPY public ./public/
COPY --from=gobuilder /app/app ./

EXPOSE 8080

CMD ["./app"]
