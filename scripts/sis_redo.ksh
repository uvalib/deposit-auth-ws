#
# Rename a previously processed import file so it is processed again
#

# check argument count
if [ $# -ne 1 ]; then
   echo "ERROR: specify a file timestamp for reprocessing"
   exit 1
fi

# import filesystem root
if [ -z "$IMPORT_FS" ]; then
   IMPORT_FS="/tmp/import"
fi

# import file attributes
TIMESTAMP=$1
FILEBASE=$IMPORT_FS/UV_Libra_From_SIS*.txt.done-$TIMESTAMP
FILE=$(ls $FILEBASE 2>/dev/null)

# do we have that file
if [ -z "$FILE" ]; then
   echo "ERROR: that import file does not exist"
   exit 1
fi

# create the new name
NEWNAME=${FILE%.done-$TIMESTAMP}

# and rename it
mv $FILE $NEWNAME
echo ""
echo "file renamed; now run the import process"
exit 0

#
# end of file
#
