package util

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Pagination struct {
	Page   int64
	Size   int64
	Offset int64
	Total  int64
}

func NewPagination(c *gin.Context) *Pagination {
	p := &Pagination{
		Page: cast.ToInt64(c.Query("page")),
		Size: cast.ToInt64(c.Query("size")),
	}
	if p.Page == 0 || p.Size == 0 {
		p = &Pagination{
			Page: 1,
			Size: 10,
		}
	}
	p.Offset = (p.Page - 1) * p.Size
	return p
}

func (p *Pagination) GormPaginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(cast.ToInt(p.Offset)).Limit(cast.ToInt(p.Size))
	}
}
