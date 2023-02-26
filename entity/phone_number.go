package entity

type PhoneNumber struct {
	ID          uint64 `gorm:"primaryKey" json:"id"`
	CountryCode string `json:"country_code" binding:"required"`
	Number      string `json:"number" binding:"required"`
	UserID      uint64 `gorm:"foreignKey" json:"user_id" binding:"required"`
	User        *User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"transaksi,omitempty"`
}
