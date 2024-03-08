package main

import (
	"api"
	"flag"
)

var addr = flag.String("addr", ":5990", "http listener port")

func main(){
  flag.Parse()

  server := api.NewServer(*addr)
  server.Start()

}
