package aws

import (
	"backend/diagnostics"
)

func WriteAWSPartErr(err error, errorType string) {
	filepath := "aws/parts.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filepath, errorType, err, "Unable to upload video part")
}
