package dao

import (
   "database/sql"
   "depositauthws/api"
   "depositauthws/logger"
   "fmt"
   // needed by the linter
   _ "github.com/go-sql-driver/mysql"
   "strconv"
)

type dbStruct struct {
   *sql.DB
}

type mapper struct {
   FieldClass  string
   FieldSource string
   FieldMapped string
}

//
// DB -- the database instance
//
var DB *dbStruct

//
// NewDB -- create the database singletomn
//
func NewDB(dataSourceName string) error {
   db, err := sql.Open("mysql", dataSourceName)
   if err != nil {
      return err
   }
   if err = db.Ping(); err != nil {
      return err
   }
   DB = &dbStruct{db}
   return nil
}

//
// CheckDB -- check our database health
//
func (db *dbStruct) CheckDB() error {
   return db.Ping()
}

//
// DepositAuthorizationExists -- determine if the supplied deposit authorization already exists
//
func (db *dbStruct) DepositAuthorizationExists(e api.Authorization) (bool, error) {

   rows, err := db.Query("SELECT COUNT(*) FROM depositauth WHERE computing_id = ? AND degree = ? AND plan = ? AND title = ?", e.ComputingID, e.Degree, e.Plan, e.Title)
   if err != nil {
      return false, err
   }
   defer rows.Close()

   var count int
   for rows.Next() {
      err := rows.Scan(&count)
      if err != nil {
         return false, err
      }
   }

   if err := rows.Err(); err != nil {
      return false, err
   }

   return count != 0, nil
}

//
// GetDepositAuthorizationByID -- get all by ID (should only be 1)
//
func (db *dbStruct) GetDepositAuthorizationByID(id string) ([]*api.Authorization, error) {

   rows, err := db.Query("SELECT * FROM depositauth WHERE id = ? LIMIT 1", id)
   if err != nil {
      return nil, err
   }
   defer rows.Close()

   return depositAuthorizationResults(rows)
}

//
// SearchDepositAuthorizationByID -- get all greater than a specified ID
//
func (db *dbStruct) SearchDepositAuthorizationByID(id string) ([]*api.Authorization, error) {

   rows, err := db.Query("SELECT * FROM depositauth WHERE id > ? ORDER BY id ASC", id)
   if err != nil {
      return nil, err
   }
   defer rows.Close()

   return depositAuthorizationResults(rows)
}

//
// SearchDepositAuthorizationByCid -- get all similar to the a specified computing ID
//
func (db *dbStruct) SearchDepositAuthorizationByCid(cid string) ([]*api.Authorization, error) {

   rows, err := db.Query("SELECT * FROM depositauth WHERE computing_id LIKE ? ORDER BY id ASC", fmt.Sprintf("%s%%", cid))
   if err != nil {
      return nil, err
   }
   defer rows.Close()

   return depositAuthorizationResults(rows)
}

//
// SearchDepositAuthorizationByCreateDate -- get all greater than a specified created date
//
func (db *dbStruct) SearchDepositAuthorizationByCreateDate(createdAt string) ([]*api.Authorization, error) {

   rows, err := db.Query("SELECT * FROM depositauth WHERE created_at >= ? ORDER BY id ASC", createdAt)
   if err != nil {
      return nil, err
   }
   defer rows.Close()

   return depositAuthorizationResults(rows)
}

//
// SearchDepositAuthorizationByExportDate -- get all greater than a specified exported date
//
func (db *dbStruct) SearchDepositAuthorizationByExportDate(exportedAt string) ([]*api.Authorization, error) {

   rows, err := db.Query("SELECT * FROM depositauth WHERE exported_at >= ? ORDER BY id ASC", exportedAt)
   if err != nil {
      return nil, err
   }
   defer rows.Close()

   return depositAuthorizationResults(rows)
}

//
// CreateDepositAuthorization -- create a new deposit authorization
//
func (db *dbStruct) CreateDepositAuthorization(reg api.Authorization) (*api.Authorization, error) {

   stmt, err := db.Prepare("INSERT INTO depositauth( employee_id, computing_id, first_name, middle_name, last_name, career, program, plan, degree, title, doctype, approved_at ) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
   if err != nil {
      return nil, err
   }

   res, err := stmt.Exec(reg.EmployeeID,
      reg.ComputingID,
      reg.FirstName,
      reg.MiddleName,
      reg.LastName,
      reg.Career,
      reg.Program,
      reg.Plan,
      reg.Degree,
      reg.Title,
      reg.DocType,
      reg.ApprovedAt)
   if err != nil {
      return nil, err
   }

   lastID, err := res.LastInsertId()
   if err != nil {
      return nil, err
   }

   reg.ID = strconv.FormatInt(lastID, 10)
   return &reg, nil
}

//
// DeleteDepositAuthorizationByID -- delete by ID
//
func (db *dbStruct) DeleteDepositAuthorizationByID(id string) (int64, error) {

   stmt, err := db.Prepare("DELETE FROM depositauth WHERE id = ? LIMIT 1")
   if err != nil {
      return 0, err
   }

   res, err := stmt.Exec(id)
   if err != nil {
      return 0, err
   }

   rowCount, err := res.RowsAffected()
   if err != nil {
      return 0, err
   }

   return rowCount, nil
}

//
// GetDepositAuthorizationForExport -- get all available for export
//
func (db *dbStruct) GetDepositAuthorizationForExport() ([]*api.Authorization, error) {

   rows, err := db.Query("SELECT * FROM depositauth WHERE accepted_at IS NOT NULL AND exported_at IS NULL ORDER BY id ASC")
   if err != nil {
      return nil, err
   }
   defer rows.Close()

   return depositAuthorizationResults(rows)
}

//
// UpdateExportedDepositAuthorization -- update all export items with the time of export
//
func (db *dbStruct) UpdateExportedDepositAuthorization(exports []*api.Authorization) error {

   stmt, err := db.Prepare("UPDATE depositauth SET exported_at = NOW( ) WHERE id = ? LIMIT 1")
   if err != nil {
      return err
   }

   for _, rec := range exports {
      _, err := stmt.Exec(rec.ID)
      if err != nil {
         return err
      }
   }

   return nil
}

//
// UpdateFulfilledDepositAuthorizationByID -- update an item that has been 'fulfilled'
//
func (db *dbStruct) UpdateFulfilledDepositAuthorizationByID(id string, did string) error {

   stmt, err := db.Prepare("UPDATE depositauth SET exported_at = NULL, accepted_at = NOW( ), status = ?, libra_id = ? WHERE id = ? LIMIT 1")
   if err != nil {
      return err
   }

   _, err = stmt.Exec("submitted", did, id)
   if err != nil {
      return err
   }

   return nil
}

//
// GetFieldMapperList -- get the list of field maps
//
func (db *dbStruct) GetFieldMapperList() ([]*mapper, error) {

   rows, err := db.Query("SELECT field_class, field_name, field_value FROM fieldmapper")
   if err != nil {
      return nil, err
   }
   defer rows.Close()

   results := make([]*mapper, 0)

   for rows.Next() {
      mapping := new(mapper)
      err := rows.Scan(
         &mapping.FieldClass,
         &mapping.FieldSource,
         &mapping.FieldMapped)
      if err != nil {
         return nil, err
      }

      results = append(results, mapping)
   }

   if err := rows.Err(); err != nil {
      return nil, err
   }

   return results, nil
}

//
// private implementation methods
//

func depositAuthorizationResults(rows *sql.Rows) ([]*api.Authorization, error) {

   var optionalApprovedAt sql.NullString
   var optionalAcceptedAt sql.NullString
   var optionalExportedAt sql.NullString
   var optionalUpdatedAt sql.NullString

   results := make([]*api.Authorization, 0)
   for rows.Next() {
      reg := new(api.Authorization)
      err := rows.Scan(&reg.ID,
         &reg.EmployeeID,
         &reg.ComputingID,
         &reg.FirstName,
         &reg.MiddleName,
         &reg.LastName,
         &reg.Career,
         &reg.Program,
         &reg.Plan,
         &reg.Degree,
         &reg.Title,
         &reg.DocType,
         &reg.LibraID,
         &reg.Status,
         &optionalApprovedAt,
         &optionalAcceptedAt,
         &optionalExportedAt,
         &reg.CreatedAt,
         &optionalUpdatedAt)
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

      results = append(results, reg)
   }
   if err := rows.Err(); err != nil {
      return nil, err
   }

   logger.Log(fmt.Sprintf("Deposit authorization request returns %d row(s)", len(results)))
   return results, nil
}
