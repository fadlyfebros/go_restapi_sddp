package moduls

type Parfum struct {
	ID       int64   `gorm:"primaryKey" json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	Category string  `json:"category"`
}

type Cart struct {
	ID     int64      `gorm:"primaryKey" json:"id"`
	UserID int64      `json:"user_id"`
	Items  []CartItem `gorm:"foreignKey:CartID" json:"items"`
}

type CartItem struct {
	ID       int64   `gorm:"primaryKey" json:"id"`
	CartID   int64   `json:"cart_id"`
	ParfumID int64   `json:"parfum_id"`
	Quantity int     `json:"quantity"`
	UserID   int64   `json:"user_id"`
	Parfum   Parfum  `gorm:"foreignKey:ParfumID;references:ID"`
}

type Order struct {
	ID          int64   `gorm:"primaryKey" json:"id"`
	UserID      int64   `json:"user_id"`
	CartID      int64   `json:"cart_id"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`
}