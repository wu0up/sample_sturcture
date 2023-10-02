package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, response interface{}, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	switch resp := response.(type) {
	case []string:
		c.JSON(http.StatusOK, resp)
	case []map[string]interface{}:
		c.JSON(http.StatusOK, resp)
	default:
		// 处理未知类型的响应，或者可以返回适当的错误响应
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown response type"})
	}

}
