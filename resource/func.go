package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/kordar/gocrud"
	response "github.com/kordar/goframework_resp"
)

var Manager = gocrud.NewResourceManager[interface{}, *gin.Context]()

func GetApiNameAndDriver(ctx *gin.Context) (string, string) {
	apiName := ctx.Param("apiName")
	driverName := Manager.DriverName(apiName)
	return apiName, driverName
}

func GetInfo(ctx *gin.Context) {

	apiName, driverName := GetApiNameAndDriver(ctx)
	body := gocrud.NewSearchBody[interface{}, *gin.Context](driverName, ctx)
	if err := ctx.ShouldBind(&body); err != nil {
		response.Error(ctx, err, nil)
		return
	}

	if vo, err := Manager.SelectOne(apiName, body); err == nil {
		response.Success(ctx, "success", vo)
	} else {
		response.Error(ctx, err, nil)
	}

}

func GetList(ctx *gin.Context) {

	apiName, driverName := GetApiNameAndDriver(ctx)
	body := gocrud.NewSearchBody[interface{}, *gin.Context](driverName, ctx)
	if err := ctx.ShouldBind(&body); err != nil {
		response.Error(ctx, err, nil)
		return
	}

	if vo, err := Manager.Select(apiName, body); err == nil {
		response.Data(ctx, "success", vo.Data, vo.Count)
	} else {
		response.Error(ctx, err, nil)
	}
}

func Add(ctx *gin.Context) {

	apiName, driverName := GetApiNameAndDriver(ctx)
	body := gocrud.NewFormBody[interface{}, *gin.Context](driverName, ctx)
	if err := ctx.ShouldBind(&body); err != nil {
		response.Error(ctx, err, nil)
		return
	}

	if obj, err := Manager.Add(apiName, body); err == nil {
		response.Success(ctx, "success", obj)
	} else {
		response.Error(ctx, err, nil)
	}

}

func Update(ctx *gin.Context) {
	apiName, driverName := GetApiNameAndDriver(ctx)
	body := gocrud.NewFormBody[interface{}, *gin.Context](driverName, ctx)
	if err := ctx.ShouldBind(&body); err != nil {
		response.Error(ctx, err, nil)
		return
	}

	if obj, err := Manager.Update(apiName, body); err == nil {
		response.Success(ctx, "success", obj)
	} else {
		response.Error(ctx, err, nil)
	}
}

func Delete(ctx *gin.Context) {

	apiName, driverName := GetApiNameAndDriver(ctx)
	body := gocrud.NewRemoveBody[interface{}, *gin.Context](driverName, ctx)
	if err := ctx.ShouldBind(&body); err != nil {
		response.Error(ctx, err, nil)
		return
	}

	if err := Manager.Delete(apiName, body); err == nil {
		response.Success(ctx, "success", nil)
	} else {
		response.Error(ctx, err, nil)
	}
}

func Edit(ctx *gin.Context) {

	apiName, driverName := GetApiNameAndDriver(ctx)
	body := gocrud.NewEditorBody[interface{}, *gin.Context](driverName, ctx)

	if err := ctx.ShouldBind(&body); err != nil {
		response.Error(ctx, err, nil)
		return
	}

	if err := Manager.Edit(apiName, body); err == nil {
		response.Success(ctx, "success", nil)
	} else {
		response.Error(ctx, err, nil)
	}

}

func Configs(ctx *gin.Context) {

	apiName, _ := GetApiNameAndDriver(ctx)
	//locale := ctx.GetHeader("Locale")
	if configs, err := Manager.Configs(apiName); err == nil {
		response.Success(ctx, "success", configs)
	} else {
		response.Error(ctx, err, nil)
	}
}
