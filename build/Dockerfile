#
# swis-api / Dockerfile
#

#
# stage 0 -- build image executable, combine codes and resource links
#

# https://hub.docker.com/_/golang
ARG ALPINE_VERSION GOLANG_VERSION
FROM golang:${GOLANG_VERSION}-alpine AS swapi-build

ARG APP_NAME
ENV APP_NAME ${APP_NAME}

RUN --mount=type=cache,target=/var/cache/apk \
	apk add git

WORKDIR /go/src/${APP_NAME}
COPY ./go.mod .

ARG GOMODCACHE GOCACHE
RUN --mount=type=cache,target="$GOMODCACHE" go mod download

COPY cmd/swis-api/ /go/src/swis-api/
COPY . .

# force use locked version
ARG APP_VERSION GOARCH GOOS
#RUN go install go.savla.dev/swis/v5@v${APP_VERSION}
RUN --mount=type=cache,target="$GOMODCACHE" \
	--mount=type=cache,target="$GOCACHE" \
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -o /go/bin/swis


#
# stage 1 -- release, combine build image and config(s), pre-runtime
#

FROM alpine:${ALPINE_VERSION} AS swapi-release

ARG APP_NAME APP_ROOT DOCKER_INTERNAL_PORT DOCKER_USER

ENV APP_NAME ${APP_NAME}

RUN apk add tzdata
RUN adduser -D -h ${APP_ROOT} -s /bin/sh ${DOCKER_USER}

# compose release image
COPY --from=swapi-build /go/bin/swis /bin/${APP_NAME}
COPY ./.github/assets/favicon.ico ${APP_ROOT}/

WORKDIR ${APP_ROOT}
EXPOSE ${DOCKER_INTERNAL_PORT}
USER ${DOCKER_USER}
ENTRYPOINT ["/bin/swis-api"]

