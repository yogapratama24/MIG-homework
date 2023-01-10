package main

import (
	"fmt"
	"homework_mitramas/route"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error read env file with err: %s", err)
	}
	e := route.Init()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("APP_PORT"))))
}
