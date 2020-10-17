package diagnostics

import (
	"fmt"
	"os"
	"time"
)

func WriteError(dbError error, fileName, methodName string) {
	if dbError != nil {
		filePath := fmt.Sprintf("diagnostics/%s", fileName)
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
		_, _ = file.Write([]byte("Time:\t" + time.Now().Format(time.ANSIC) + "\n"))
		_, _ = file.Write([]byte("Error:\t" + dbError.Error() + "\n"))
		_, _ = file.Write([]byte("From:\t" + methodName + "\n\n\n"))
		_ = file.Close()
	}
}
