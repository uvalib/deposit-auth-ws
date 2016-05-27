package dao

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "depositauthws/api"
    "log"
    "strconv"
)

type DB struct {
    *sql.DB
}

type Mapper struct {
    FieldClass  string
    FieldSource string
    FieldMapped string
}

var Database * DB

func NewDB( dataSourceName string ) error {
    db, err := sql.Open( "mysql", dataSourceName )
    if err != nil {
        return err
    }
    if err = db.Ping( ); err != nil {
        return err
    }
    Database = &DB{ db }
    return nil
}

func ( db *DB ) Check( ) error {
    if err := db.Ping( ); err != nil {
        return err
    }
    return nil
}

func ( db *DB ) GetDepositAuthorizationById( id string ) ( [] * api.Authorization, error ) {

    rows, err := db.Query( "SELECT * FROM depositauth WHERE id = ? LIMIT 1", id )
    if err != nil {
        return nil, err
    }
    defer rows.Close( )

    return depositAuthorizationResults( rows )
}

func ( db *DB ) SearchDepositAuthorization( id string ) ( [] * api.Authorization, error ) {

    rows, err := db.Query( "SELECT * FROM depositauth WHERE id > ? ORDER BY id ASC", id )
    if err != nil {
        return nil, err
    }
    defer rows.Close( )

    return depositAuthorizationResults( rows )
}

func ( db *DB ) CreateDepositAuthorization( reg api.Authorization ) ( * api.Authorization, error ) {

    stmt, err := db.Prepare( "INSERT INTO depositauth( employee_id, computing_id, first_name, middle_name, last_name, career, program, plan, degree, title, doctype, approved_at ) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)" )
    if err != nil {
        return nil, err
    }

    res, err := stmt.Exec( reg.EmployeeId,
                           reg.ComputingId,
                           reg.FirstName,
                           reg.MiddleName,
                           reg.LastName,
                           reg.Career,
                           reg.Program,
                           reg.Plan,
                           reg.Degree,
                           reg.Title,
                           reg.DocType,
                           reg.ApprovedAt )
    if err != nil {
        return nil, err
    }

    lastId, err := res.LastInsertId( )
    if err != nil {
        return nil, err
    }

    reg.Id = strconv.FormatInt( lastId, 10 )
    return &reg, nil
}

func ( db *DB ) DeleteDepositAuthorizationById( id string ) ( int64, error ) {

    stmt, err := db.Prepare( "DELETE FROM depositauth WHERE id = ? LIMIT 1" )
    if err != nil {
        return 0, err
    }

    res, err := stmt.Exec( id )
    if err != nil {
        return 0, err
    }

    rowCount, err := res.RowsAffected( )
    if err != nil {
        return 0, err
    }

    return rowCount, nil
}

func ( db *DB ) GetDepositAuthorizationForExport( ) ( [] * api.Authorization, error ) {

    rows, err := db.Query( "SELECT * FROM depositauth WHERE accepted_at IS NOT NULL AND exported_at IS NULL ORDER BY id ASC" )
    if err != nil {
        return nil, err
    }
    defer rows.Close( )

    return depositAuthorizationResults( rows )
}

func ( db *DB ) UpdatedExportedDepositAuthorization( exports [] * api.Authorization ) error {

    stmt, err := db.Prepare( "UPDATE depositauth SET exported_at = NOW( ) WHERE id = ? LIMIT 1" )
    if err != nil {
        return err
    }

    for _, rec := range exports {
        _, err := stmt.Exec( rec.Id )
        if err != nil {
            return err
        }
    }

    return nil
}

func ( db *DB ) GetFieldMapperList( ) ( [] * Mapper, error ) {

    rows, err := db.Query( "SELECT field_class, field_name, field_value FROM fieldmapper" )
    if err != nil {
        return nil, err
    }
    defer rows.Close( )

    results := make([ ] * Mapper, 0 )

    for rows.Next() {
        mapping := new( Mapper )
        err := rows.Scan(
            &mapping.FieldClass,
            &mapping.FieldSource,
            &mapping.FieldMapped )
        if err != nil {
            return nil, err
        }

        results = append( results, mapping )
    }

    if err := rows.Err( ); err != nil {
        return nil, err
    }

    return results, nil
}

//
// private implementation methods
//

func depositAuthorizationResults( rows * sql.Rows ) ( [] * api.Authorization, error ) {

    var optionalApprovedAt sql.NullString
    var optionalAcceptedAt sql.NullString
    var optionalExportedAt sql.NullString
    var optionalUpdatedAt sql.NullString

    results := make([ ] * api.Authorization, 0 )
    for rows.Next() {
        reg := new( api.Authorization )
        err := rows.Scan( &reg.Id,
            &reg.EmployeeId,
            &reg.ComputingId,
            &reg.FirstName,
            &reg.MiddleName,
            &reg.LastName,
            &reg.Career,
            &reg.Program,
            &reg.Plan,
            &reg.Degree,
            &reg.Title,
            &reg.DocType,
            &reg.LibraId,
            &reg.Status,
            &optionalApprovedAt,
            &optionalAcceptedAt,
            &optionalExportedAt,
            &reg.CreatedAt,
            &optionalUpdatedAt )
        if err != nil {
            return nil, err
        }

        if optionalApprovedAt.Valid {
            reg.ApprovedAt = optionalApprovedAt.String
        }

        if optionalAcceptedAt.Valid {
            reg.AcceptedAt = optionalAcceptedAt.String
        }

        if optionalExportedAt.Valid {
            reg.ExportedAt = optionalExportedAt.String
        }

        if optionalUpdatedAt.Valid {
            reg.UpdatedAt = optionalUpdatedAt.String
        }

        results = append( results, reg )
    }
    if err := rows.Err( ); err != nil {
        return nil, err
    }

    log.Printf( "Returning %d row(s)", len( results ) )
    return results, nil
}