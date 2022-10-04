package main

import (
	"flag"
	"fmt"
	"hexagonal_boilerplate/app"
	"hexagonal_boilerplate/shared/driver"
)

func main() {
	appMap := map[string]func() driver.RegistryContract{
		"intl":   app.NewIntl(),
		"extl":   app.NewExtl(),
		"worker": app.NewWorker(),
	}
	flag.Parse()

	app, exist := appMap[flag.Arg(0)]
	if exist {
		driver.Run(app())
	} else {
		fmt.Println("You may try 'go run main.go <app_name>' :")
		for appName := range appMap {
			fmt.Printf(" - %s\n", appName)
		}
	}

}
