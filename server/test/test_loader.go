package main

import (
	"db"
	"fmt"
	"log"
	"schema"
)

func main(){
  dbHandler, err := db.GetDBConn()
  if err != nil {
    log.Fatalf(err.Error())
  }
  application := &schema.Application{
    CompanyName: "Google",
    Role: "SWE",
    ApplicationURL: "google.com",
    Status: "Applied",
  }
  err = dbHandler.AddApplication(application)
  if err != nil {
    log.Fatalf(err.Error())
  }
  fmt.Println("Document added into the database")
}
