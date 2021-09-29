package orders

type course struct {
	ID        int    `json:"id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Thumbnail string `json:"thumbnail" binding:"required"`
	Price     int    `json:"price" binding:"required"`
	Level     string `json:"level" binding:"required"`
}

type user struct {
	ID    int    `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type CreateOrderInput struct {
	User   user   `json:"user" binding:"required"`
	Course course `json:"course" binding:"required"`
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
