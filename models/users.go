package userModel

type User struct{
	ID int64 `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"size:255;not null;index:idx_username"`
	Email string `gorm:"size:200;unique;index:idx_email;not null"`
	Password string `gorm:"size:255;not null"`
	LoginToken string `gorm:"size:255"`
	TempToken string `gorm:"size:255"`
	EmailVerify bool `gorm:"column:e_verify;type:tinyint(1);not null;default:0"`
	UserDisable bool `gorm:"type:tinyint(1);column:isDisable;not null; default:0"`
	CreateAt int64 `gorm:"autoCreateTime"`
	UpdateAt int64 `gorm:"autoUpdateTime"`
}