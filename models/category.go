package models

import (
	"log"
	"strconv"
	"time"
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

// 大分類登録
func CreateBroadCategory(c echo.Context, body structs.BroadCategoryBody) {
	t := time.Now()
	broadCategory := structs.BroadCategory{ShopId: stringToUint(c.Param("id")), Name: body.BroadCategoryName, CreateAt: t.Format("2006/1/2 15:04:05"), UpdateAt: t.Format("2006/1/2 15:04:05")}
	getDB().Create(&broadCategory)
}

// 大分類削除
func DeleteBroadCategory(c echo.Context) {
	getDB().Where("broad_category_id = ?", c.Param("id")).Delete([]structs.SubCategory{})
	getDB().Delete(&structs.BroadCategory{}, stringToUint(c.Param("id")))
}

// 小分類登録
func CreateSubCategory(body structs.SubCategoryBody) {
	t := time.Now()
	subCategory := structs.SubCategory{
		BroadCategoryId: body.BroadCategoryId, 
		Name: body.Name, 
		Price: body.Price,
		ImageUrl: body.ImageUrl,
		CreateAt: t.Format("2006/1/2 15:04:05"), 
		UpdateAt: t.Format("2006/1/2 15:04:05"),
	}
	getDB().Create(&subCategory)
}

// 小分類削除
func DeleteSubCategory(c echo.Context) {
	getDB().Delete(structs.SubCategory{}, c.Param("id"))
}

func stringToUint(s string) uint {
	uintVal, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Fatalln("パラメータが正しくありません")
	}
	return uint(uintVal)
}