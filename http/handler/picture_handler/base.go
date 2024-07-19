package picture_handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web/service/picture_service"
	"web/utils/errs"
)

func NewPictureHandler() *PictureHandler {
	return &PictureHandler{
		service: picture_service.NewPictureService(),
	}
}

type PictureHandler struct {
	service *picture_service.PictureService
}

// ApiShowPlay 展示播放按钮
func (ph *PictureHandler) ApiShowPlay(ctx *gin.Context) {
	var response = make(map[string]interface{})
	response["data"] = 1
	ctx.JSON(errs.SucResp(response))
}

// ApiBookNavList 获取栏目展示-顶部Nav栏目
func (ph *PictureHandler) ApiBookNavList(ctx *gin.Context) {
	typeId, _ := strconv.Atoi(ctx.Query("type"))
	response, err := ph.service.FindBookNavList(typeId)
	if err != nil {
		ctx.JSON(errs.ErrResp(err))
		return
	}
	ctx.JSON(errs.SucResp(response))
}
