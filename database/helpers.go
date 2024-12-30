package database

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Ordering(c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		orderBy := c.Query("orderBy")
		if orderBy != "" {
			ordering := "DESC"
			if c.Query("ordering") == "ASC" {
				ordering = "ASC"
			}
			return db.Order(orderBy + " " + ordering)
		}
		return db
	}
}

func Paginate(c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := c.QueryInt("page", 1)
		pageSize := c.QueryInt("pageSize", 15)

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
