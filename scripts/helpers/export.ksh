#
#
#

# ensure we have and endpoint
if [ -z "$DEPOSITAUTH_URL" ]; then
   echo "ERROR: DEPOSITAUTH_URL is not defined"
   exit 1
fi

# and an API token
if [ -z "$API_TOKEN" ]; then
   echo "ERROR: API_TOKEN is not defined"
   exit 1
fi

# issue the command
curl -X POST $DEPOSITAUTH_URL/export?auth=$API_TOKEN

echo "done"
exit 0

#
# end of file
#
