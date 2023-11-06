###########
# builder #
###########

FROM golang:1.21-bullseye AS builder
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    upx-ucl

WORKDIR /build
COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 go build -o ./bin/portpoc \
    -ldflags='-w -s -extldflags "-static"' \
    . \
 && upx-ucl --best --ultra-brute ./bin/portpoc

###########
# release #
###########

FROM gcr.io/distroless/static-debian11:latest AS release

COPY --from=builder /build/bin/portpoc /bin/
WORKDIR /workdir
ENTRYPOINT ["/bin/portpoc"]
