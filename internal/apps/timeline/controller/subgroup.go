package controller

import (
	. "aio/internal/apps/timeline/models"
	. "aio/internal/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterSubgroup(api *gin.RouterGroup) {
	r := api.Group("/subgroup")
	r.GET("", subgroupList)
	r.POST("", subgroupCreate)
	r.PUT("", subgroupUpdate)
	r.DELETE("", subgroupDelete)
}

func subgroupList(c *gin.Context) {
	group_id := c.Query("group_id")
	var list []Subgroup
	q := Conn.Timeline.Model(&Subgroup{})
	if group_id != "" {
		q = q.Where("group_id = ?", group_id)
	}
	q.Preload("Group").Find(&list).Order("group_id ASC").Order("id ASC")
	Success(c, list)
}

func subgroupCreate(c *gin.Context) {
	var subgroup Subgroup
	c.ShouldBindJSON(&subgroup)
	Conn.Timeline.Model(&Subgroup{}).Create(&subgroup)
	Success(c, subgroup.ID)
}

func subgroupUpdate(c *gin.Context) {
	var subgroup Subgroup
	c.ShouldBindJSON(&subgroup)
	Success(c, Conn.Timeline.Model(&Subgroup{}).Where("id =?", subgroup.ID).Save(&subgroup).RowsAffected)
}

func subgroupDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var list []Subgroup
	if Conn.Timeline.Model(&Item{}).Where("subgroup_id", id).Find(&list).RowsAffected > 0 {
		Fail(c, "存在下级级联，无法删除")
	} else {
		Success(c, Conn.Timeline.Model(&Subgroup{}).Delete(&Subgroup{}, "id", id).RowsAffected)
	}
}
