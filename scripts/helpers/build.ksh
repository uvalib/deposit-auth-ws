if [ -z "$GOPATH" ]; then
   echo "ERROR: GOPATH is not defined"
   exit 1
fi

res=0
if [ $res -eq 0 ]; then
  GOOS=darwin go build -a -o bin/deposit-auth-ws.darwin depositauthws
  res=$?
fi

if [ $res -eq 0 ]; then
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/deposit-auth-ws.linux depositauthws
  res=$?
fi

exit $res
