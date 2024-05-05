package initialisers

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnv(){
		// Load .env file
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")
		}
}