#
#
#

# import filesystem root
if [ -z "$IMPORT_FS" ]; then
   IMPORT_FS="/tmp/import"
fi

# export filesystem root
if [ -z "$EXPORT_FS" ]; then
   EXPORT_FS="/tmp/export"
fi

echo "***********************************************************************"
echo "* Import: $IMPORT_FS"
echo "***********************************************************************"
ls -l $IMPORT_FS/UV_Libra_From_SIS*.txt 2>&1 | grep -v "No such file or directory"

echo ""
echo "***********************************************************************"
echo "* Export: $EXPORT_FS"
echo "***********************************************************************"
ls -l $EXPORT_FS/UV_LIBRA_IN*.txt 2>&1 | grep -v "No such file or directory"

exit 0

#
# end of file
#
