package main

import (
	"github.com/girakdev/girack-backend/controller"
	_ "github.com/lib/pq"
)

func main() {
	router := controller.Router()
	router.Run()
}
