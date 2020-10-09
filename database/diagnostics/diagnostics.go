package diagnostics

import (
	"os"
	"time"
)

func WriteError(dbError error, methodName string) {
	if dbError != nil {
		file, err := os.OpenFile("diagnostics.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
		_, _ = file.Write([]byte("Time:\t" + time.Now().Format(time.ANSIC) + "\n"))
		_, _ = file.Write([]byte("Error:\t" + dbError.Error() + "\n"))
		_, _ = file.Write([]byte("From:\t" + methodName + "\n\n\n"))
		_ = file.Close()
	}
}
