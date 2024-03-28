package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DBStringConnection = ""
	Port               = 0
	SecretKey          []byte
)

func Load() {
	var error error

	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	Port, error = strconv.Atoi(os.Getenv("API_PORT"))
	if error != nil {
		Port = 9000
	}
	DBStringConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}
