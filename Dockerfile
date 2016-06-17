FROM alpine:3.4

# Add bash cos we dont get by default
RUN apk --update add bash

# Create the run user and group
RUN addgroup webservice && adduser webservice -G webservice -D

# Specify home 
ENV APP_HOME /depositauth-ws
WORKDIR $APP_HOME

# Create necessary directories
RUN mkdir -p $APP_HOME/scripts $APP_HOME/bin $APP_DIR/data
RUN chown -R webservice $APP_HOME && chgrp -R webservice $APP_HOME

# Specify the user
USER webservice

# port and run command
EXPOSE 8080
CMD scripts/entry.sh

# Move in necessary helper scripts and binary
COPY scripts/entry.sh $APP_HOME/scripts/
COPY scripts/*.ksh $APP_HOME/scripts/
COPY data/sample_from_sis.txt $APP_HOME/data/
COPY bin/deposit-auth-ws.linux $APP_HOME/bin/deposit-auth-ws
