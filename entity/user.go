package entity

type User struct {
	ID           uint64        `gorm:"primaryKey" json:"id"`
	Name         string        `json:"name" binding:"required"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
}
