package dao

import (
	"github.com/uvalib/deposit-auth-ws/depositauthws/api"
)

// our storage interface
type Storage interface {
	Check() error
	GetInbound(after string) ([]*api.Authorization, error)
	GetMatchingDepositAuthorization(e api.Authorization) ([]*api.Authorization, error)
	GetDepositAuthorizationByID(id string) ([]*api.Authorization, error)
	SearchDepositAuthorizationByCid(cid string) ([]*api.Authorization, error)
	SearchDepositAuthorizationByCreateDate(createdAt string) ([]*api.Authorization, error)
	SearchDepositAuthorizationByExportDate(exportedAt string) ([]*api.Authorization, error)
	CreateInbound(authID string) error
	CreateDepositAuthorization(reg api.Authorization) (*api.Authorization, error)
	DeleteDepositAuthorizationByID(id string) (int64, error)
	GetDepositAuthorizationForExport() ([]*api.Authorization, error)
	UpdateExportedDepositAuthorization(exports []*api.Authorization) error
	UpdateDepositAuthorizationByIDSetFulfilled(id string, did string) error
	UpdateDepositAuthorizationByIDSetTitle(id string, title string) error
	GetFieldMapperList() ([]*mapper, error)
	//Destroy() error
}

type mapper struct {
	FieldClass  string
	FieldSource string
	FieldMapped string
}

// our singleton store
var Store Storage

// our factory
func NewDatastore() error {
	var err error
	// mock implementation here
	Store, err = newDBStore()
	return err
}

//
// end of file
//
