package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"path"
)

/**
 * [Debug 调试函数]
 * @作者    como
 * @时间    2019-06-25
 * @版权    思智捷管理系统
 * @版本    1.0.0
 * @param {[type]}   info interface{}    [description]
 * @param {[type]}   v    ...interface{} [description]
 */
func Debug(info interface{}, v ...interface{}) {
	isMode := beego.AppConfig.String("runmode") == "dev"
	if isMode {
		logs.Debug(info, v)
	}
}

/**
 * [getFileExt 获取文件名后缀]
 * @作者     como
 * @时间     2019-06-26
 * @版权     思智捷管理系统
 * @版本     1.0.0
 * @param  {[type]}   fileName string        [description]
 * @return {[type]}            [description]
 */
func GetFileExt(fileName string) string {
	return path.Ext(fileName)
}
