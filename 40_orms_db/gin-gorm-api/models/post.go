package models

type Post struct {
    ID     uint   `gorm:"primaryKey" json:"id"`
    Title  string `json:"title"`
    Body   string `json:"body"`
    UserID uint   `json:"user_id"`
}
