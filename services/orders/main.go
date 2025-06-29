package main

import "github.com/wrtgvr/go-food-order-ms/services/orders/app"

func main() {
	app := app.NewApp(":50051")

	app.MustRun()
}
