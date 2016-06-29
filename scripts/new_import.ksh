#
#
#

# validate input parameters
if [ $# -eq 3 ]; then
   COMPUTEID=$1
   FIRSTNAME=$2
   LASTNAME=$3
else
   COMPUTEID=dpg3k
   FIRSTNAME=David
   LASTNAME=Goldstein
fi

# import filesystem root
if [ -z "$IMPORT_FS" ]; then
   IMPORT_FS="/tmp/import"
fi

# the sameple data file
SAMPLE=data/sample_from_sis.txt

# bomb out sample data does not
if [ ! -f $SAMPLE ]; then
   echo "ERROR: $SAMPLE data does not exist, aborting"
   exit 1
fi

# create the target filename
DATESTAMP=$(date "+%y%m%d")
TARGET=$IMPORT_FS/UV_Libra_From_SIS_$DATESTAMP.txt

# bomb out if it already exists
if [ -f $TARGET ]; then
   echo "ERROR: $TARGET already exists, aborting"
   exit 1
fi

cat $SAMPLE | sed -e "s/XXCOMPUTEIDXX/$COMPUTEID/g" | sed -e "s/XXFIRSTNAMEXX/$FIRSTNAME/g" | sed -e "s/XXLASTNAMEXX/$LASTNAME/g" | sed -e "s/XXRANDOMXX/$RANDOM/g" > $TARGET

echo "Created $TARGET"
exit 0

#
# end of file
#
