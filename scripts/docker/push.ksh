# set the definitions
INSTANCE=deposit-auth-ws
NAMESPACE=uvadave

# push the current image
docker push $NAMESPACE/$INSTANCE

# return status
exit $?
