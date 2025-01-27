package response

func Success(StatusCode int, Data interface{}, Message string) defaultResponse{
	res := defaultResponse{
		StatusCode: StatusCode,
		Data: Data,
		Error: false,
		Message: Message,
	}

	return res;
}

func BadRequest(StatusCode int, Message string) defaultResponse {
	res := defaultResponse{
		StatusCode: StatusCode,
		Data: "No Data",
		Error: true,
		Message: Message,
	}
	return res;
}

func ServerError(Message string) defaultResponse {
	res := defaultResponse{
		StatusCode: 500,
		Data: "No Data",
		Error: true,
		Message: Message,
	}
	return res;
}



type defaultResponse struct{
	Message string `json:"message"`
	StatusCode int `json:"statusCode"`
	Data interface{} `json:"data"`
	Error bool `json:"error"`
}