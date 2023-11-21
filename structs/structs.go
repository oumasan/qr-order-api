package structs

import (
	"time"
)

// type BroadCategory struct {
// 	ID uint
// 	Name string
// }

// type SubCategory struct {
// 	ID uint
// 	Name string
// 	Price uint
// 	ImageUrl string
// }

type BroadCategory struct {
	ID        uint      `json:"id"`
	shopId    uint      `json:"shopId"`
	Name      string    `json:"name"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

type SubCategory struct {
	ID        uint      `json:"id  param:"id""`
	BroadCategoryId uint `json:"broadCategoryId"`
	Name      string    `json:"name"`
	Price     uint      `json:"price"`
	ImageUrl  string    `json:"imageUrl"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}

type Category struct {
	BroadCategory BroadCategory `json:"broadCategory"`
	SubCategoryList []SubCategory `json:"subCategoryList"`
}