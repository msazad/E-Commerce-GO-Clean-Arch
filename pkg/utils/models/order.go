package models

type OrderPaymentDetails struct {
	UserID     int     `json:"user_id"`
	Username   string  `json:"username"`
	Razor_id   string  `json:"razor_id"`
	OrderID    int     `json:"order_id"`
	FinalPrice float64 `json:"final_price"`
}

type InvoiceData struct {
	Title       string
	Quantity    int
	Price       int
	TotalAmount int
}

type Invoice struct {
	Name         string
	Address      string
	InvoiceItems []*InvoiceData
}
type Order struct {
	UserID          int `json:"user_id"`
	AddressID       int `json:"address_id"`
	PaymentMethodID int `json:"payment_id"`
	CouponID        int `json:"coupon_id"`
}
