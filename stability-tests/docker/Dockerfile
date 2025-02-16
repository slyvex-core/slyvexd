ARG SLYVEXD_IMAGE
ARG SLYVEXMINER_IMAGE

FROM ${SLYVEXD_IMAGE} as slyvexd
FROM ${SLYVEXMINER_IMAGE} as slyvexminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/slyvex-core/slyvexd

WORKDIR /go/src/github.com/slyvex-core/slyvexd

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=slyvexd /app/ /app/
COPY --from=slyvexminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/slyvex-core/slyvexd/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]