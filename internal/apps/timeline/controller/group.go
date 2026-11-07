package controller

import (
	. "aio/internal/apps/timeline/models"
	. "aio/internal/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterGroup(api *gin.RouterGroup) {
	r := api.Group("/group")
	r.GET("", groupList)
	r.POST("", groupCreate)
	r.PUT("", groupUpdate)
	r.DELETE("", groupDelete)
}

func groupList(c *gin.Context) {
	var list []Group
	Conn.Timeline.Model(&Group{}).Find(&list)
	Success(c, list)
}

func groupCreate(c *gin.Context) {
	var group Group
	c.ShouldBindJSON(&group)
	Conn.Timeline.Model(&Group{}).Create(&group)
	Success(c, group.ID)
}

func groupUpdate(c *gin.Context) {
	var group Group
	c.ShouldBindJSON(&group)
	Success(c, Conn.Timeline.Model(&Group{}).Where("id =?", group.ID).Save(&group).RowsAffected)
}

func groupDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var list []Group
	if Conn.Timeline.Model(&Item{}).Where("id", id).Find(&list).RowsAffected > 0 {
		Fail(c, "存在下级级联，无法删除")
	} else {
		Success(c, Conn.Timeline.Model(&Group{}).Delete(&Group{}, "id", id).RowsAffected)
	}
}
