package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func RestApiGetFunc() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Println(err)
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)

		// İkinci yöntem yapılara atamak
		var varTodo Todo

		json.Unmarshal(bodyBytes, &varTodo) // json kütüphanesi ile kolayca her iki yönde dönüştürme yapabiliriz.
		fmt.Println(varTodo)
	}
	defer resp.Body.Close() // açılan bağlantının fonksiyondan çıkılmadan önce kapatılması gerekir her ne olursa olsun.
}

func RestApiPostFunc() {
	varTodo := Todo{UserId: 1, Id: 1, Title: "Hesap Bilgisi", Body: "Bu kullanıcı yeni kayıt olmuştur"}
	jsonTodo, jsonErr := json.Marshal(varTodo)

	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	resp, httpErr := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo))

	if httpErr != nil {
		fmt.Println(httpErr)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todoResponse Todo
	json.Unmarshal(bodyBytes, &todoResponse)
	fmt.Println(todoResponse)

	defer resp.Body.Close()
}
