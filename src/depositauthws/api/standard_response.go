package api

type StandardResponse struct {
   Status        int          `json:"status"`
   Message       string       `json:"message"`
   Details  [] * Registration `json:"details,omitempty"`
}

