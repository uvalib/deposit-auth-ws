package dao

import (
	"database/sql"
	"fmt"
	"github.com/uvalib/deposit-auth-ws/depositauthws/api"
	"github.com/uvalib/deposit-auth-ws/depositauthws/config"
	"github.com/uvalib/deposit-auth-ws/depositauthws/logger"
	"time"

	// needed by the linter
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

// this is our DB implementation
type storage struct {
	*sql.DB
}

//
// newDBStore -- create a DB version of the storage singleton
//
func newDBStore() (Storage, error) {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowOldPasswords=1&tls=%s&sql_notes=false&timeout=%s&readTimeout=%s&writeTimeout=%s",
		config.Configuration.DbUser,
		config.Configuration.DbPassphrase,
		config.Configuration.DbHost,
		config.Configuration.DbName,
		config.Configuration.DbSecure,
		config.Configuration.DbTimeout,
		config.Configuration.DbTimeout,
		config.Configuration.DbTimeout)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	//taken from https://github.com/go-sql-driver/mysql/issues/461
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(2)

	return &storage{db}, nil
}

//
// Check -- check our database health
//
func (s *storage) Check() error {
	return s.Ping()
}

//
// GetInbound -- get all inbound after the specified ID
//
func (s *storage) GetInbound(after string) ([]*api.Authorization, error) {

	rows, err := s.Query("SELECT i.id, d.* FROM depositauth d, inbound i WHERE d.id = i.deposit_id AND i.id > ? ORDER BY i.id ASC", after)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return depositAuthorizationResults(rows)
}

//
// DepositAuthorizationExists -- determine if the supplied deposit authorization already exists
//
func (s *storage) GetMatchingDepositAuthorization(e api.Authorization) ([]*api.Authorization, error) {

	rows, err := s.Query("SELECT 0, d.* FROM depositauth d WHERE d.computing_id = ? AND d.degree = ? AND d.plan = ? ORDER BY d.id ASC", e.ComputingID, e.Degree, e.Plan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return depositAuthorizationResults(rows)
}

//
// GetDepositAuthorizationByID -- get all by ID (should only be 1)
//
func (s *storage) GetDepositAuthorizationByID(id string) ([]*api.Authorization, error) {

	rows, err := s.Query("SELECT 0, d.* FROM depositauth d WHERE d.id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return depositAuthorizationResults(rows)
}

//
// SearchDepositAuthorizationByCid -- get all similar to the a specified computing ID
//
func (s *storage) SearchDepositAuthorizationByCid(cid string) ([]*api.Authorization, error) {

	rows, err := s.Query("SELECT 0, d.* FROM depositauth d WHERE d.computing_id LIKE ? ORDER BY d.id ASC", fmt.Sprintf("%s%%", cid))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return depositAuthorizationResults(rows)
}

//
// SearchDepositAuthorizationByCreateDate -- get all greater than a specified created date
//
func (s *storage) SearchDepositAuthorizationByCreateDate(createdAt string) ([]*api.Authorization, error) {

	rows, err := s.Query("SELECT 0, d.* FROM depositauth d WHERE d.created_at >= ? ORDER BY d.id ASC", createdAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return depositAuthorizationResults(rows)
}

//
// SearchDepositAuthorizationByExportDate -- get all greater than a specified exported date
//
func (s *storage) SearchDepositAuthorizationByExportDate(exportedAt string) ([]*api.Authorization, error) {

	rows, err := s.Query("SELECT 0, d.* FROM depositauth d WHERE d.exported_at >= ? ORDER BY d.id ASC", exportedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return depositAuthorizationResults(rows)
}

//
// CreateInbound -- create a new inbound record
//
func (s *storage) CreateInbound(authID string) error {

	stmt, err := s.Prepare("INSERT INTO inbound( deposit_id ) VALUES(?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(authID)
	return err
}

//
// CreateDepositAuthorization -- create a new deposit authorization
//
func (s *storage) CreateDepositAuthorization(reg api.Authorization) (*api.Authorization, error) {

	stmt, err := s.Prepare("INSERT INTO depositauth( employee_id, computing_id, first_name, middle_name, last_name, career, program, plan, degree, title, doctype, approved_at ) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
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
func (s *storage) DeleteDepositAuthorizationByID(id string) (int64, error) {

	stmt, err := s.Prepare("DELETE FROM depositauth WHERE id = ? LIMIT 1")
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
func (s *storage) GetDepositAuthorizationForExport() ([]*api.Authorization, error) {

	rows, err := s.Query("SELECT 0, d.* FROM depositauth d WHERE d.accepted_at IS NOT NULL AND d.exported_at IS NULL ORDER BY d.id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return depositAuthorizationResults(rows)
}

//
// UpdateExportedDepositAuthorization -- update all export items with the time of export
//
func (s *storage) UpdateExportedDepositAuthorization(exports []*api.Authorization) error {

	stmt, err := s.Prepare("UPDATE depositauth SET exported_at = NOW( ) WHERE id = ? LIMIT 1")
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
// UpdateDepositAuthorizationByIDSetFulfilled -- update an item that has been 'fulfilled'
//
func (s *storage) UpdateDepositAuthorizationByIDSetFulfilled(id string, did string) error {

	stmt, err := s.Prepare("UPDATE depositauth SET exported_at = NULL, accepted_at = NOW( ), status = ?, libra_id = ? WHERE id = ? LIMIT 1")
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
// UpdateDepositAuthorizationByIDSetTitle -- update an items title
//
func (s *storage) UpdateDepositAuthorizationByIDSetTitle(id string, title string) error {

	stmt, err := s.Prepare("UPDATE depositauth SET title = ?, updated_at = NOW( ) WHERE id = ? LIMIT 1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(title, id)
	if err != nil {
		return err
	}

	return nil
}

//
// GetFieldMapperList -- get the list of field maps
//
func (s *storage) GetFieldMapperList() ([]*mapper, error) {

	rows, err := s.Query("SELECT field_class, field_name, field_value FROM fieldmapper")
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
		err := rows.Scan(
			&reg.InboundID,
			&reg.ID,
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

//
// end of file
//
