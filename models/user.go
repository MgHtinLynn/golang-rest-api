package models

type User struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Password   string `json:"password"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
type CreateUserInput struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Address string `json:"address" binding:"required"`
}
type UpdateUserInput struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
