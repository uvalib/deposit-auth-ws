#
# split a sis file into 2 based on a list of computing ID's that appear in a seperate file
#

# check command line count
if [ $# -ne 2 ]; then
   echo "use: %0 <sis file> <computing id file>"
   exit 1
fi

SIS_FILE=$1
CID_FILE=$2

# check the input files exist
if [ ! -f $SIS_FILE ]; then
   echo "ERROR: SIS file $SIS_FILE does not exist, aborting"
   exit 1
fi

if [ ! -f $CID_FILE ]; then
   echo "ERROR: CID file $CID_FILE does not exist, aborting"
   exit 1
fi

SIS_OUT_INCLUDED=$SIS_FILE.included
SIS_OUT_EXCLUDED=$SIS_FILE.excluded

# clean and create the output files
rm -fr $SIS_OUT_INCLUDED > /dev/null 2>&1
rm -fr $SIS_OUT_EXCLUDED > /dev/null 2>&1
touch $SIS_OUT_INCLUDED
touch $SIS_OUT_EXCLUDED

while read line; do

   cid=$(echo $line | awk -F'|' '{print $2}')
   included=$(grep "^$cid$" $CID_FILE | wc -l | awk '{print $1}')
   #echo "[$cid] $included"

   if [ "$included" == "1" ]; then
      echo $line >> $SIS_OUT_INCLUDED
   else
      echo $line >> $SIS_OUT_EXCLUDED
   fi

done < $SIS_FILE

# get output file record count
INCLUDED_COUNT=$(wc -l < $SIS_OUT_INCLUDED | awk '{print $1}')
EXCLUDED_COUNT=$(wc -l < $SIS_OUT_EXCLUDED | awk '{print $1}')

echo "Terminating normally"
echo "$SIS_OUT_INCLUDED contains $INCLUDED_COUNT record(s)"
echo "$SIS_OUT_EXCLUDED contains $EXCLUDED_COUNT record(s)"

exit 0

#
# end of file
#
