# set the definitions
INSTANCE=deposit-auth-ws
NAMESPACE=uvadave

# pull the current runable image
docker pull $NAMESPACE/$INSTANCE:latest

# return status
exit $?
