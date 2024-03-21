package response

type K3Response struct {
	Result *Result `json:"Result"`
}

type Result struct {
	Status         K3Status          `json:"ResponseStatus"`
	ID             int               `json:"Id"`
	Number         string            `json:"Number"`
	NeedReturnData []*NeedReturnData `json:"NeedReturnData"`
}
type NeedReturnData struct {
	FVOUCHERGROUPNO string `json:"FVOUCHERGROUPNO"`
	FVOUCHERID      string `json:"FVOUCHERID"`
}

type K3Status struct {
	ErrorCode       int                `json:"ErrorCode"`
	IsSuccess       bool               `json:"IsSuccess"`
	Errors          []*Errors          `json:"Errors"`
	SuccessEntities []*SuccessEntities `json:"SuccessEntitys"`
	SuccessMessages []any              `json:"SuccessMessages"`
	MsgCode         int                `json:"MsgCode"`
}

type Errors struct {
	FieldName any    `json:"FieldName"`
	Message   string `json:"Message"`
	DIndex    int    `json:"DIndex"`
}

type SuccessEntities struct {
	ID     int    `json:"Id"`
	Number string `json:"Number"`
	DIndex int    `json:"DIndex"`
}
