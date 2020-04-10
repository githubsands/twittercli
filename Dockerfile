FROM golang:1.13-alpine as builder

# Install dependencies and build the binaries.
RUN apk add --no-cache --update alpine-sdk
# Start a new, final image.
FROM alpine as final

# Define a root volume for data persistence.
VOLUME /root/.twittercli

RUN apk --no-cache add \
    bash \
    ca-certificates \
&&  go install

# Copy the binaries from the builder image.
COPY --from=builder /go/bin/twittercli /bin/

EXPOSE 8000

# Specify the start command and entrypoint as the dcrlnd daemon.
ENTRYPOINT ["start_twittercli.sh"]
