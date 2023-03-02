package contract

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	CONTRACT_ADDRESS string
	BLOCKCHAIN_WSS   string
	PRIVATE_KEY      string
	BLOCKCHAIN_URL   string
	EVENT_SIGNATURE  string
}

func LoadENV() Config {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}
	config := Config{}
	config.CONTRACT_ADDRESS = os.Getenv("CONTRACT_ADDRESS")
	config.BLOCKCHAIN_WSS = os.Getenv("BLOCKCHAIN_WSS")
	config.PRIVATE_KEY = os.Getenv("PRIVATE_KEY")
	config.BLOCKCHAIN_URL = os.Getenv("BLOCKCHAIN_URL")
	config.EVENT_SIGNATURE = os.Getenv("EVENT_SIGNATURE")

	return config
}
