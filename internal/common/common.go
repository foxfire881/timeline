package common

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "成功", "data": data})
}

func Fail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"status": 1, "msg": msg, "data": nil})
	c.Abort()
}

func GlobalRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 异常时返回统一格式
				Fail(c, fmt.Sprintf("全局异常: %v", err))
			}
		}()
		c.Next()
	}
}

type PageQuery struct {
	Page    int `form:"page,default=1" json:"page"`
	Perpage int `form:"perPage,default=10" json:"perPage"`
}

type Page[T any] struct {
	Size    int   `json:"size"`
	Current int   `json:"current"`
	Total   int64 `json:"total"`
	Pages   int   `json:"pages"`
	Items   []T   `json:"items"`
}

func Paginate0[T any](c *gin.Context, baseQuery *gorm.DB, size int, page int) Page[T] {
	var total int64
	var err error
	err = baseQuery.Session(&gorm.Session{}).Count(&total).Error
	if err != nil {
		Fail(c, err.Error())
	}

	list := make([]T, 0) //不能用var list []T,否则前端序列化为null报错
	if page < 1 {
		page = 1
	}
	err = baseQuery.Limit(size).Offset((page - 1) * size).Find(&list).Error
	if err != nil {
		Fail(c, err.Error())
	}

	return Page[T]{
		Size:    size,
		Current: page,
		Total:   total,
		Pages:   (int(total) + size - 1) / size,
		Items:   list,
	}
}

func Paginate[T any](c *gin.Context, baseQuery *gorm.DB, size int, page int) Page[T] {
	var total int64
	var err error
	err = baseQuery.Session(&gorm.Session{}).Count(&total).Error
	if err != nil {
		Fail(c, err.Error())
	}

	list := make([]T, 0) //不能用var list []T,否则前端序列化为null报错
	if page < 1 {
		page = 1
	}

	if reflect.TypeOf(baseQuery.Statement.Model).Elem().Name() == reflect.TypeFor[T]().Name() {
		err = baseQuery.Limit(size).Offset((page - 1) * size).Find(&list).Error //Find可以保留之前Preload
	} else {
		err = baseQuery.Limit(size).Offset((page - 1) * size).Scan(&list).Error //Scan导致之前Preload失效
	}

	if err != nil {
		Fail(c, err.Error())
	}

	return Page[T]{
		Size:    size,
		Current: page,
		Total:   total,
		Pages:   (int(total) + size - 1) / size,
		Items:   list,
	}
}
