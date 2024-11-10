package moduls


// Struktur data untuk tabel produk
type Product struct {
    ID          int64   `gorm:"primaryKey" json:"id"`
    NamaProduct string  `gorm:"type:varchar(300)" json:"nama_product"`
    Size        string  `gorm:"type:varchar(100)" json:"size"`
    Stock       int     `gorm:"type:int" json:"stock"`
    Price       float64 `gorm:"type:decimal(10,2)" json:"price"`
}