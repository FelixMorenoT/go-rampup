package main

import "github.com/FelixMorenoT/go-rampup/internal/app"

func main() {
	app := app.NewApp()

	app.Start("8080")
}
