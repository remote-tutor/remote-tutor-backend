package aws

import "backend/diagnostics"

func WriteAWSSignedURLErr(err error, errorType string) {
	filepath := "aws/signed-url.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filepath, errorType, err, "Unable to generate signed URL")
}

