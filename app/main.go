package main

import (
  "app/router"
  "app/db"
)

func main() {
  db.InitDB()
  defer db.CloseDB()
  router.CreateRouter().Run()
}
