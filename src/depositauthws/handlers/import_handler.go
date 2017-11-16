package handlers

import (
	"depositauthws/authtoken"
	"depositauthws/config"
	"depositauthws/dao"
	"depositauthws/logger"
	"depositauthws/sis"
	"fmt"
	"net/http"
)

//
// ImportHandler -- authorization import request handler
//
func ImportHandler(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("auth")

	// parameters OK ?
	if isEmpty(token) {
		status := http.StatusBadRequest
		encodeImportResponse(w, status, http.StatusText(status), 0, 0, 0, 0)
		return
	}

	// validate the token
	if authtoken.Validate(config.Configuration.AuthTokenEndpoint, token, config.Configuration.Timeout) == false {
		status := http.StatusForbidden
		encodeImportResponse(w, status, http.StatusText(status), 0, 0, 0, 0)
		return
	}

	// get the details ready to be imported
	imports, err := sis.Exchanger.Import()
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		status := http.StatusInternalServerError
		encodeImportResponse(w, status,
			fmt.Sprintf("%s (%s)", http.StatusText(status), err),
			0, 0, 0, 0)
		return
	}

	// if we have nothing to import, bail out
	if len(imports) == 0 {
		status := http.StatusOK
		encodeImportResponse(w, status, http.StatusText(status), 0, 0, 0, 0)
		return
	}

	// import each record and keep track of progress
	newCount := 0
	updateCount := 0
	duplicateCount := 0
	errorCount := 0
	for _, e := range imports {

		// check to see if this record already exists
		matching, err := dao.DB.GetMatchingDepositAuthorization(*e)
		if err != nil {
			logger.Log(fmt.Sprintf("Error querying record; '%s' for (%s/%s/%s/%s)", err, e.ComputingID, e.Degree, e.Plan, e.Title))
			errorCount++
		} else {
			matchCount := len(matching)
			if matchCount == 0 {
				// no match, must be a new record
				rec, err := dao.DB.CreateDepositAuthorization(*e)
				if err != nil {
					logger.Log(fmt.Sprintf("Error inserting record; '%s' for (%s/%s/%s/%s)", err, e.ComputingID, e.Degree, e.Plan, e.Title))
					errorCount++
				} else {
					err = dao.DB.CreateInbound(rec.ID)
					if err != nil {
						logger.Log(fmt.Sprintf("Error creating inbound record; '%s' for (%s/%s/%s/%s)", err, e.ComputingID, e.Degree, e.Plan, e.Title))
						errorCount++
					} else {
						logger.Log(fmt.Sprintf("NEW record (%s/%s/%s/%s)", e.ComputingID, e.Degree, e.Plan, e.Title))
						newCount++
					}
				}
			} else if matchCount == 1 {
				// does the title already match
				if matching[0].Title == e.Title {
					logger.Log(fmt.Sprintf("DUPLICATE record (%s/%s/%s/%s)", e.ComputingID, e.Degree, e.Plan, e.Title))
					duplicateCount++
				} else {
					// titles differ, update the title and mark as an updated item
					err = dao.DB.UpdateDepositAuthorizationByIDSetTitle(matching[0].ID, e.Title)
					if err != nil {
						logger.Log(fmt.Sprintf("Error updating record; '%s' for (%s/%s/%s/%s)", err, e.ComputingID, e.Degree, e.Plan, e.Title))
						errorCount++
					} else {
						err = dao.DB.CreateInbound(matching[0].ID)
						if err != nil {
							logger.Log(fmt.Sprintf("Error creating inbound record; '%s' for (%s/%s/%s/%s)", err, e.ComputingID, e.Degree, e.Plan, e.Title))
							errorCount++
						} else {
							logger.Log(fmt.Sprintf("UPDATE record (%s/%s/%s/%s)", e.ComputingID, e.Degree, e.Plan, e.Title))
							updateCount++
						}
					}
				}
			} else {
				logger.Log(fmt.Sprintf("multiple records exist, ignoring (%s/%s/%s/%s)", e.ComputingID, e.Degree, e.Plan, e.Title))
				errorCount++
			}
		}
	}

	// log summary
	logger.Log(fmt.Sprintf("Import summary: %d new, %d update(s), %d duplicate(s), %d error(s)",
		newCount, updateCount, duplicateCount, errorCount))

	// did we encounter any errors
	if errorCount != 0 {
		status := http.StatusCreated
		encodeImportResponse(w, status,
			fmt.Sprintf("%s (%d errors encountered)", http.StatusText(status), errorCount),
			newCount, updateCount, duplicateCount, errorCount)
		return
	}

	// its all over
	status := http.StatusOK
	encodeImportResponse(w, status, http.StatusText(status), newCount, updateCount, duplicateCount, errorCount)
}

//
// end of file
//
