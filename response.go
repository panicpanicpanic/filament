package filament

// Response is a generic slice of results from LIFX API
type Response struct {
	Results []struct {
		ID     string `json:"id"`
		Status string `json:"status"`
		Label  string `json:"label"`
	} `json:"results"`
}
