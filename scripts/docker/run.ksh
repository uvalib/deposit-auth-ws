# set the definitions
INSTANCE=deposit-auth-ws
NAMESPACE=uvadave

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

# run the instance
docker run -d -p $host_ip:8080:8080 -p $host_ip:8081:8081 --name $INSTANCE $NAMESPACE/$INSTANCE:latest

# all good
exit 0
