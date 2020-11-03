package utils

import "backend/diagnostics"

var filepath = "utils/convert-types.log"

func ConvertToBoolErr(err error, value string) {
	diagnostics.WriteToSpecificDiagnosticsFile(filepath, "ConvertToBoolErr", err, value)
}

func ConvertToIntErr(err error, value string) {
	diagnostics.WriteToSpecificDiagnosticsFile(filepath, "ConvertToIntErr", err, value)
}

func ConvertToUIntErr(err error, value string) {
	diagnostics.WriteToSpecificDiagnosticsFile(filepath, "ConvertToUIntErr", err, value)
}
