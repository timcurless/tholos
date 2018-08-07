FROM golang:1.10.3-alpine3.8 AS build
LABEL maintainer=<tim.curless@thinkahead.com>

RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep

COPY . /go/src/tholos/
WORKDIR /go/src/meld/
RUN dep ensure --vendor-only && go build -o /bin/tholos

FROM alpine:3.8
COPY --from=build /bin/tholos /bin/tholos
ENTRYPOINT ["/bin/tholos"]
CMD ["--help"]
