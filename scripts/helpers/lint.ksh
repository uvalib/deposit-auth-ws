if [ -z "$GOPATH" ]; then
   echo "ERROR: GOPATH is not defined"
   exit 1
fi

golint depositauthws/...
res=$?

echo "Exiting with status $res"
exit $res