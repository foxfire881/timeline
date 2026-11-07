package models

type Group struct {
	ID      int    `gorm:"primaryKey" json:"id"`
	Content string `json:"content"`
}

func (Group) TableName() string {
	return "group"
}
