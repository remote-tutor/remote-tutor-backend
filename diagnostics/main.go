package diagnostics

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func getFile(filePath string, errTime time.Time) (*os.File, error) {
	writeToMainDiagnosticsFile(filePath, errTime)
	filePath = fmt.Sprintf("diagnostics/%s", filePath)
	return os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}

func writeToMainDiagnosticsFile(filePath string, errTime time.Time) {
	file, err := os.OpenFile("diagnostics/main.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		_, _ = file.Write([]byte("Time:\t" + errTime.Format(time.ANSIC) + "\n"))
		_, _ = file.Write([]byte("Details In:\t" + filePath + "\n\n\n"))
		_ = file.Close()
	}
}

func WriteToSpecificDiagnosticsFile(filepath, errorType string, err error, value interface{}) {
	// value parameter is the value causing the error
	// errorType could be thought of as the method name OR the type of error (Create, Update, or Delete)
	if err != nil {
		errTime := time.Now()
		file, fileErr := getFile(filepath, errTime)
		if fileErr == nil {
			_, _ = file.Write([]byte("Method:\t" + errorType + "\n"))
			_, _ = file.Write([]byte("Time:\t" + errTime.Format(time.ANSIC) + "\n"))
			_, _ = file.Write([]byte("Error:\t" + err.Error() + "\n"))
			_, _ = file.Write([]byte("Value:\t" + fmt.Sprintf("%+v", value) + "\n"))
			b, err := json.Marshal(value)
			if err == nil {
				_, _ = file.Write([]byte("JSON:\t" + string(b) + "\n\n\n"))
			}
			_ = file.Close()
		}
	}
}

