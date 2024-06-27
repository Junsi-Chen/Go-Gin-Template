package util

import "gorm.io/gorm"

// 分页查询的封装

// PageInfo 存放page对象的结构体
type PageInfo struct {
	Page     int `json:"page" binding:"required"`     // 当前页
	PageSize int `json:"pageSize" binding:"required"` // 数据条数
}

// PageRecords 响应体
type PageRecords struct {
	Records  interface{} `json:"records"`  // 数据集合
	Total    int64       `json:"total"`    // 总条数
	Page     int64       `json:"page"`     // 当前页
	PageSize int64       `json:"pageSize"` // 数据条数
}

// LimitPage 分页查询
func LimitPage(pageInfo *PageInfo) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// 处理page
		page := pageInfo.Page
		if page <= 0 {
			page = 1
		}
		// 处理pageSize
		pageSize := pageInfo.PageSize
		switch {
		case pageSize > 100:
			pageSize = 100
			break
		case pageSize <= 0:
			pageSize = 10
			break
		}
		// 计算分页偏移
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
