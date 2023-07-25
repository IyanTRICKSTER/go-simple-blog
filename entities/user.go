package entities

type User struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Username   string `gorm:"not null" json:"username"`
	Email      string `gorm:"not null" json:"email"`
	Password   string `gorm:"not null" json:"password"`
	ProfilePic string `gorm:"not null;type:text;longtext" json:"profilePic"`
	CreatedAt  int64  `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt  int64  `gorm:"autoUpdateTime:milli" json:"updatedAt"`
	DeletedAt  *int64 `json:"deletedAt"`
	Posts      []Post `gorm:"foreignKey:UserID;references:ID"`
}
