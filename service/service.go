/*
 */
package service

import (
	"context"
	"fmt"
	"log"
	"log/registry"
	"net/http"
)

//启动一个新web服务
/*
输入：context, host+port（服务地址）, reg（服务注册信息）
作用：启动一个具有特定地址和功能的服务
*/
func Start(ctx context.Context, host, port string, reg registry.Registration,
	registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, reg.ServiceName, host, port)
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port

	//如果启动时出现错误，则会被打印出来
	go func() {
		log.Println(srv.ListenAndServe())
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		//发生错误，取消serv
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. Press any key to stop. \n", serviceName)
		//等待用户按键
		var s string
		fmt.Scanln(&s)
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
