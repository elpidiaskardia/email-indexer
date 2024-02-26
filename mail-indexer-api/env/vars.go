package env

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var (
    BaseURL      string
    AuthUsername string
    AuthPassword string
    ContentType  string
    UserAgent    string
)

func loadEnvFile(filename string) (map[string]string, error) {
	env := make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			env[parts[0]] = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return env, nil
}

func init() {
	env, err := loadEnvFile("./.env")
	if err != nil {
		log.Fatal("Error loading the file .env:", err)
	}
    BaseURL = env["BASE_URL"]
    AuthUsername = env["AUTH_USERNAME"]
    AuthPassword = env["AUTH_PASSWORD"]
    ContentType = env["CONTENT_TYPE"]
    UserAgent = env["USER_AGENT"]
}