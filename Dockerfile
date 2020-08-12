# build image 
FROM golang:1.14.5-alpine as builder

WORKDIR go/src/twittercli
ADD . /go/src/twittercli
RUN cd /go/src/twittercli \
        && mkdir -p go/bin \
        && GOMODULE111=on \
        && go install \
        && cd go/bin && ls

# final image
FROM alpine as final

RUN apk --no-cache add \
    bash \
    jq \
    ca-certificates

COPY --from=builder /go/bin/twittercli /bin/

#TODO: make this configurable
EXPOSE 8080

# copies from just the built artifact from the previous stage into this new stage
ENTRYPOINT ["/twittercli"]
