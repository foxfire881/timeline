package controller

import (
	. "aio/internal/apps/timeline/models"
	. "aio/internal/apps/timeline/util"
	. "aio/internal/common"

	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterItem(api *gin.RouterGroup, web *gin.RouterGroup) {
	r := api.Group("/item")
	r.GET("", itemList)
	r.POST("", itemCreate)
	r.PUT("", itemUpdate)
	r.DELETE("", itemDelete)

	r = web.Group("")
	r.GET("/admin", itemAdmin)
}

func itemAdmin(c *gin.Context) {
	data := gin.H{"appname": "Timeline", "page": "timeline/admin"}
	c.HTML(200, "public/page.gtpl", data)
}

func itemList(c *gin.Context) {
	var ipage Page[Item]
	var query QueryItem
	c.ShouldBindQuery(&query)

	q := Conn.Timeline.Model(&Item{})
	if query.GroupID != 0 {
		q = q.Where("`group_id`=?", query.GroupID)
	}
	if query.SubgroupID != "" {
		q = q.Where("`subgroup_id`=?", query.SubgroupID)
	}
	if query.StartStr != "" {
		start := FormatTimeS2I(query.StartStr)
		q = q.Where("`start`>=?", *start)
	}
	if query.EndStr != "" {
		end := FormatTimeS2I(query.EndStr)
		q = q.Where("`end`<=?", *end)
	}
	if query.Key != "" {
		q = q.Where("`content` like ?", "%"+query.Key+"%").Or("`title` like ?", "%"+query.Key+"%")
	}

	q = q.Preload("Group")
	q = q.Preload("Subgroup")
	q = q.Order(query.Order + " " + query.Asc)
	if query.All == false {
		ipage = Paginate[Item](c, q, query.PerPage, query.Page)
		Success(c, ipage)
	} else {
		var items []Item
		q.Find(&items)
		Success(c, items)
	}
}

func itemCreate(c *gin.Context) {
	var item Item
	c.ShouldBindJSON(&item)
	item.Fixed()
	Conn.Timeline.Model(&Item{}).Create(&item)
	Success(c, item.ID)
}

func itemUpdate(c *gin.Context) {
	var item Item
	c.ShouldBindJSON(&item)
	item.Fixed()
	Success(c, Conn.Timeline.Model(&Item{}).Where("id =?", item.ID).Save(&item).RowsAffected) //这里要用Save，不能用Updates。Updates会忽略值为nil的字段不更新，Save则是强制所有字段不管是不是nil都更新
}

func itemDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	Success(c, Conn.Timeline.Model(&Item{}).Delete(&Item{}, "id", id).RowsAffected)
}
