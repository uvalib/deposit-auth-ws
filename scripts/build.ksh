export GOPATH=$(pwd)

res=0
if [ $res -eq 0 ]; then
  env GOOS=darwin go build -o bin/deposit-auth-ws.darwin depositauthws
  res=$?
fi

if [ $res -eq 0 ]; then
  env GOOS=linux go build -o bin/deposit-auth-ws.linux depositauthws
  res=$?
fi

exit $res
