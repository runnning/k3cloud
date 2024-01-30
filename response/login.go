package response

type LoginResponse struct {
	Message           string `json:"Message"`
	MessageCode       string `json:"MessageCode"`
	LoginResultType   int    `json:"LoginResultType"`
	Context           any    `json:"Context"`
	KDSVCSessionID    string `json:"KDSVCSessionId"`
	FormID            any    `json:"FormId"`
	RedirectFormParam any    `json:"RedirectFormParam"`
	FormInputObject   any    `json:"FormInputObject"`
	ErrorStackTrace   any    `json:"ErrorStackTrace"`
	Lcid              int    `json:"Lcid"`
	AccessToken       any    `json:"AccessToken"`
	KdAccessResult    any    `json:"KdAccessResult"`
	IsSuccessByAPI    bool   `json:"IsSuccessByAPI"`
}
