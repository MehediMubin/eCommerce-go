package database

type Product struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImgUrl string `json:"imageUrl"`
}

var ProductList []Product

func NewProduct(id int, title, description string, price float64, imgUrl string) *Product {
	return &Product{
		ID: id,
		Title: title,
		Description: description,
		Price: price,
		ImgUrl: imgUrl,
	}
}

func init() {
	product1 := NewProduct(1, "Orange", "Orange is orange", 3.23, "https://www.dole.com/sites/default/files/media/2025-01/oranges.png")

	product2 := NewProduct(2, "Apple", "I love Apple", 2.23, "https://5.imimg.com/data5/AK/RA/MY-68428614/apple.jpg")

	product3 := NewProduct(3, "Banana", "Banana is interesting", 4.23, "https://png.pngtree.com/png-clipart/20220716/ourmid/pngtree-banana-yellow-fruit-banana-skewers-png-image_5944324.png")

	product4 := NewProduct(4, "Grape", "Grape is sour", 10.23, "https://hips.hearstapps.com/hmg-prod/images/766/grapes-main-1506688521.jpg?resize=640:*")
	
	product5 := NewProduct(5, "Mango", "Mango is fruit ka raja", 6.23, "https://5.imimg.com/data5/SELLER/Default/2023/9/344928632/OW/RQ/XC/25352890/yellow-mango-500x500.jpeg")

	ProductList = append(ProductList, *product1, *product2, *product3, *product4, *product5)
}