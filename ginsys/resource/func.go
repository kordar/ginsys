package resource

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/kordar/gocrud"
	response "github.com/kordar/goframework_resp"
)

var Manager = gocrud.NewResourceManager()

func apiAndDriverName(ctx *gin.Context) (string, string) {
	apiName := ctx.Param("apiName")
	driverName := Manager.DriverName(apiName, ctx)
	return apiName, driverName
}

func GetInfo(ctx *gin.Context) {

	apiName, driverName := apiAndDriverName(ctx)
	body := gocrud.NewSearchBody(driverName, ctx)
	if err := ctx.ShouldBind(&body); err != nil {
		response.Error(ctx, err, nil)
		return
	}

	if vo, err := Manager.ReadOne(apiName, body); err == nil {
		response.Success(ctx, "success", vo)
	} else {
		response.Error(ctx, err, nil)
	}

}

func GetList(ctx *gin.Context) {

	apiName, driverName := apiAndDriverName(ctx)
	body := gocrud.NewSearchBody(driverName, ctx)
	if err := ctx.ShouldBind(&body); err != nil {
		response.Error(ctx, err, nil)
		return
	}

	if vo, err := Manager.Read(apiName, body); err == nil {
		response.Data(ctx, "success", vo.Data, vo.Count)
	} else {
		response.Error(ctx, err, nil)
	}
}

func Add(ctx *gin.Context) {

	apiName, driverName := apiAndDriverName(ctx)
	body := gocrud.NewFormBody(driverName, ctx)
	if err := ctx.ShouldBind(&body); err != nil {
		response.Error(ctx, err, nil)
		return
	}

	if obj, err := Manager.Create(apiName, body); err == nil {
		response.Success(ctx, "success", obj)
	} else {
		response.Error(ctx, err, nil)
	}

}

func Update(ctx *gin.Context) {

	apiName, driverName := apiAndDriverName(ctx)
	body := gocrud.NewFormBody(driverName, ctx)
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

	apiName, driverName := apiAndDriverName(ctx)
	body := gocrud.NewRemoveBody(driverName, ctx)
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

	apiName, driverName := apiAndDriverName(ctx)
	body := gocrud.NewEditorBody(driverName, ctx)
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

	apiName := ctx.Param("apiName")
	if configs, err := Manager.Configs(apiName, ctx); err == nil {
		response.Success(ctx, "success", configs)
	} else {
		response.Error(ctx, err, nil)
	}
}

func Upload(ctx *gin.Context) {
	apiName, driverName := apiAndDriverName(ctx)
	body := gocrud.NewFormBody(driverName, ctx)
	if err := ctx.ShouldBind(&body); err != nil {
		response.Error(ctx, err, nil)
		return
	}

	s, err := Manager.GetResourceService(apiName, ctx)
	if err != nil {
		response.Error(ctx, err, nil)
		return
	}

	if uploadResourceService, ok := s.(UploadResourceService); ok {
		if obj, err2 := uploadResourceService.Upload(body); err2 == nil {
			response.Success(ctx, "success", obj)
		} else {
			response.Error(ctx, err2, nil)
		}
	} else {
		response.Error(ctx, errors.New("upload is not supported: please implement the UploadResourceService interface"), nil)
	}

}

func Download(ctx *gin.Context) {

	apiName, driverName := apiAndDriverName(ctx)
	body := gocrud.NewSearchBody(driverName, ctx)
	if err := ctx.ShouldBind(&body); err != nil {
		response.Error(ctx, err, nil)
		return
	}

	s, err := Manager.GetResourceService(apiName, ctx)
	if err != nil {
		response.Error(ctx, err, nil)
		return
	}

	if downloadResourceService, ok := s.(DownloadResourceService); ok {
		if err2 := downloadResourceService.Download(body); err2 == nil {
			// response.Data(ctx, "success", nil, 0)
		} else {
			response.Error(ctx, err2, nil)
		}
	} else {
		response.Error(ctx, errors.New("download is not supported: please implement the DownloadResourceService interface"), nil)
	}
}
