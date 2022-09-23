package main

import (
	"context"
	"fmt"
	"log"
	"log/registry"
	"net/http"
)

func main() {
	registry.SetupRegistryService()
	http.Handle("/services", &registry.RegistryService{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//配置server
	var srv http.Server
	srv.Addr = registry.ServerPort

	//goroutine
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("Registry service started, press any key to quit.\n")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("shutting down registry service")
}
