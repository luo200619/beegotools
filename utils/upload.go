package utils

import (
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/luo200619/tools"
)

//类型接口
type Sizer interface {
	Size() int64
}

/**
 * 文件上传配置类
 */
type UploadConfig struct {
	Ext      map[string]bool //上传后缀限制
	Size     int64           //上传大小限制
	SavePath string          //文件保存路径
	SaveName string          //文件保存名字
}

/**
 * [UploadConf 获取文件上传配置]
 * @作者 como
 * @时间 2019-06-26
 * @版权 思智捷管理系统
 * @版本 1.0.0
 */
func GetUploadConf() UploadConfig {
	confExt := beego.AppConfig.String("upload_ext")
	var limitMap = make(map[string]bool)
	if confExt == "" {
		limitMap = map[string]bool{".jpg": true, ".png": true, ".bmp": true, ".jpeg": true, ".gif": true}
	} else {
		limitMap = limitCusConfigExt(confExt)
	}
	size := beego.AppConfig.DefaultInt64("upload_size", 2)
	savepath := beego.AppConfig.String("upload_savepath")
	if savepath == "" {
		savepath = "/static/uploads/"
	}
	savename := beego.AppConfig.String("upload_savename")
	//获取到当前时间豪秒的时间戳 + 六位随机数 的字符串作为名字
	curTimeStr := strconv.FormatInt(time.Now().UnixNano()+int64(tools.Mt_rand(100000, 999999)), 10)
	switch savename {
	case "time":
		savename = strconv.FormatInt(time.Now().UnixNano(), 10)
	case "md5":
		savename = tools.Md5([]byte(curTimeStr))
	default:
		savename = tools.Md5([]byte(curTimeStr))
	}
	config := UploadConfig{Ext: limitMap, Size: size, SavePath: savepath, SaveName: savename}
	return config
}

/**
 * [limitCusConfigExt 自定义限制的文件后缀]
 * @作者     como
 * @时间     2019-06-28
 * @版权     思智捷管理系统
 * @版本     1.0.0
 * @param  {[type]}   confExt string        [description]
 * @return {[type]}           [description]
 */
func limitCusConfigExt(confExt string) map[string]bool {
	conf := tools.Explode(",", confExt)
	limitConfExt := make(map[string]bool)
	for _, item := range conf {
		limitConfExt[item] = true
	}
	return limitConfExt
}

/**
 * [UploadHandler  文件上传处理器]
 * @作者    como
 * @时间    2019-06-27
 * @版权    思智捷管理系统
 * @版本    1.0.0
 * @param {[type]}   c         *beego.Controller [description]
 * @param {[type]}   inputName string            [description]
 * @param {[type]}   conf      UploadConfig      [description]
 */
func UploadHandler(file multipart.File, h *multipart.FileHeader, inputName string, conf UploadConfig) Result {
	var result Result
	result = _uploadValidate(file, h, conf)
	if isTrue, ok := result.Err.(bool); ok && isTrue {
		return result
	}
	result = _uploadSaveFileHandler(file, h, inputName, conf)
	return result
}

/**
 * [_uploadSaveFileHandler 实现文件保存功能]
 * @作者     como
 * @时间     2019-06-27
 * @版权     思智捷管理系统
 * @版本     1.0.0
 * @param  {[type]}   h    *multipart.FileHeader [description]
 * @param  {[type]}   conf UploadConfig          [description]
 * @return {[type]}        [description]
 */
func _uploadSaveFileHandler(file multipart.File, h *multipart.FileHeader, inputName string, conf UploadConfig) Result {
	var result Result
	ext := GetFileExt(h.Filename)
	err := os.MkdirAll("."+conf.SavePath, 0777)
	if err != nil {
		return AppResult(err.Error())
	}
	f, err1 := os.OpenFile("."+conf.SavePath+conf.SaveName+ext, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err1 != nil {
		return AppResult(err1.Error())
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		result = AppResult(err.Error())
	} else {
		data := map[string]string{
			"savepath": conf.SavePath,
			"savename": conf.SaveName + ext,
		}
		result = AppResult("SUCCESS", data, false)
	}

	return result
}

/**
 * [_uploadValidate 文件上传验证]
 * @作者     como
 * @时间     2019-06-27
 * @版权     思智捷管理系统
 * @版本     1.0.0
 * @param  {[type]}   file multipart.File [description]
 * @param  {[type]}   conf UploadConfig   [description]
 * @return {[type]}        [description]
 */
func _uploadValidate(file multipart.File, h *multipart.FileHeader, conf UploadConfig) Result {
	var result Result
	if fileSizer, ok := file.(Sizer); ok {
		size := fileSizer.Size()
		if size > conf.Size*1024*1024 {
			result = AppResult("上传的文件超过大小的限制")
			return result
		}
		ext := GetFileExt(h.Filename)
		if _, ok := conf.Ext[ext]; !ok {
			result = AppResult("上传的文件类型不合法")
			return result
		}
	} else {
		result = AppResult("未知错误，无法获取文件大小")
		return result
	}
	return result
}
