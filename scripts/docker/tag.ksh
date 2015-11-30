# set the definitions
INSTANCE=deposit-auth-ws
NAMESPACE=uvadave
TAG=$1

# did we provide a tag correctly...
if [ -n "$TAG" ]; then

   IMAGEID=$(docker images $NAMESPACE/$INSTANCE | grep latest | awk '{print $3 }')
   if [ -n "$IMAGEID" ]; then
      # tag the latest image approriatly
      docker tag $IMAGEID $NAMESPACE/$INSTANCE:$TAG
   fi
fi

# all good
exit 0
