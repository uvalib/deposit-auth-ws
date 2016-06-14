package handlers

import (
    "net/http"
    "depositauthws/dao"
    "depositauthws/sis"
)

func HealthCheck( w http.ResponseWriter, r *http.Request ) {

    status := http.StatusOK
    db_err := dao.Database.Check( )
    import_err := sis.Exchanger.CheckImport( )
    export_err := sis.Exchanger.CheckExport( )

    var db_msg, import_msg, export_msg string

    if db_err != nil || import_err != nil || export_err != nil {

        status = http.StatusInternalServerError

        if db_err != nil {
            db_msg = db_err.Error( )
        }

        if import_err != nil {
            import_msg = import_err.Error( )
        }

        if export_err != nil {
            export_msg = export_err.Error( )
        }
    }

    EncodeHealthCheckResponse( w, status, db_msg, import_msg, export_msg )
}