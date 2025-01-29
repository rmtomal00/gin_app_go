package response

import "github.com/gin-gonic/gin"

func Success(ctx *gin.Context, StatusCode int, Data interface{}, Message string){
	res := defaultResponse{
		StatusCode: StatusCode,
		Data: Data,
		Error: false,
		Message: Message,
	}

	ctx.JSON(StatusCode, res)
	ctx.Abort()
}

func BadRequest(ctx *gin.Context, StatusCode int, Message string) {
	res := defaultResponse{
		StatusCode: StatusCode,
		Data: "No Data",
		Error: true,
		Message: Message,
	}
	
	ctx.JSON(StatusCode, res)
	ctx.Abort()
}

func ServerError(ctx *gin.Context, Message string) {

	res := defaultResponse{
		StatusCode: 500,
		Data: "No Data",
		Error: true,
		Message: Message,
	}
	ctx.JSON(500, res)
	ctx.Abort()
	
}



type defaultResponse struct{
	Message string `json:"message"`
	StatusCode int `json:"statusCode"`
	Data interface{} `json:"data"`
	Error bool `json:"error"`
}