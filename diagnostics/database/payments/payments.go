package payments

import (
	"backend/diagnostics"
	paymentsModel "backend/models/payments"
)

func WritePaymentErr(err error, errorType string, payment *paymentsModel.Payment) {
	filepath := "database/payments/payments.log"
	diagnostics.WriteToSpecificDiagnosticsFile(filepath, errorType, err, payment)
}

func WritePaymentsErr(err error, errorType string, payments []paymentsModel.Payment) {
	filepath := "database/payments/bulk-payments.log"
	for i := 0; i < len(payments); i++ {
		diagnostics.WriteToSpecificDiagnosticsFile(filepath, errorType, err, payments[i])
	}
}