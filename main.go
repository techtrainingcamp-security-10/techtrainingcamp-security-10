package main

import (
	"fmt"
	"techtrainingcamp-security-10/internal/resource"
	"techtrainingcamp-security-10/internal/route"
)

func main() {
	server, err := resource.NewServer()
	if err != nil {
		fmt.Println(err)
		return
	}
	router, err := route.NewRoute(server.Redis.Conn, server.DbR)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = router.Run(":8081")
	if err != nil {
		fmt.Println(err)
	}
	server.Close()
	return
}
