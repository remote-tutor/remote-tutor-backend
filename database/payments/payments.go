package payments

import (
	dbInstance "backend/database"
	paymentsDiagnostics "backend/diagnostics/database/payments"
	paymentsModel "backend/models/payments"
	"time"
)

// CreatePayment inserts a new payment to the database
func CreatePayment(payment *paymentsModel.Payment) error {
	err := dbInstance.GetDBConnection().Create(payment).Error
	paymentsDiagnostics.WritePaymentErr(err, "Create", payment)
	return err
}

// UpdatePayments inserts/removes payments for multiple users
func UpdatePayments(paymentsToAdd, paymentsToRemove []paymentsModel.Payment) error {
	var err error
	if len(paymentsToAdd) > 0 && len(paymentsToRemove) > 0 {
		err = dbInstance.GetDBConnection().Create(&paymentsToAdd).Delete(paymentsToRemove).Error
	} else if len(paymentsToAdd) > 0 {
		err = dbInstance.GetDBConnection().Create(&paymentsToAdd).Error
	} else {
		err = dbInstance.GetDBConnection().Delete(&paymentsToRemove).Error
	}

	paymentsDiagnostics.WritePaymentsErr(err, "CreateDelete", append(paymentsToAdd, paymentsToRemove...))
	return err
}

// UpdatePayment updates the payment data in the database
func UpdatePayment(payment *paymentsModel.Payment) error {
	err := dbInstance.GetDBConnection().Save(payment).Error
	paymentsDiagnostics.WritePaymentErr(err, "Update", payment)
	return err
}

// DeletePayment deletes the payment
func DeletePayment(payment *paymentsModel.Payment) error {
	err := dbInstance.GetDBConnection().Unscoped().Delete(payment).Error
	paymentsDiagnostics.WritePaymentErr(err, "Delete", payment)
	return err
}

// GetPaymentsByUserAndMonthAndClass gets the payment of specific user in a specific month
func GetPaymentsByUserAndMonthAndClass(userID uint, startDate, endDate time.Time, class string) []paymentsModel.Payment {
	payments := make([]paymentsModel.Payment, 0)
	dbInstance.GetDBConnection().Where("user_id = ? AND start_date >= ? AND end_date < ? AND class_hash = ?",
		userID, startDate, endDate, class).Find(&payments)
	return payments
}

// GetPaymentByUserAndWeekAndClass returns the payment of the user in a specific week (if found)
func GetPaymentByUserAndWeekAndClass(userID uint, eventDate time.Time, class string) paymentsModel.Payment {
	var payment paymentsModel.Payment
	dbInstance.GetDBConnection().Where("user_id = ? AND start_date <= ? AND end_date >= ? AND class_hash = ?",
		userID, eventDate, eventDate, class).First(&payment)
	return payment
}

// GetPaymentsByWeekAndClass returns the payments of all selected users in a specific week and class
func GetPaymentsByWeekAndClass(usersIDs []uint, date, endDate time.Time, class string) []paymentsModel.Payment {
	payments := make([]paymentsModel.Payment, 0)
	dbInstance.GetDBConnection().Where("user_id IN (?) AND start_date >= ? AND end_date <= ? AND class_hash = ?",
		usersIDs, date, endDate, class).Find(&payments)
	return payments
}
