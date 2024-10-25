package domain

type User struct {
	ID       uint   `gorm:"primaryKey"`      // Khóa chính
	Username string `gorm:"unique;not null"` // Tên người dùng (không trùng lặp)
	Email    string `gorm:"unique;not null"` // Email (không trùng lặp)
	Password string `gorm:"not null"`        // Mật khẩu

	// Quan hệ với Role (Many-to-One)
	RoleID uint // Khóa ngoại tham chiếu đến Role
	Role   Role `gorm:"foreignKey:RoleID"` // Mối quan hệ với Role

	// Quan hệ với Task (One-to-Many, người dùng có thể tạo nhiều công việc)
	Tasks []Task `gorm:"foreignKey:CreatedBy"`

	// Quan hệ với các công việc được assign (Many-to-Many)
	AssignedTasks []Task `gorm:"many2many:task_assignments;"`
}
