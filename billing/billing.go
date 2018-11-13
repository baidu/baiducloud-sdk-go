package billing

// Billing defines the struct of payment
type Billing struct {
	PaymentTiming string      `json:"paymentTiming"`
	Reservation   Reservation `json:"reservation,omitempty"`
}

// Reservation is the reservation info of billing
type Reservation struct {
	ReservationLength   int    `json:"reservationLength"`
	ReservationTimeUnit string `json:"reservationTimeUnit"`
}
