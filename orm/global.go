package orm

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	Name     string `json:"name" gorm:"unique"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Role     int8   `json:"role"`
	Forbid   int8   `json:"-"`
	Avatar   string `json:"avatar"`
	gorm.Model
}

type Video struct {
	UserId    uint   `json:"user_id"`
	Type      string `json:"type"`
	VideoUrl  string `json:"video_url"`
	VideoName string `json:"video_name"`
	gorm.Model
}
type Good struct {
	VideoId uint `json:"video_id"`
	UserId  uint `json:"user_id"`
	gorm.Model
}

type Collect struct {
	VideoId uint `json:"video_id"`
	UserId  uint `json:"user_id"`
	gorm.Model
}
type Comment struct {
	VideoId uint   `json:"video_id"`
	UserId  uint   `json:"user_id"`
	ReplyId uint   `json:"reply_id"`
	Text    string `json:"text"`
	gorm.Model
}
type BulletChat struct {
	VideoId   uint   `json:"video_id"`
	VideoTime string `json:"video_time"`
	UserId    uint   `json:"user_id"`
	Text      string `json:"text"`
	gorm.Model
}
