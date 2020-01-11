package entity

// Event represents Food Menu Event
type Event struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Description string
	Location    string
	Time        string
	Image       string `gorm:"type:varchar(255)"`
}

// User represents application user
type User struct {
	ID       uint
	FullName string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null; unique"`
	Phone    string `gorm:"type:varchar(100);not null; unique"`
	Password string `gorm:"type:varchar(255)"`
	RoleID   uint
}

// Role repesents application user roles
type Role struct {
	ID    uint
	Name  string `gorm:"type:varchar(255)"`
	Users []User
}

//Session represents login user session
type Session struct {
	ID         uint
	UUID       string `gorm:"type:varchar(255);not null"`
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}
