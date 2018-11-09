package bce

type Billing struct {
	PaymentTiming string `json:"paymentTiming"`
	BillingMethod string `json:"billingMethod"`
}
type Reservation struct {
	ReservationLength   int    `json:"reservationLength"`
	ReservationTimeUnit string `json:"reservationTimeUnit"`
}
