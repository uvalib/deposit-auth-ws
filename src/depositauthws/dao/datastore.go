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

func ( db *DB ) GetDepositRequest( id string ) ( [] * api.Registration, error ) {

    rows, err := db.Query( "SELECT * FROM depositrequest WHERE id = ? LIMIT 1", id )
    if err != nil {
        return nil, err
    }
    defer rows.Close( )

    return depositRequestResults( rows )
}

func ( db *DB ) SearchDepositRequest( id string ) ( [] * api.Registration, error ) {

    rows, err := db.Query( "SELECT * FROM depositrequest WHERE id > ?", id )
    if err != nil {
        return nil, err
    }
    defer rows.Close( )

    return depositRequestResults( rows )
}

func ( db *DB ) CreateDepositRequest( reg api.Registration ) ( * api.Registration, error ) {

    stmt, err := db.Prepare( "INSERT INTO depositrequest( requester, user, department, degree ) VALUES(?,?,?,?)" )
    if err != nil {
        return nil, err
    }

    res, err := stmt.Exec( reg.Requester, reg.For, reg.Department, reg.Degree )
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

func ( db *DB ) DeleteDepositRequest( id string ) ( int64, error ) {

    stmt, err := db.Prepare( "DELETE FROM depositrequest WHERE id = ? LIMIT 1" )
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

func ( db *DB ) GetFieldSet( field_name string ) ( [] string, error ) {
    rows, err := db.Query( "SELECT field_value FROM fieldvalues WHERE field_name = ?", field_name )
    if err != nil {
        return nil, err
    }
    defer rows.Close( )

    return fieldSetResults( rows )
}

func depositRequestResults( rows * sql.Rows ) ( [] * api.Registration, error ) {

    var optional sql.NullString

    results := make([ ] * api.Registration, 0 )
    for rows.Next() {
        reg := new( api.Registration )
        err := rows.Scan( &reg.Id,
            &reg.Requester,
            &reg.For,
            &reg.Department,
            &reg.Degree,
            &reg.Status,
            &reg.RequestDate,
            &optional )
        if err != nil {
            return nil, err
        }
        if optional.Valid {
            reg.DepositDate = optional.String
        }
        results = append( results, reg )
    }
    if err := rows.Err( ); err != nil {
        return nil, err
    }

    log.Printf( "Returning %d row(s)", len( results ) )
    return results, nil
}

func fieldSetResults( rows * sql.Rows ) ( [] string, error ) {

    results := make([ ] string, 0 )
    for rows.Next() {
        var s string
        err := rows.Scan( &s )
        if err != nil {
            return nil, err
        }
        results = append( results, s )
    }

    log.Printf( "Returning %d row(s)", len( results ) )
    return results, nil
}