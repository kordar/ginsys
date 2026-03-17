# ginsys

基于 Gin 的一组“开箱即用”增强能力，聚合了：

- 服务启动封装（GinServer）
- Validator 自定义校验与多语言翻译集成
- 统一 JSON 返回结构 + i18n
- 常用中间件（Recovery/CORS/Locale 处理）
- 面向 gocrud 的资源（resource）HTTP 封装

## 安装

```bash
go get github.com/kordar/ginsys@latest
```

Go 版本：`go1.18+`

## 快速开始

下面示例演示：

- 初始化 i18n 配置（从 `language/` 目录加载 ini/toml/yaml）
- 初始化 JSON 返回结构与 i18n Response
- 启用 Validator 翻译器（中英）
- 注入中间件（Recovery、Locale 转换、CORS）
- 添加自定义校验

```go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/kordar/ginsys"
	"github.com/kordar/ginsys/middleware"
	"github.com/kordar/ginsys/resp"
	"github.com/kordar/ginsys/util"

	"github.com/kordar/gocfg"
	"github.com/kordar/goframework_resp"
	"github.com/kordar/gotrans"
)

type DemoForm struct {
	Name string `json:"name" binding:"required,tom"`
}

type ValidTom struct{}

func (v ValidTom) DefaultTpl() (tpl string, override bool) { return "", true }
func (v ValidTom) Tpl() (section string, key string)       { return "dictionary", "ttt" }
func (v ValidTom) I18n(fe validator.FieldError, locale string) []string {
	return nil
}
func (v ValidTom) Tag() string { return "tom" }
func (v ValidTom) Valid(fl validator.FieldLevel) bool {
	fmt.Println("value:", fl.Field().String())
	return false
}

func router(r *gin.Engine) *gin.Engine {
	r.GET("/hello", func(ctx *gin.Context) {
		goframework_resp.Success(ctx, "success", nil)
	})
	r.POST("/tt", func(ctx *gin.Context) {
		var f DemoForm
		if err := util.DefaultGetValidParams(ctx, &f); err != nil {
			goframework_resp.Error(ctx, err, nil)
			return
		}
		goframework_resp.Success(ctx, "success", nil)
	})
	return r
}

func main() {
	gocfg.InitConfigWithParentDir("language", "ini")

	resp.InitJsonResp001()
	resp.InitI18nResponse()

	ginsys.NewGinServer().
		Router(router).
		OpenValidateAndTranslations(gotrans.NewEnTranslation(), gotrans.NewZhTranslation()).
		Middleware(
			middleware.RecoveryMiddleware(),
			resp.TransLocaleMiddleware("Locale", "zh"),
		).
		AddValidate(ValidTom{}).
		Middleware(middleware.CorsMiddleware()).
		StartD("0.0.0.0", "9099")
}
```

请求示例（指定语言）：

```bash
curl -X POST http://localhost:9099/tt \
  -H 'Content-Type: application/json' \
  -H 'Locale: zh-CN' \
  -d '{"name":"abc"}'
```

## 目录结构

- `ginserver.go`：GinServer，封装 Gin Engine 的创建/中间件/路由/启动，以及校验与翻译集成
- `resp/`：统一响应与 i18n（可通过 Header Locale 选择语言）
- `middleware/`：Recovery、CORS 等
- `resource/`：面向 `gocrud.ResourceManager` 的 REST 风格处理函数封装
- `util/`：常用工具（绑定校验触发、读取 i18n 配置）
- `tenant/`：租户相关全局配置

## 响应与国际化（resp）

常用初始化组合：

- `resp.InitJsonResp001()`：注册统一 JSON 输出结构
- `resp.InitI18nResponse()`：为 success/error/valid 等响应注册 i18n 输出逻辑
- `resp.TransLocaleMiddleware(key, defaultLocale)`：将外部 locale（如 `zh-CN`）映射为翻译器实际注册的 locale（默认映射：`zh-CN -> zh`）

默认读取 Header：`Locale`，默认语言：`en`。可用：

- `resp.SetHeaderKey("Locale")`
- `resp.SetDefaultLocale("en")`
- `resp.SetTransLocaleMapValue("zh-CN", "zh")`

## 校验与翻译

- `NewGinServer/NewGinEngineServer` 会将 Gin 内置 validator 引擎注入到 `govalidator`
- `OpenValidateAndTranslations(...)` 会初始化 `gotrans` 的翻译器集合
- `AddValidate(...)` 支持注册自定义校验，并在翻译器已初始化时，将校验的 i18n 文案绑定到 validator

## gocrud 资源封装（resource）

`resource` 包提供了一组可直接挂到 Gin 路由上的处理函数（按 `:apiName` 分发），内部使用 `resource.Manager`（`gocrud.ResourceManager`）执行 CRUD。

处理函数：

- `resource.GetInfo`
- `resource.GetList`
- `resource.Add`
- `resource.Update`
- `resource.Delete`
- `resource.Edit`
- `resource.Configs`
- `resource.Upload`
- `resource.Download`

你需要按 `gocrud` 的方式提前为 `resource.Manager` 配置资源、驱动与执行器。

其中 `Upload/Download` 依赖资源服务实现以下可选接口：

- `resource.UploadResourceService`：`Upload(body gocrud.FormBody) (obj any, err error)`
- `resource.DownloadResourceService`：`Download(body gocrud.SearchBody) error`

当资源服务未实现对应接口时，接口会返回明确错误提示：

- upload：`upload is not supported: please implement the UploadResourceService interface`
- download：`download is not supported: please implement the DownloadResourceService interface`
