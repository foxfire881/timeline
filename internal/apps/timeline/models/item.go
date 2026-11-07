package models

import (
	. "aio/internal/apps/timeline/util"

	"gorm.io/gorm"
)

type Item struct {
	ID             int64  `gorm:"primaryKey" json:"id"`
	Start          *int64 `json:"-"`
	End            *int64 `json:"-"`
	Content        string `json:"content"`
	Title          string `json:"title"`
	GroupID        int    `json:"group_id"`
	SubgroupID     int    `json:"subgroup_id"`
	Type           string `json:"type"`
	Color1         string `json:"color1"`
	Color2         string `json:"color2"`
	PrecisionStart string `json:"precision_start"`
	PrecisionEnd   string `json:"precision_end"`

	StartStr string   `gorm:"-" json:"start_str"`
	EndStr   string   `gorm:"-" json:"end_str"`
	Group    Group    `gorm:"foreignKey:GroupID;ignore" json:"group"`
	Subgroup Subgroup `gorm:"foreignKey:SubgroupID;ignore" json:"subgroup"`
}

type QueryItem struct {
	GroupID    int    `form:"group_id"`
	SubgroupID string `form:"subgroup_id"`
	StartStr   string `form:"start_str"`
	EndStr     string `form:"end_str"`
	Key        string `form:"key"`
	Order      string `form:"order,default=id"`
	Asc        string `form:"asc,default=desc"`
	All        bool   `form:"all,default=false"`
	Page       int    `form:"page,default=1"`
	PerPage    int    `form:"perPage,default=5"`
}

func (Item) TableName() string {
	return "item"
}

// 数据库取出后有HOOK，自动触发
func (item *Item) AfterFind(tx *gorm.DB) error {
	item.StartStr = FormatTimeI2S(item.Start, item.PrecisionStart)
	item.EndStr = FormatTimeI2S(item.End, item.PrecisionEnd)
	return nil // 必须返回 nil
}

// 入库前绑定，需要手动触发
func (item *Item) Fixed() error {
	item.Start = FormatTimeS2I(item.StartStr)
	item.End = FormatTimeS2I(item.EndStr)
	return nil // 必须返回 nil
}
