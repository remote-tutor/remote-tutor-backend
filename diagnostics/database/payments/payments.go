package payments

import (
	"backend/diagnostics"
	paymentsModel "backend/models/payments"
)

func WritePaymentErr(err error, errorType string, payment *paymentsModel.Payment) {
	filepath := "database/payments/payments.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filepath, errorType, err, payment)
}
