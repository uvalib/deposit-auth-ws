# set the definitions
INSTANCE=deposit-auth-ws
NAMESPACE=uvadave

# define as appropriate if you do not want the defaults
# export DATABASE_URL=
# export DATABASE_USER=
# export DATABASE_PASSWORD=

# get the IP address of the docker engine
host_ip=$(ifconfig eth0 2>/dev/null | grep "inet addr" | awk -F: '{print $2}' | awk '{print $1}')

# stop the running instance
docker stop $INSTANCE

# remove the instance
docker rm $INSTANCE

# remove the previously tagged version
docker rmi $NAMESPACE/$INSTANCE:current  

# tag the latest as the current
docker tag -f $NAMESPACE/$INSTANCE:latest $NAMESPACE/$INSTANCE:current

# run the instance passing the environment if appropriate
if [ -n "$DATABASE_URL" -a -n "$DATABASE_USER" -a -n "$DATABASE_PASSWORD" ];
   docker run -d -p $host_ip:8080:8080 -p $host_ip:8081:8081 --name $INSTANCE -e DATABASE_URL=$DATABASE_URL -e DATABASE_USER=$DATABASE_USER -e DATABASE_PASSWORD=$DATABASE_PASSWORD $NAMESPACE/$INSTANCE:latest
else
   docker run -d -p $host_ip:8080:8080 -p $host_ip:8081:8081 --name $INSTANCE $NAMESPACE/$INSTANCE:latest
fi

# all good
exit 0
