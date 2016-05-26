FROM centos:7

# Create the run user and group
RUN groupadd -r webservice && useradd -r -g webservice webservice

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
COPY scripts/entry.sh $APP_HOME/scripts/entry.sh
COPY scripts/new_import.ksh $APP_HOME/scripts/new_import.ksh
COPY scripts/sis_status.ksh $APP_HOME/scripts/sis_status.ksh
COPY data/sample_from_sis.txt $APP_HOME/data/sample_from_sis.txt
COPY bin/deposit-auth-ws.linux $APP_HOME/bin/deposit-auth-ws
