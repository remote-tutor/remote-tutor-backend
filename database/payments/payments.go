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

func GiveAccessToAllStudents(startDate, endDate time.Time, class string) {
	// get all students' IDs who have payment in the given week
	usersWithPayments := make([]uint, 0)
	dbInstance.GetDBConnection().Where("start_date <= ? AND end_date >= ? AND class_hash = ?",
		startDate, endDate, class).Table("payments").Pluck("user_id", &usersWithPayments)
	// get all students' IDs who DOESN'T have payment in the given week
	usersWithNoPayments := make([]uint, 0)
	query := dbInstance.GetDBConnection()
	if len(usersWithPayments) > 0 {
		query = query.Where("user_id NOT IN (?) AND admin = 0", usersWithPayments)
	}
	query.Table("class_users").Pluck("user_id", &usersWithNoPayments)
	payments := make([]paymentsModel.Payment, len(usersWithNoPayments))
	for i := 0; i < len(payments); i++ {
		payments[i].StartDate = startDate
		payments[i].EndDate = endDate
		payments[i].ClassHash = class
		payments[i].UserID = usersWithNoPayments[i]
	}
	if len(payments) > 0 {
		dbInstance.GetDBConnection().Create(&payments)
	}
}
