if [ -z "$DBPASSWD" ]; then
   echo "ERROR: DBPASSWD must be defined"
   exit 1
fi

bin/deposit-auth-ws.darwin --dbpassword $DBPASSWD
