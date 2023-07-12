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
COPY --from=firstbuildstage $APP_PATH/rowing-registration-api $APP_PATH
COPY --from=firstbuildstage $APP_PATH/data/translations/en.toml $APP_PATH/data/translations/en.toml

# Define entrypoint and port
ENV PORT=8397
CMD ["./rowing-registration-api"]
