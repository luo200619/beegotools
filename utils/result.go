package utils

/**
 * 统一返回格式
 */
type Result struct {
	Err  interface{} `json:"err"`
	Info string      `json:"info"`
	Data interface{} `json:"data"`
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
func AppResult(info string, params ...interface{}) Result {
	length := len(params)
	var result Result
	if length == 0 {
		result = Result{Info: info, Err: true, Data: nil}
	} else if length == 1 {
		result = Result{Info: info, Data: params[0], Err: true}
	} else if length == 2 {
		result = Result{Info: info, Data: params[0], Err: params[1]}
	}
	return result
}
