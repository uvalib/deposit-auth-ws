# set the definitions
INSTANCE=deposit-auth-ws
NAMESPACE=uvadave

# move in the yaml file
cp src/main/resources/service.yaml target/

# build the image
docker build -t $NAMESPACE/$INSTANCE .

# all good
exit 0
