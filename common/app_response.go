package common

type succesRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging"`
	Filter interface{} `json:"filter"`
}

func NewSuccessResponse(data, paging, filter interface{}) *succesRes {
	return &succesRes{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessResponse(data interface{}) *succesRes {
	return &succesRes{Data: data, Paging: nil, Filter: nil}
}
