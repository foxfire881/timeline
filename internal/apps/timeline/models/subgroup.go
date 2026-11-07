package models

type Subgroup struct {
	ID      int    `gorm:"primaryKey" json:"id"`
	Content string `json:"content"`
	GroupID int    `json:"group_id"`
	Stack   bool   `json:"stack"`

	Group Group `gorm:"foreignKey:GroupID;ignore" json:"group"`
}

func (Subgroup) TableName() string {
	return "subgroup"
}
