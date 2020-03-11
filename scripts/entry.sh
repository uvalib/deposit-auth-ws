# set blank options variables
DBSECURE_OPT=""
DBHOST_OPT=""
DBNAME_OPT=""
DBUSER_OPT=""
DBPASSWD_OPT=""
IMPORT_OPT=""
EXPORT_OPT=""
SECRET_OPT=""
DEBUG_OPT=""

# secure database access
if [ -n "$DBSECURE" ]; then
   DBSECURE_OPT="--dbsecure=$DBSECURE"
fi

# database host
if [ -n "$DBHOST" ]; then
   DBHOST_OPT="--dbhost $DBHOST"
fi

# database name
if [ -n "$DBNAME" ]; then
   DBNAME_OPT="--dbname $DBNAME"
fi

# database user
if [ -n "$DBUSER" ]; then
   DBUSER_OPT="--dbuser $DBUSER"
fi

# database password
if [ -n "$DBPASSWD" ]; then
   DBPASSWD_OPT="--dbpassword $DBPASSWD"
fi

# import filesystem root
if [ -n "$IMPORT_FS" ]; then
   IMPORT_OPT="--importfs $IMPORT_FS"
fi

# export filesystem root
if [ -n "$EXPORT_FS" ]; then
   EXPORT_OPT="--exportfs $EXPORT_FS"
fi

# shared secret
if [ -n "$AUTH_SHARED_SECRET" ]; then
   SECRET_OPT="--secret $AUTH_SHARED_SECRET"
fi

# service debugging
if [ -n "$DEPOSITAUTH_DEBUG" ]; then
   DEBUG_OPT="--debug=$DEPOSITAUTH_DEBUG"
fi

bin/deposit-auth-ws $DBSECURE_OPT $DBHOST_OPT $DBNAME_OPT $DBUSER_OPT $DBPASSWD_OPT $IMPORT_OPT $EXPORT_OPT $SECRET_OPT $DEBUG_OPT

#
# end of file
#
