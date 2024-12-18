package main

import (
	"github.com/Rail-KH/HTTP-Calculator/internal/application"
)

func main() {
	app := application.New()
	app.RunServer()
}
