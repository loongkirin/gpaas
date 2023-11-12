package router

import (
	"github.com/gin-gonic/gin"
)

type Entry struct {
}

func (e *Entry) InitAllRouter(router *gin.RouterGroup) {
	sysAuth := SystemAuthorityRouter{}
	sysAuth.InitSystemAuthorityRouter(router)

	sysDesinger := SystemDesignerRouter{}
	sysDesinger.InitSystemDesignerRouter(router)

}
