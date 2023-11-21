package models

import (
	"github.com/labstack/echo"
	"qr-order-system/structs"
)

// 大分類取得
func GetBroadCategory(c echo.Context) []structs.BroadCategory {
	broadCategoryList := []structs.BroadCategory{}
	getDB().Where("shop_id = ?", c.Param("id")).Find(&broadCategoryList)
	return broadCategoryList
} 

// 小分類取得
func GetSubCategory(broadCategory structs.BroadCategory) []structs.SubCategory {
	subCategoryList := []structs.SubCategory{}
	getDB().Where("broad_category_id = ?", broadCategory.ID).Find(&subCategoryList)
	return subCategoryList
}