package controllers

import (
	_ "encoding/json"
	"path"
	"strings"

	"github.com/astaxie/beego"
	"github.com/luo200619/beegotools/utils"
)

/**
 * 基类控制器
 */
type Base struct {
	beego.Controller
}

/**
 * [func 显示模版]
 * @作者     como
 * @时间     2019-06-22
 * @版权     思智捷管理系统
 * @版本     1.0.0
 * @param  {[type]}   this *BaseController) Display(fileName string [description]
 * @return {[type]}        [description]
 */
func (this *Base) Display(viewPath string, fileName ...string) {
	length := len(fileName)
	controllerName, actionName := this.GetControllerAndAction()
	var templatehtml string
	if length == 0 {
		templatehtml = path.Join(viewPath, strings.ToLower(controllerName)+beego.AppConfig.String("tpl_split")+strings.ToLower(actionName)) + beego.AppConfig.String("tpl_suffix")
	} else if length == 1 {
		templatehtml = path.Join(viewPath, fileName[0]) + beego.AppConfig.String("tpl_suffix")
	} else if length == 2 {
		templatehtml = path.Join(viewPath, fileName[0], beego.AppConfig.String("tpl_split"), fileName[1]) + beego.AppConfig.String("tpl_suffix")
	} else {
		templatehtml = ""
	}
	this.TplName = templatehtml
}

/**
 * [func 输出数据]
 * @作者     como
 * @时间     2019-06-22
 * @版权     思智捷管理系统
 * @版本     1.0.0
 * @param  {[type]}   this *BaseController) Assign(key string,val interface{} [description]
 * @return {[type]}        [description]
 */
func (this *Base) Assign(key string, val interface{}) {
	this.Data[key] = val
}

/**
 * [func ajax请求返回]
 * @作者     como
 * @时间     2019-06-23
 * @版权     思智捷管理系统
 * @版本     1.0.0
 * @param  {[type]}   this *BaseController) ajaxReturn(data interface{} [description]
 * @return {[type]}        [description]
 */
func (this *Base) AjaxReturn(data interface{}) {
	this.Assign("json", data)
	this.ServeJSON()
	this.StopRun()
}

/**
 * [func 格式化统一返回]
 * @作者     como
 * @时间     2019-06-23
 * @版权     思智捷管理系统
 * @版本     1.0.0
 * @param  {[type]}   this *Base)        AppResult(info string,params ...interface{} [description]
 * @return {[type]}        [description]
 */
func (this *Base) AppResult(info string, params ...interface{}) utils.Result {
	return utils.AppResult(info, params)
}

/**
 * [func 信息调用工具]
 * @作者     como
 * @时间     2019-06-23
 * @版权     思智捷管理系统
 * @版本     1.0.0
 * @param  {[type]}   this *Base)        Debug(v ...interface{} [description]
 * @return {[type]}        [description]
 */
func (this *Base) Debug(info interface{}, v ...interface{}) {
	utils.Debug(info, v)
}

/**
 * [func 自定义404页面]
 * @作者     como
 * @时间     2019-06-23
 * @版权     思智捷管理系统
 * @版本     1.0.0
 * @param  {[type]}   this *Base)        Error404( [description]
 * @return {[type]}        [description]
 */
func (this *Base) Error404() {
	this.Ctx.WriteString("出错啦,找不到访问了")
	this.StopRun()
}

/**
 * [func 自定义500页面]
 * @作者     como
 * @时间     2019-06-23
 * @版权     思智捷管理系统
 * @版本     1.0.0
 * @param  {[type]}   this *Base)        Error500( [description]
 * @return {[type]}        [description]
 */
func (this *Base) Error500() {
	this.Ctx.WriteString("服务器内部出错")
	this.StopRun()
}

/**
 * [func 文件上传示例]
 * @作者     como
 * @时间     2019-06-26
 * @版权     思智捷管理系统
 * @版本     1.0.0
 * @param  {[type]}   this *Base)        UploadManager( [description]
 * @return {[type]}        [description]
 */
func (this *Base) UploadManager(inputName string, conf utils.UploadConfig) utils.Result {
	file, h, err := this.Ctx.Request.FormFile(inputName)
	if err != nil {
		return utils.AppResult(err.Error())
	}
	defer file.Close()
	result := utils.UploadHandler(file, h, inputName, conf)
	return result
}
