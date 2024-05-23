package sis

import (
	"fmt"
	"github.com/uvalib/deposit-auth-ws/depositauthws/api"
	"github.com/uvalib/deposit-auth-ws/depositauthws/logger"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

type sisImplementation struct {
	ImportFs string
	ExportFs string
}

// Exchanger -- the file exchanger used to communicate with SIS
var Exchanger *sisImplementation

// NewExchanger -- create a new exchanger... use as a singleton
func NewExchanger(ImportFs string, ExportFs string) error {

	// check the import and export file systems
	//err := checkReadableFs( ImportFs )
	//if err != nil {
	//    return err
	//}
	//err = checkWritableFs( ExportFs )
	//if err != nil {
	//    return err
	//}

	Exchanger = &sisImplementation{ImportFs: ImportFs, ExportFs: ExportFs}
	return nil
}

// Import -- import any available records
func (sis *sisImplementation) Import() ([]*api.Authorization, error) {

	// get a list of files to import
	files, err := locateImportFiles(sis.ImportFs)
	if err != nil {
		return nil, err
	}

	results := make([]*api.Authorization, 0)

	// go through the list of import files
	for _, f := range files {
		imports, err := importFromFile(f)
		if err == nil {
			// add to our list
			for _, i := range imports {
				results = append(results, i)
			}
			// and remove the file
			markFileComplete(f)
		} else {
			logger.Log(fmt.Sprintf("ERROR: while importing from: %s", f))
			// handle the error
		}
	}

	return results, nil
}

// Export -- export the list of supplied records
func (sis *sisImplementation) Export(exports []*api.Authorization) error {
	filename := createExportFilename(sis.ExportFs)
	return exportToFile(filename, exports)
}

// CheckExport -- check the export filesystem to ensure it is available and writable
func (sis *sisImplementation) CheckExport() error {
	return checkWritableFs(sis.ExportFs)
}

// CheckImport -- check the import filesystem to ensure it is available and readable
func (sis *sisImplementation) CheckImport() error {
	return checkReadableFs(sis.ImportFs)
}

// import records from the supplied file
func importFromFile(filename string) ([]*api.Authorization, error) {

	logger.Log(fmt.Sprintf("INFO: importing from: %s", filename))

	// open the file for reading
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	results := make([]*api.Authorization, 0)

	// tokenize by newline...
	contents := strings.Split(string(b), "\n")

	//previous := ""
	for i := range contents {
		s := contents[i]

		// ignore empty records
		if len(s) == 0 {
			continue
		}

		converted := convertToUtf8(s)

		r := createImportRecord(converted)
		if r != nil {
			results = append(results, r)
		} else {
			logger.Log(fmt.Sprintf("ERROR: bad record, ignoring [%s]", s))
			// handle the error here
		}
	}

	logger.Log(fmt.Sprintf("INFO: import success from: %s, %d record(s) loaded", filename, len(results)))
	return results, nil
}

// export the supplied records to the supplied file
func exportToFile(filename string, exports []*api.Authorization) error {

	// open the file, creating it if necessary and apending if it already exists
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// write each record to the export file
	for _, i := range exports {
		rec := createExportRecord(i)
		if _, err = f.WriteString(rec); err != nil {
			return err
		}
	}

	return nil
}

// attempt to convert the character string to utf8
func convertToUtf8(si string) string {

	//logger.Log( fmt.Sprintf( "==> IN: [%s]", si ) )
	ba := []byte(si)
	buf := make([]rune, len(ba))
	for i, b := range ba {
		buf[i] = rune(b)
	}
	so := string(buf)
	//    logger.Log( fmt.Sprintf( "==> OUT: [%s]", so ) )
	return (so)
}

// mark the file as done so we don't process it again
func markFileComplete(filename string) error {

	newname := fmt.Sprintf("%s.done-%d", filename, time.Now().UnixNano())
	logger.Log(fmt.Sprintf("INFO: renaming %s -> %s", filename, newname))
	return os.Rename(filename, newname)
}

// look for any available files that can be imported
func locateImportFiles(filesystem string) ([]string, error) {
	files, err := ioutil.ReadDir(filesystem)
	if err != nil {
		return nil, err
	}

	results := make([]string, 0)

	// the pattern to search for
	rx, _ := regexp.Compile("UV_Libra_From_SIS_\\d{6}.txt$")

	for _, f := range files {
		if rx.MatchString(f.Name()) == true {
			fullname := filepath.Join(filesystem, f.Name())
			results = append(results, fullname)
		}
	}

	sort.Strings(results)
	return results, nil
}

// create a filename that can be used for export
func createExportFilename(filesystem string) string {
	yymmddFormat := "060102"
	t := time.Now()
	yymmddDate := t.Format(yymmddFormat)
	// we strive to exist in UTC so attempt to localize the timestamp
	location, err := time.LoadLocation("EST")
	if err == nil {
		yymmddDate = t.In(location).Format(yymmddFormat)
	}
	filename := fmt.Sprintf("UV_LIBRA_IN_%s.txt", yymmddDate)
	fullname := filepath.Join(filesystem, filename)
	logger.Log(fmt.Sprintf("INFO: exporting to: %s", fullname))
	return fullname
}

// create an export record from an authorization record
func createExportRecord(rec *api.Authorization) string {
	delimiter := "|"
	r := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s\n",
		rec.EmployeeID, delimiter,
		rec.ComputingID, delimiter,
		rec.FirstName, delimiter,
		rec.MiddleName, delimiter,
		rec.LastName, delimiter,
		rec.Career, delimiter,
		rec.Program, delimiter,
		rec.Plan, delimiter,
		fullIdentifierURL(rec.LibraID), delimiter,
		rec.DocType, delimiter,
		rec.Degree, delimiter,
		dateToExportFormat(rec.AcceptedAt))
	return r
}

// create an export record from an authorization record
func createImportRecord(s string) *api.Authorization {

	//log.Printf("==> [%s]", s)

	delimiter := "|"
	tokens := strings.Split(s, delimiter)
	if len(tokens) == 12 {
		rec := api.Authorization{}
		rec.EmployeeID = tokens[0]
		rec.ComputingID = tokens[1]
		rec.FirstName = tokens[2]
		rec.MiddleName = tokens[3]
		rec.LastName = tokens[4]
		rec.Career = tokens[5]
		rec.Program = tokens[6]
		rec.Plan = tokens[7]
		rec.Title = tokens[8]
		rec.DocType = tokens[9]
		rec.Degree = tokens[10]
		rec.ApprovedAt = dateToNativeFormat(tokens[11])

		//log.Printf( "Title [%s]", rec.Title )

		return &rec
	}
	return nil
}

// convert the document identifier (DOI) to a fully qualified URL
func fullIdentifierURL(identifier string) string {
	base := "http://dx.doi.org"
	if len(identifier) == 0 {
		return base
	}

	return fmt.Sprintf("%s/%s", base, strings.Replace(identifier, "doi:", "", 1))
}

// convert a date from the native format to the export format
func dateToExportFormat(date string) string {

	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return date
	}

	return t.Format("01/02/2006")
}

// convert a date from the export format to the native format
func dateToNativeFormat(date string) string {

	layout := "01/02/2006"
	t, err := time.Parse(layout, date)
	if err != nil {
		return date
	}
	return t.Format("2006-01-02")
}

// check the supplied filesystem to ensure it is available and readable
func checkReadableFs(filesystem string) error {

	f, err := os.Stat(filesystem)
	if err != nil {
		return err
	}

	if f.IsDir() == false {
		return fmt.Errorf("%s is not a directory", filesystem)
	}

	_, err = ioutil.ReadDir(filesystem)
	if err != nil {
		return err
	}

	return nil
}

// check the supplied filesystem to ensure it is available and writable
func checkWritableFs(filesystem string) error {

	f, err := os.Stat(filesystem)
	if err != nil {
		return err
	}

	if f.IsDir() == false {
		return fmt.Errorf("%s is not a directory", filesystem)
	}

	tmpfile, err := ioutil.TempFile(filesystem, "tmpfile")
	if err != nil {
		return err
	}

	defer os.Remove(tmpfile.Name())

	return tmpfile.Close()
}

//
// end of file
//
