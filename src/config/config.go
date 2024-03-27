package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	DBStringConnection = ""
	DBUser             = ""
	DBPassword         = ""
	DBName             = ""
	Port               = 0
)

func loadEnvFromFile() error {
	file, error := os.Open(".env")
	if error != nil {
		log.Fatal(error)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		switch parts[0] {
		case "DB_USER":
			DBUser = parts[1]
		case "DB_PASSWORD":
			DBPassword = parts[1]
		case "DB_NAME":
			DBName = parts[1]
		case "API_PORT":
			Port, error = strconv.Atoi(parts[1])
			if error != nil {
				Port = 9000
			}
		default:
			log.Fatal(error)
		}
	}

	if error := scanner.Err(); error != nil {
		return error
	}

	return nil
}

func Load() {
	loadEnvFromFile()
	DBStringConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		DBUser,
		DBPassword,
		DBName,
	)
}
