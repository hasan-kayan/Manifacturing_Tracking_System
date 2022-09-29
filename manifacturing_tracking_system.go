package main

import (
	"fmt"
	"strings"
	"time"
)

const sharp = "#"

type material struct {
	name    string
	id      int
	stock   int
	quality string
}
type furniture struct {
	name                           string
	id                             int
	color                          string
	price                          float32
	sizeWidth                      float32
	sizeHeight                     float32
	sizeLength                     float32
	foldability                    bool
	quantityOfRawMaterialsRequired float32 //KG
}

func (property furniture) furnitureProperty() {
	if property.foldability == true {
		fmt.Println("Kolaylıkla taşınabilir")
	} else {
		fmt.Println("Taşınması zor")
	}
}

func main() {
	factorySize()
	fmt.Println("Fabrikanız Hazırlanıyor...")
	progressBar()
	structOfContents()

}
func factorySize() {
	var factorySizeScan string
	var factorySize int
	var numberOfWorkers int
	fmt.Println("Fabrikanızın boyutununu giriniz.[küçük , orta , büyük ]")
	fmt.Scanf("%s", &factorySizeScan)
	if factorySizeScan == "küçük" {
		numberOfWorkers = 5
		factorySize = 300
	} else if factorySizeScan == "orta" {
		numberOfWorkers = 10
		factorySize = 750
	} else if factorySizeScan == "büyük" {
		numberOfWorkers = 15
		factorySize = 1250
	}
	fmt.Println("Fabrikada Çalışan Sayısı ", numberOfWorkers, "\nFabrika Boyutunuz ", factorySize, "m2")

}
func progressBar() {
	for i := 0; i < 101; i++ {
		fmt.Printf("\r[%-100s] %d%%", strings.Repeat(sharp, i), i)
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("")

}

func structOfContents() {
	furniture1 := []furniture{
		{
			name:                           "Seat",
			id:                             0,
			color:                          "Red",
			price:                          500,
			sizeWidth:                      90,
			sizeHeight:                     50,
			sizeLength:                     110,
			quantityOfRawMaterialsRequired: 15.4,
			foldability:                    true,
		},
		{
			name:                           "Chair",
			id:                             1,
			color:                          "Blue",
			price:                          215,
			sizeWidth:                      35,
			sizeHeight:                     96,
			sizeLength:                     41,
			quantityOfRawMaterialsRequired: 27.8,
			foldability:                    true,
		},
		{
			name:                           "Table",
			id:                             2,
			color:                          "Green",
			price:                          315,
			sizeWidth:                      80,
			sizeHeight:                     75,
			sizeLength:                     140,
			quantityOfRawMaterialsRequired: 54.7,
		},
		{
			name:                           "Wardrobe",
			id:                             3,
			color:                          "Yellow",
			price:                          613,
			sizeWidth:                      50,
			sizeHeight:                     190,
			sizeLength:                     210,
			quantityOfRawMaterialsRequired: 46,
		},
		{
			name:                           "Coffee Table",
			id:                             4,
			color:                          "Black",
			price:                          305,
			sizeWidth:                      50,
			sizeHeight:                     42,
			sizeLength:                     50,
			quantityOfRawMaterialsRequired: 12.4,
			foldability:                    true,
		},
		{
			name:                           "TV Unit",
			id:                             5,
			color:                          "White",
			price:                          425,
			sizeWidth:                      35,
			sizeHeight:                     50,
			sizeLength:                     180,
			quantityOfRawMaterialsRequired: 21,
		},
	}

	material := []material{
		{
			name:    "Wood",
			id:      0,
			stock:   500,
			quality: "Very Good",
		},
		{
			name:    "MDF",
			id:      1,
			stock:   500,
			quality: "Good",
		},
		{
			name:    "Plywood",
			id:      2,
			stock:   500,
			quality: "Bad",
		},
		{
			name:    "Hardboard",
			id:      3,
			stock:   500,
			quality: "Too Bad",
		},
	}
	var orderFurniture int
	var orderMaterial int
	fmt.Println("Sipariş Vermek İstediğiniz Ürünü ve Yapılmasını İstediğiniz Hammaddesini Giriniz.[Seat : 0 , Chair : 1 , Table : 2 , Wardobe : 3 , Coffee Table : 4, TV Unit : 5],[Wood : 0 , MDF : 1 , Plywood : 2 , Hardboard : 3 ")
	fmt.Scan(&orderFurniture, &orderMaterial)
	for i := 0; i < len(furniture1); i++ {
		if orderFurniture == furniture1[i].id {
			if orderMaterial == material[i].id {
				fmt.Println("Siparişiniz Hazırlanıyor...")
				progressBar()
				fmt.Println("Siparişiniz Hazırlandı.")
				fmt.Println("Siparişinizin Özellikleri :")
				fmt.Println("Ürün Adı :", furniture1[orderFurniture].name)
				fmt.Println("Ürün ID :", furniture1[orderFurniture].id)
				fmt.Println("Ürün Fiyatı :", furniture1[orderFurniture].price)
				furniture.furnitureProperty(furniture1[orderFurniture])
				fmt.Println("Ürün İçin Gerekli Ham Madde Miktarı :", furniture1[orderFurniture].quantityOfRawMaterialsRequired)
				fmt.Println("Ham Madde Adı :", material[orderMaterial].name)
				fmt.Println("Ham Madde ID :", material[orderMaterial].id)
				fmt.Println("Ham Madde Stok Miktarı :", material[orderMaterial].stock)
				fmt.Println("Ham Madde Kalitesi :", material[orderMaterial].quality)
			}
		}
	}
}
