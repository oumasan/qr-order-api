package models

import (
	"time"
	"qr-order-system/structs"
	"github.com/labstack/echo"
	"fmt"
)

// 注文確定
func CreateAccountant(body structs.CartBody) {
	t := time.Now()
	for _, v := range body.Cart {
		accountant := structs.Accountant {
			SubCategoryId: v.ID,
			ShopId: body.ShopId,
			AccountantId: body.AccountantId,
			Name: v.Name, 
			Price: v.Price,
			Count: v.Count,
			CreateAt: t.Format("2006/1/2 15:04:05"), 
			UpdateAt: t.Format("2006/1/2 15:04:05"),
		}
		fmt.Printf("(%%+v) %+v\n", accountant)
		getDB().Create(&accountant)
	} 
}

// 会計取得
func GetAccountant(c echo.Context) []structs.Accountant {
	accountantList := []structs.Accountant{}
	fmt.Printf("(%%+v) %+v\n", c.Param("id"))
	getDB().Where("accountant_id = ?", c.Param("id")).Find(&accountantList)
	fmt.Printf("(%%+v) %+v\n", accountantList)
	return accountantList
}