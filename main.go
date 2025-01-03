package main

import (
	"log"
	Route "server/routes"
	Services "server/services"

	"github.com/joho/godotenv"
)

//Conv "server/utils"

func main() {
	//_log("get_fiber 0 1")
	log.Println()
	godotenv.Load()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	//Services.Handler_fiber(test,)

	Services.MongoDB()
	Route.Route_path()

}
func ENV() {
	godotenv.Load()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	//ip := os.Getenv("acu_ip")

	/*dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")*/

}

