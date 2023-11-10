package resource

import (
	"github.com/gin-gonic/gin"
	response "github.com/kordar/ginsys/resp"
	crud "github.com/kordar/gocrud"
)

var Manager = crud.NewResourceManager()

func GetInfo(ctx *gin.Context) {

	body := crud.NewSearchBody(ctx)
	if err := ctx.ShouldBind(&body); err != nil {
		response.Error(ctx, err, nil)
		return
	}

	apiName := ctx.Param("apiName")
	if vo, err := Manager.SelectOne(apiName, body); err == nil {
		response.Success(ctx, "success", vo)
	} else {
		response.Error(ctx, err, nil)
	}

}

func GetList(ctx *gin.Context) {
	body := crud.NewSearchBody(ctx)
	if err := ctx.ShouldBind(&body); err != nil {
		response.Error(ctx, err, nil)
		return
	}

	apiName := ctx.Param("apiName")
	if vo, err := Manager.Select(apiName, body); err == nil {
		response.Data(ctx, "success", vo.Data, vo.Count)
	} else {
		response.Error(ctx, err, nil)
	}
}

func Add(ctx *gin.Context) {
	body := crud.NewFormBody(ctx)
	if err := ctx.ShouldBind(&body); err != nil {
		response.Error(ctx, err, nil)
		return
	}

	apiName := ctx.Param("apiName")
	if obj, err := Manager.Add(apiName, body); err == nil {
		response.Success(ctx, "success", obj)
	} else {
		response.Error(ctx, err, nil)
	}

}

func Update(ctx *gin.Context) {
	body := crud.NewFormBody(ctx)
	if err := ctx.ShouldBind(&body); err != nil {
		response.Error(ctx, err, nil)
		return
	}
	apiName := ctx.Param("apiName")
	if obj, err := Manager.Update(apiName, body); err == nil {
		response.Success(ctx, "success", obj)
	} else {
		response.Error(ctx, err, nil)
	}
}

func Delete(ctx *gin.Context) {
	body := crud.NewRemoveBody(ctx)
	if err := ctx.ShouldBind(&body); err != nil {
		response.Error(ctx, err, nil)
		return
	}

	apiName := ctx.Param("apiName")
	if err := Manager.Delete(apiName, body); err == nil {
		response.Success(ctx, "success", nil)
	} else {
		response.Error(ctx, err, nil)
	}
}

func Edit(ctx *gin.Context) {
	body := crud.NewEditorBody(ctx)
	if err := ctx.ShouldBind(&body); err != nil {
		response.Error(ctx, err, nil)
		return
	}

	apiName := ctx.Param("apiName")
	if err := Manager.Edit(apiName, body); err == nil {
		response.Success(ctx, "success", nil)
	} else {
		response.Error(ctx, err, nil)
	}

}

func Configs(ctx *gin.Context) {
	locale := ctx.GetHeader("Locale")
	apiName := ctx.Param("apiName")
	if configs, err := Manager.Configs(apiName, locale); err == nil {
		response.Success(ctx, "success", configs)
	} else {
		response.Error(ctx, err, nil)
	}
}
