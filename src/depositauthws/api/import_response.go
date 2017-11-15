package api

//
// ImportResponse -- response to the import request
//
type ImportResponse struct {
   Status         int    `json:"status"`
   Message        string `json:"message"`
   NewCount       int    `json:"new_count"`
   UpdatedCount   int    `json:"update_count"`
   DuplicateCount int    `json:"duplicate_count"`
   ErrorCount     int    `json:"error_count"`
}

//
// end of file
//
