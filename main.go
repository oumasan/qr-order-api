package main

import (
	"log"
	"fmt"
	"strconv"
	"time"
	"net/http"
	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/labstack/echo/middleware"

	"qr-order-system/models"
	"qr-order-system/structs"
)

type Shop struct {
	ID        uint      `json:"id  param:"id""`
	Name      string    `json:"name"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

var (
	DB  *gorm.DB
	err error
)

func main() {
	// インスタンス生成
	e := echo.New()
	// CORSの設定追加
    e.Use(middleware.CORS())

	// DBアクセス
	dsn := "user:password@tcp(mysql:3306)/qr_order_system?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(dsn + "database can't connect")
	}
	
	// ショップ情報取得
	e.GET("/shops/:id", getShop)

	// ショップごとのメニュー取得
	e.GET("/shops/:id/category", func(c echo.Context) error {
		response := []structs.Category{}
		// 大分類取得
		broadCategoryList := models.GetBroadCategory(c)
		
		// 小分類取得
		for index, broadCategory := range broadCategoryList {
			subCategory := models.GetSubCategory(broadCategory)
			category := structs.Category{broadCategory, subCategory}
			response = append(response, category)
			
			fmt.Print(index)
		}

		return c.JSON(http.StatusOK, response)
	})

	// 大分類登録
	e.POST("/shops/:id/broad-category", func(c echo.Context) error {
		// ショップID
		//broadCategoryName := c.Param("broadCategoryName")
		var requestBody structs.BroadCategoryBody
		if err := c.Bind(&requestBody); err != nil {
			return err
		}
		models.CreateBroadCategory(c, requestBody)
		return c.JSON(http.StatusOK, true)
	})

	// 大分類削除
	e.DELETE("/broad-category/:id", func(c echo.Context) error {
		models.DeleteBroadCategory(c)
		return c.JSON(http.StatusOK, true)
	})

	// 小分類登録
	e.POST("/shops/:id/sub-category", func(c echo.Context) error {
		var requestBody structs.SubCategoryBody
		if err := c.Bind(&requestBody); err != nil {
			return err
		}
		models.CreateSubCategory(requestBody)
		return c.JSON(http.StatusOK, true)
	})

	// 小分類削除
	e.DELETE("/sub-category/:id", func(c echo.Context) error {
		models.DeleteSubCategory(c)
		return c.JSON(http.StatusOK, true)
	})

	e.Logger.Fatal(e.Start(":9090"))
}

// ショップ取得
func getShop(c echo.Context) error {
	shop := Shop{}
	DB.First(&shop, c.Param("id"))
	return c.JSON(http.StatusOK, shop)
}

func stringToUint(s string) uint {
	uintVal, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Fatalln("パラメータが正しくありません")
	}
	return uint(uintVal)
}