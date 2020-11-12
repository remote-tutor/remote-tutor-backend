package aws

import "backend/diagnostics"

func WriteAWSUploadError(err error, errorType string) {
	filepath := "aws/upload.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filepath, errorType, err, "Unable to upload file to AWS")
}
