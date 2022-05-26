package database

type Notification struct {
	ID		uint		`gorm:"primaryKey;autoIncrement"`
	UserId	uint		`gorm:"column:user_id"`
	Header	string		`gorm:"column:header"`
	Message	string		`gorm:"column:message"`	
}