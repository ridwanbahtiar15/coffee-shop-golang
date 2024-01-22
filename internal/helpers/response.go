package helpers

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type Meta struct {
	Page      int    `json:"page,omitempty"`
	TotalData int    `json:"total_data,omitempty"`
	NextPage  string `json:"next,omitempty"`
	PrevPage  string `json:"prev,omitempty"`
}

func GetPagination(resultPage int, totalData int, nextPage string, prevPage string) Meta {
	return Meta{
		Page:      resultPage,
		TotalData: totalData,
		NextPage:  nextPage,
		PrevPage:  prevPage,
	}
}

func GetResponse(message string, data interface{}, meta *Meta) Response {
	return Response{
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}