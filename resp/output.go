package resp

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type IOutput interface {
	Type() string // "data" | "browser"
	Header() map[string]string
	Params() interface{}
	Data() []byte
}

// OutputData 直接导出二进制
type OutputData struct {
	bytes  []byte
	params interface{}
}

func NewOutputData(bytes []byte, params interface{}) *OutputData {
	return &OutputData{bytes: bytes, params: params}
}

func (o OutputData) Type() string {
	return "data"
}

func (o OutputData) Header() map[string]string {
	return map[string]string{}
}

func (o OutputData) Params() interface{} {
	return o.params
}

func (o OutputData) Data() []byte {
	return o.bytes
}

// BaseToBrowser 导出到浏览器
type BaseToBrowser struct {
	bytes     []byte
	filename  string
	extension string
	header    map[string]string
}

func (o BaseToBrowser) Type() string {
	//TODO implement me
	return "browser"
}

func (o BaseToBrowser) Header() map[string]string {
	//TODO implement me
	if o.filename == "" {
		o.header["Content-Disposition"] = "attachment; filename=" + url.QueryEscape(o.createFileName())
	} else {
		o.header["Content-Disposition"] = fmt.Sprintf("attachment; filename=%s.%s", o.filename, o.extension)
	}
	return o.header
}

func (o BaseToBrowser) Params() interface{} {
	//TODO implement me
	return nil
}

func (o BaseToBrowser) Data() []byte {
	//TODO implement me
	return o.bytes
}

func (o BaseToBrowser) createFileName() string {
	name := time.Now().Format("2006-01-02-15-04-05")
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("excle-%v-%v."+o.extension, name, rand.Int63n(time.Now().Unix()))
}

// OutputExcelToBrowser 导出excel到浏览器
type OutputExcelToBrowser struct {
	BaseToBrowser
}

func NewOutputExcelToBrowser(bytes []byte, filename string,
	ext string, header map[string]string) *OutputExcelToBrowser {
	if header == nil {
		header = map[string]string{}
	}
	if header["Content-Type"] == "" {
		header["Content-Type"] = "application/vnd.ms-excel;charset=utf8"
	}
	// ---------------
	return &OutputExcelToBrowser{BaseToBrowser: BaseToBrowser{
		bytes:     bytes,
		filename:  filename,
		extension: ext,
		header:    header,
	}}
}
