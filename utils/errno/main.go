package errno

// 定义错误码
type errors struct {
	Code     int
	HttpCode int
	Message  string
}

//错误级别
// 没错误 0
// 系统错误 1
// 数据库错误 2

//模块编号
// 任务 01
// 远端存储 02
// Yserver服务端 03
// 后台计划任务 04
// Yserver上传 05

var all = map[int]errors{
	0: errors{
		Code:     0,
		HttpCode: 200,
		Message:  "成功",
	},
}

func New(code int, err error) *errors {
	if code != 0 || err != nil {
		newerr := new(errors)
		newerr.Code = code
		newerr.Message = err.Error()
		newerr.HttpCode = all[code].HttpCode
		return newerr
	}
	return nil
}
