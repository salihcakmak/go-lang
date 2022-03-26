package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() {
	productResponse, getErr := http.Get("http://localhost:3000/products")

	if getErr != nil {
		fmt.Println(getErr)
	}

	productsByte, readErr := ioutil.ReadAll(productResponse.Body)

	if readErr != nil {
		fmt.Println(readErr)
	}

	var products []Product
	json.Unmarshal(productsByte, &products)
	fmt.Println(products)

	defer productResponse.Body.Close()
}

func AddProduct() {
	product := Product{ProductName: "Telefon", CategoryId: 1, UnitPrice: 7999.90}
	productByte, _ := json.Marshal(&product)

	response, postErr := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(productByte))
	if postErr != nil {
		fmt.Println(postErr)
	}

	bodyy, _ := ioutil.ReadAll(response.Body)

	var p Product
	json.Unmarshal(bodyy, &p)
	fmt.Println(p)
}
