package utils

import (
	"net/http"
	"log"
	"os"
)
type userError string
func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

type  Handler func(w http.ResponseWriter, request *http.Request) error

func Wrapper(handler Handler) func(w http.ResponseWriter, request *http.Request)  {
	return func(w http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(
					w, //向writer汇报错误
					http.StatusText(http.StatusInternalServerError), //错误描述信息（字符串）
					http.StatusInternalServerError) //系统内部错误
			}

		}()

		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "application/json")
		//返回数据格式是json
		if request.Method=="OPTIONS"{
			log.Printf("OPTIONS")
			return
		}
		// 执行业务代码操作，上面定义的 defer 就是防止 业务代码中出现 panic
		err := handler(w, request)

		// 如果业务代码执行出错
		if err != nil {
			//日志输出错误信息
			log.Printf("Error occurred handling request: %s",err.Error())

			//判断错误类型是否为 自定义错误
			if userErr, ok := err.(userError); ok {
				http.Error(w,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}
			// 判断系统错误的类型
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound  //文件无法找到错误
			case os.IsPermission(err):
				code = http.StatusForbidden // 权限不够错误
			default:
				code = http.StatusInternalServerError //其他错误
			}
			//向writer 中写入错误信息
			http.Error(w,http.StatusText(code), code)
		}

	}

}