package payments

import (
	dbInstance "backend/database"
	paymentsModel "backend/models/payments"
	"time"
)

// CreatePayment inserts a new payment to the database
func CreatePayment(payment *paymentsModel.Payment) {
	dbInstance.GetDBConnection().Create(payment)
}

// UpdatePayment updates the payment data in the database
func UpdatePayment(payment *paymentsModel.Payment) {
	dbInstance.GetDBConnection().Save(payment)
}

// GetPaymentsByUserAndMonth gets the payment of specific user in a specific month
func GetPaymentsByUserAndMonth(userID uint, startDate, endDate time.Time) []paymentsModel.Payment {
	payments := make([]paymentsModel.Payment, 0)
	dbInstance.GetDBConnection().Where("user_id = ? AND start_date >= ? AND end_date < ?",
		userID, startDate, endDate).Find(&payments)
	return payments
}
