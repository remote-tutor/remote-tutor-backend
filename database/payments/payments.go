package payments

import (
	dbInstance "backend/database"
	"backend/diagnostics"
	paymentsModel "backend/models/payments"
	"time"
)

// CreatePayment inserts a new payment to the database
func CreatePayment(payment *paymentsModel.Payment) error {
	err := dbInstance.GetDBConnection().Create(payment).Error
	diagnostics.WriteError(err, "database.log", "CreatePayment")
	return err
}

// UpdatePayment updates the payment data in the database
func UpdatePayment(payment *paymentsModel.Payment) error {
	err := dbInstance.GetDBConnection().Save(payment).Error
	return err
}

// DeletePayment deletes the payment
func DeletePayment(payment *paymentsModel.Payment) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(payment).Error
	diagnostics.WriteError(err, "database.log", "DeletePayment")
	return err
}

// GetPaymentsByUserAndMonth gets the payment of specific user in a specific month
func GetPaymentsByUserAndMonth(userID uint, startDate, endDate time.Time) []paymentsModel.Payment {
	payments := make([]paymentsModel.Payment, 0)
	dbInstance.GetDBConnection().Where("user_id = ? AND start_date >= ? AND end_date < ?",
		userID, startDate, endDate).Find(&payments)
	return payments
}

// GetPaymentByUserAndWeek returns the payment of the user in a specific week (if found)
func GetPaymentByUserAndWeek(userID uint, eventDate time.Time) paymentsModel.Payment {
	var payment paymentsModel.Payment
	dbInstance.GetDBConnection().Where("user_id = ? AND start_date <= ? AND end_date >= ?",
		userID, eventDate, eventDate).First(&payment)
	return payment
}
