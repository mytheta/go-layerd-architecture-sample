FROM golang:1.12 AS build-env
ADD . /go/src/github.com/mytheta/go-layerd-architecture-sample/
WORKDIR /go/src/github.com/mytheta/go-layerd-architecture-sample
RUN make build

FROM alpine
RUN apk update && apk add ca-certificates
COPY --from=build-env /go/src/github.com/mytheta/go-layerd-architecture-sample/main /usr/local/bin/emole
COPY --from=build-env /go/src/github.com/mytheta/go-layerd-architecture-sample/interface/env/conf /usr/local/bin/interface/env/conf
EXPOSE 30000
CMD /usr/local/bin/emole -e /usr/local/bin/interface/env/conf