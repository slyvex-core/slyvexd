# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/slyvex-core/slyvexd

WORKDIR /go/src/github.com/slyvex-core/slyvexd

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /go/src/github.com/slyvex-core/slyvexd/cmd/slyvexctl

RUN GOOS=linux go build -a -installsuffix cgo -o slyvexctl .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/slyvex-core/slyvexd/cmd/slyvexctl/slyvexctl /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
