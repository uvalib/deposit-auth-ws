package sis

import (
   "time"
   "depositauthws/api"
    "log"
    "path/filepath"
    "fmt"
    "os"
    "io/ioutil"
    "regexp"
    "strings"
)

type SIS struct {
    ImportFs string
    ExportFs string
}

var Exchanger * SIS

// create a new exchanger... use as a singleton
func NewExchanger( ImportFs string, ExportFs string ) error {

    // check the import and export file systems
    err := checkReadableFs( ImportFs )
    if err != nil {
        return err
    }
    err = checkWritableFs( ExportFs )
    if err != nil {
        return err
    }

    Exchanger = &SIS{ ImportFs: ImportFs, ExportFs: ExportFs }
    return nil
}

// import any available records
func ( sis * SIS ) Import( ) ( [ ] * api.Authorization, error ) {

    // get a list of files to import
    files, err := locateImportFiles( sis.ImportFs )
    if err != nil {
        return nil, err
    }

    results := make([ ] * api.Authorization, 0 )

    // go through the list of import files
    for _, f := range files {
        imports, err := importFromFile( f )
        if err == nil {
            // add to our list
            for _, i := range imports {
               results = append( results, i )
            }
            // and remove the file
            markFileComplete( f )
        } else {
            log.Printf( "ERROR: while importing from %s", f )
            // handle the error
        }
    }

    return results, nil
}

// export the list of supplied records
func ( sis * SIS ) Export( exports [ ] * api.Authorization ) error {
    filename := createExportFilename( sis.ExportFs )
    return exportToFile( filename, exports )
}

// check the export filesystem to ensure it is available and writable
func ( sis * SIS ) CheckExport( ) error {
    return checkWritableFs( sis.ExportFs )
}

// check the import filesystem to ensure it is available and readable
func ( sis * SIS ) CheckImport( ) error {
    return checkReadableFs( sis.ImportFs )
}

// import records from the supplied file
func importFromFile( filename * string ) ( [ ] * api.Authorization, error ) {

    log.Printf( "Importing from: %s", *filename )
    // open the file for reading
    b, err := ioutil.ReadFile( *filename )
    if err != nil {
        return nil, err
    }

    results := make([ ] * api.Authorization, 0 )

    // tokenize by newline...
    contents := strings.Split( string( b ), "\n" )
    previous := ""
    for i := range contents {
        s := contents[ i ]
        // ignore empty records
        if len( s ) == 0 {
            continue
        }

        // handle the case where we get a truncated record (if the line is too long ???)
        if truncatedImportRecord( s ) == true {
           previous = s
        } else {
            r := createImportRecord( previous + s )
            if r != nil {
                results = append( results, r )
            } else {
                log.Printf( "ERROR: bad SIS record [%s]", previous + s )
                // handle the error here
            }
            previous = ""
        }
    }

    log.Printf( "%d record(s) loaded", len( results ) )
    return results, nil
}

// export the supplied records to the supplied file
func exportToFile( filename string, exports [ ] * api.Authorization ) error {

    // open the file, creating it if necessary and apending if it already exists
    f, err := os.OpenFile( filename, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0644 )
    if err != nil {
        return err
    }
    defer f.Close( )

    // write each record to the export file
    for _, i := range exports {
        rec := createExportRecord( i )
        if _, err = f.WriteString( rec ); err != nil {
            return err
        }
    }

    return nil
}

// determine that the record is truncated because it is 170 characters long
func truncatedImportRecord( rec string ) bool {
   return len( rec ) >= 170
}

// mark the file as done so we dont process it again
func markFileComplete( filename * string ) error {

    newname := fmt.Sprintf( "%s.done-%d", *filename, time.Now( ).UnixNano( ) )
    log.Printf( "Renaming %s -> %s", *filename, newname )
    return os.Rename( *filename, newname )
}

// look for any available files that can be imported
func locateImportFiles( filesystem string ) ( [ ] * string, error ) {
    files, err := ioutil.ReadDir( filesystem )
    if err != nil {
        return nil, err
    }

    results := make( [ ] * string, 0 )

    // the pattern to search for
    rx, _ := regexp.Compile( "UV_Libra_From_SIS_\\d{6}.txt$" )

    for _, f := range files {
        if rx.MatchString( f.Name( ) ) == true {
            fullname := filepath.Join( filesystem, f.Name( ) )
            results = append( results, &fullname )
        }
    }

    return results, nil
}

// create a filename that can be used for export
func createExportFilename( filesystem string )  string {
    yymmddDate := time.Now( ).Format( "060102" )
    filename := fmt.Sprintf( "UV_LIBRA_IN_%s.txt", yymmddDate )
    fullname := filepath.Join( filesystem, filename )
    log.Printf( "Exporting to: %s", fullname )
    return fullname
}

// create an export record from an authorization record
func createExportRecord( rec * api.Authorization ) string {
    delimiter := "|"
    r := fmt.Sprintf( "%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s\n",
        rec.EmployeeId, delimiter,
        rec.ComputingId, delimiter,
        rec.FirstName, delimiter,
        rec.MiddleName, delimiter,
        rec.LastName, delimiter,
        rec.Career, delimiter,
        rec.Program, delimiter,
        rec.Plan, delimiter,
        rec.LibraId, delimiter,
        rec.DocType, delimiter,
        rec.Degree, delimiter,
        dateToExportFormat( rec.AcceptedAt ) )
    return r
}

// create an export record from an authorization record
func createImportRecord( s string ) * api.Authorization {

    delimiter := "|"
    tokens := strings.Split( s, delimiter )
    if len( tokens ) == 12 {
        rec := api.Authorization{ }
        rec.EmployeeId = tokens[ 0 ]
        rec.ComputingId = tokens[ 1 ]
        rec.FirstName = tokens[ 2 ]
        rec.MiddleName = tokens[ 3 ]
        rec.LastName = tokens[ 4 ]
        rec.Career = tokens[ 5 ]
        rec.Program = tokens[ 6 ]
        rec.Plan = tokens[ 7 ]
        rec.Title = tokens[ 8 ]
        rec.DocType = tokens[ 9 ]
        rec.Degree = tokens[ 10 ]
        rec.ApprovedAt = dateToNativeFormat( tokens[ 11 ] )
        return &rec
    }
    return nil
}

// convert a date from the native format to the export format
func dateToExportFormat( date string ) string {

    layout := "2006-01-02 15:04:05"
    t, err := time.Parse( layout, date )
    if err != nil {
        return date
    }

    return t.Format( "01/02/2006" )
}

// convert a date from the export format to the native format
func dateToNativeFormat( date string ) string {

    layout := "01/02/2006"
    t, err := time.Parse( layout, date )
    if err != nil {
        return date
    }
    return t.Format( "2006-01-02" )
}

// check the supplied filesystem to ensure it is available and readable
func checkReadableFs( filesystem string ) error {
    return nil
}

// check the supplied filesystem to ensure it is available and writable
func checkWritableFs( filesystem string ) error {
    return nil
}