package resource

import "github.com/kordar/gocrud"

type UploadResourceService interface {
	Upload(body gocrud.FormBody) (obj any, err error)
}

type DownloadResourceService interface {
	Download(body gocrud.SearchBody) (err error)
}
