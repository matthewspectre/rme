package user

// GORM model for users table.

type UserModel struct {
	ID       int    `gorm:"primaryKey;autoIncrement;column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	FullName string `gorm:"column:full_name"`
	Role     int    `gorm:"column:role"`
}

func (UserModel) TableName() string {
	return "users"
}
