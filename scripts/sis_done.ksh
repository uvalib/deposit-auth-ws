#
# Show the previously processed import files
#

# import filesystem root
if [ -z "$IMPORT_FS" ]; then
   IMPORT_FS="/tmp/import"
fi

echo ""
echo "***********************************************************************"
echo "* Completed import from: $IMPORT_FS"
echo "***********************************************************************"
ls $IMPORT_FS/UV_Libra_From_SIS*.txt.done-* 2>&1 | grep -v "No such file or directory"

echo ""
echo ""
exit 0

#
# end of file
#
