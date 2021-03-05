package main

import "app/infrastructure/router"

func main() {
	r := router.NewRouter()
	r.Run()
}
