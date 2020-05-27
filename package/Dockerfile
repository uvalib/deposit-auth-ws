#
# build the target application
#
FROM golang:1.14 as builder
RUN apt-get -y update

WORKDIR /build
COPY go.mod go.sum Makefile ./
COPY depositauthws ./depositauthws
RUN make build-linux

#
# build the target container
#
FROM alpine:3.11

# update the packages
RUN apk update && apk upgrade && apk add bash tzdata ca-certificates curl && rm -fr /var/cache/apk/*

# Create the run user and group
RUN addgroup --gid 18570 sse && adduser --uid 1984 docker -G sse -D

# set the timezone appropriatly
ENV TZ=UTC
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Specify home 
ENV APP_HOME /depositauth-ws
WORKDIR $APP_HOME

# Create necessary directories
RUN mkdir -p $APP_HOME/scripts $APP_HOME/bin $APP_HOME/assets $APP_DIR/data
RUN chown -R docker $APP_HOME && chgrp -R sse $APP_HOME

# port and run command
EXPOSE 8080
CMD scripts/entry.sh

# Move in necessary helper scripts and binary
COPY data/container_bash_profile /home/docker/.profile
COPY scripts/entry.sh $APP_HOME/scripts/
COPY scripts/*.ksh $APP_HOME/scripts/
COPY assets/* $APP_HOME/assets/
COPY --from=builder /build/bin/deposit-auth-ws.linux $APP_HOME/bin/deposit-auth-ws

# Ensure permissions are correct
RUN chown docker:sse /home/docker/.profile $APP_HOME/scripts/entry.sh $APP_HOME/bin/deposit-auth-ws && chmod 755 /home/docker/.profile $APP_HOME/scripts/entry.sh $APP_HOME/bin/deposit-auth-ws

# Add the build tag
ARG BUILD_TAG
RUN test -n "$BUILD_TAG" && touch $APP_HOME/buildtag.build-$BUILD_TAG || touch $APP_HOME/buildtag.build-0

# Specify the user
USER docker

#
