package enviroments

import (
	"api-estadia-express/init/logger"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatal("error to load .env", err)
	}
}
