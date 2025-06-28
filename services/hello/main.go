package main

import "github.com/wrtgvr/go-food-order-ms/services/hello/app"

func main() {
	app := app.NewApp(":9000")

	app.MustRun()
}
