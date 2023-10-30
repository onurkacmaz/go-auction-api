package dto

type Artist struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Bio     string `json:"bio" gorm:"type:text"`
	Email   string `json:"email"`
	Phone   string `json:"phone" gorm:"type:varchar(12)"`
	Address string `json:"address" gorm:"type:text"`
}
