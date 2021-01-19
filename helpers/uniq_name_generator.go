package helpers

import (
	"fmt"
	"time"
)

func GenerateColumnName(projectName string) string {

	return fmt.Sprintf("%s - Default Column(%q)", projectName, time.Now().Format(time.RFC822))
}
