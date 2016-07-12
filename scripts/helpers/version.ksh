#
#
#

# ensure we have and endpoint
if [ -z "$DEPOSITAUTH_URL" ]; then
   echo "ERROR: DEPOSITAUTH_URL is not defined"
   exit 1
fi

# issue the command
echo "$DEPOSITAUTH_URL"
curl $DEPOSITAUTH_URL/version

exit 0

#
# end of file
#
