package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags ExaFileUploadAndDownload
// @Summary 测试链接到服务器
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "an example for breakpoint resume, 断点续传示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/breakpointContinue [post]
func Show(c *gin.Context)  {
	c.JSON(http.StatusOK, "操作成功")
}
