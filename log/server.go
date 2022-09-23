/*
实现log的业务逻辑
*/
package log

import (
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

type filelog string

//将数据写入到文件（filelog）内
/*
输入：文件名（string）
输出：写入的数据长度（int）、错误（error）
*/
func (fl filelog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

//在指定路径的日志里写入新的日志消息
/*
输入：日志路径（string）
作用：向指定路径的日志里写入特定的格式化的日志消息
*/
func Run(destination string) {
	log = stlog.New(filelog(destination), "[go] - ", stlog.LstdFlags)
}

//注册Handler
/*
输入：http请求
作用：处理特定路径（/log）的httpPost请求
*/
func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := ioutil.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

//打印消息
func write(message string) {
	log.Printf("%v\n", message)
}
