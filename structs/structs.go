package structs

type BroadCategory struct {
	ID        uint      `json:"id"`
	ShopId    uint      `json:"shopId"`
	Name      string    `json:"name"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}

type SubCategory struct {
	ID        uint      `json:"id"`
	BroadCategoryId string `json:"broadCategoryId"`
	Name      string    `json:"name"`
	Price     string      `json:"price"`
	ImageUrl  string    `json:"imageUrl"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}

type Category struct {
	BroadCategory BroadCategory `json:"broadCategory"`
	SubCategoryList []SubCategory `json:"subCategoryList"`
}

// RequestBody はJSONリクエストボディの構造体です
type BroadCategoryBody struct {
	BroadCategoryName   string `json:"broadCategoryName"`
}

// RequestBody はJSONリクエストボディの構造体です
type SubCategoryBody struct {
	BroadCategoryId string `json:"broadCategoryId"`
	Name      string    `json:"name"`
	Price     string      `json:"price"`
	ImageUrl  string    `json:"imageUrl"`
}

type OrderBody struct {
	ID uint `json:"id"`
	Name      string    `json:"name"`
	Price     uint    `json:"price"`
	Count     uint    `json:"count"`
}

type CartBody struct {
	ShopId string `json:"shopId"`
	AccountantId string `json:"accountantId"`
	Cart []OrderBody `json:"cart"`
}

type Accountant struct {
	ID uint `json:"id"`
	SubCategoryId uint `json:"subCategoryId"`
	ShopId string `json:"shopId"`
	AccountantId string `json:"accoutantId"` 
	Name      string    `json:"name"`
	Price     uint    `json:"price"`
	Count     uint    `json:"count"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}