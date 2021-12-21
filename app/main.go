package main

import (
  "app/router"
  "app/db"
)

func main() {
  db.InitDB()
  router.CreateRouter().Run()
  db.CloseDB()
}
