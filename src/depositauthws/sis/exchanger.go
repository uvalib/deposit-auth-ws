package sis

import (
   //"time"
   //"fmt"
   "depositauthws/api"
   //"encoding/json"
)

type SIS struct {
    ImportFs string
    ExportFs string
}

var Exchanger * SIS

func NewExchanger( ImportFs string, ExportFs string ) error {

    // check the import and export file systems
    err := checkImportFs( ImportFs )
    if err != nil {
        return err
    }
    err = checkExportFs( ExportFs )
    if err != nil {
        return err
    }

    Exchanger = &SIS{ ImportFs: ImportFs, ExportFs: ExportFs }
    return nil
}

func ( sis * SIS ) Import( ) ( [] * api.Authorization, error ) {
    return nil, nil
}

func ( sis * SIS ) Export( exports [] * api.Authorization ) error {
    return nil
}

func ( sis * SIS ) CheckExport( ) error {
    return checkExportFs( sis.ExportFs )
}

func ( sis * SIS ) CheckImport( ) error {
    return checkImportFs( sis.ImportFs )
}

func checkImportFs( filesystem string ) error {
    return nil
}

func checkExportFs( filesystem string ) error {
    return nil
}