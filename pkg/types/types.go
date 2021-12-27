package types

// Money  представляет собой денежную сумму в минимальных единицах (центы , копейки , дирамы и т.д.)
type Money int64

type PaymentCategory string //категории в которых был совершен платеж(авто , рестораны,аптеки итд.)

//Status представляет собой статус платежа
type PaymentStatus string

//Предопределенные статусы платежей
const (
	PaymentStatusOk         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

type Payment struct { // информация о платeже
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

type Phone string

//Account представляет информацию о счете пользователя
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}
