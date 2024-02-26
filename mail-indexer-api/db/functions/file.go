package functions

import (
	"Api_Go/models"
	"bufio"
	"bytes"
	"io/ioutil"
	"reflect"
	"strings"
)


func assignField(emailData *models.EmailData, key, value string) {
	elem := reflect.ValueOf(emailData).Elem()
	field := elem.FieldByName(key)

	if field.IsValid() && field.CanSet() && value != "" {
		field.SetString(value)
	}
}

func ProcessFile(path string) (models.EmailData, error) {
	data, err := ioutil.ReadFile(path)
    if err != nil {
		return models.EmailData{}, err
	}
	emailData := models.NewEmail() 
	var remainingLines []string

    scanner := bufio.NewScanner(bytes.NewReader(data))
	buf := make([]byte, 0, 64*2048)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)
	isContent := false
	for scanner.Scan() {
		line := scanner.Text()
		if isContent {
			remainingLines = append(remainingLines, line)
		} else {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) < 2 {
				isContent = true
			} else {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				assignField(&emailData, key, value)
			}
		}
	}

	emailData.Content = strings.Join(remainingLines, "\n")

	if err := scanner.Err(); err != nil {
		return models.EmailData{}, err
	}

	return emailData, nil
}

