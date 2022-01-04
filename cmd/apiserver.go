package main

import (
	"github.com/dhope-nagesh/gin_todo/pkg/service/router"
)

func main() {
	r := router.GetRouter()
	r.Run()
}
