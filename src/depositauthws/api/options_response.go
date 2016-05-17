package api

type OptionsResponse struct {
    Status        int          `json:"status"`
    Message       string       `json:"message"`
    Options     * Options      `json:"options,omitempty"`
}

