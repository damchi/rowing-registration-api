FROM golang:1.20-alpine3.18 as firstbuildstage

RUN apk add --no-cache git

ENV APP_PATH="$GOPATH/src/rowing-registration-api"

RUN mkdir -p $APP_PATH
COPY . $APP_PATH
WORKDIR $APP_PATH

RUN go build

# Clean Image
FROM golang:1.20-alpine3.18

ENV APP_PATH="$GOPATH/src/rowing-registration-api"
RUN mkdir -p $APP_PATH

WORKDIR $APP_PATH
COPY --from=firstbuildstage $APP_PATH/rowing-club-api $APP_PATH

# Define entrypoint and port
ENV PORT=8397
CMD ["./rowing-registration-api"]
