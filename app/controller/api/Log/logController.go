package Log

import (
	"fitness/pkg/util/response"
	"github.com/gin-gonic/gin"
)

func LogTest(c *gin.Context) {


	response.Success("handle ok")
}
