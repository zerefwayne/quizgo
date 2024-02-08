package main

import (
	"fmt"
	"log"
	"quizgo/database"
	"quizgo/types"
)

func main() {
	_, err := database.Open()
	if err != nil {
		log.Fatal("ERROR - Failed to open Database:", err)
		return
	}

	driver, err := types.NewDriver()
	if err != nil {
		log.Fatal("ERROR - Failed to create new driver:", err)
		return
	} else {

		for {
			score, err := driver.Start()
			if err != nil {
				log.Fatal("ERROR - Driver failed to start")
				return
			}

			fmt.Printf("Game Over\n\n")
			restart := ""
			fmt.Printf("Restart (y/n): ")
			fmt.Scanf("%s", &restart)

			if restart != "y" {
				fmt.Println("Final Score:", score)
				break
			}
		}

	}
}
