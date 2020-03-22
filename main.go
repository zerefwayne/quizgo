package main

import (
	"fmt"
	"github.com/zerefwayne/quizgo/types"
	"log"
)

func main() {
	driver, err := types.NewDriver()

	if err != nil {
		log.Fatal(err)
		return
	} else {
		driver.Start()

		fmt.Printf("Game Over\n\n")
		restart := ""
		fmt.Printf("Restart (y/n): ")
		fmt.Scanf("%s", &restart)

		if restart == "y" {
			driver.Start()
		}

	}
}
