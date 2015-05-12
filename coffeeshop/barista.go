package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var custUrl, coffeeName, cream, sugar, xHot, size string

func checkCustomer() {

	for {
		var customer map[string]string

		url := "http://localhost:1330/dequeueOrderQueue"
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(body, &customer)
		if err != nil {
			panic(err)
		}

		fmt.Println(customer["ip"], customer["port"], customer["name"])

		if customer["port"] != "" {
			custUrl = "http://" + customer["ip"] + customer["port"] + "/"

			k := GetCoffeeName()
			WantCream()
			WantSugar()

			if k != "frappechino" {
				Extra_hot()
			}

			Whatsize()

			RepeatOrder()

			TakeCoffee()
		} else {
			fmt.Println("No one in Order Queue")
		}
		time.Sleep(3 * time.Second)
	}
}

func GetCoffeeName() string {

	fmt.Println("How may I help you?")

	url := custUrl + "coffee_name"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	coffeeName = string(body)
	fmt.Println("Okay, you want " + string(body))

	return coffeeName
}

func WantCream() {

	fmt.Println("Extra cream?")

	url := custUrl + "cream"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	cream = string(body)
	fmt.Println("Okay, you want " + string(body))

}

func WantSugar() {

	fmt.Println("How much sugar?")

	url := custUrl + "sugar"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	sugar = string(body)
	fmt.Println("Okay, you want " + string(body))

}

func Extra_hot() {

	fmt.Println("Extra hot?")

	url := custUrl + "extra_hot"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	xHot = string(body)
	fmt.Println("Okay, you want " + string(body))

}

func Whatsize() {

	fmt.Println("What size would you like?")

	url := custUrl + "size"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	size = string(body)
	fmt.Println("Okay, you want " + string(body))

}

func RepeatOrder() {

	fmt.Printf("Let me repeat your order\n You ordered a " + size + " " + coffeeName + " ")
	if cream == "yes" {
		fmt.Printf(" with cream")
	}
	if cream == "yes" {
		fmt.Printf(", Extra hot")
	}
	if sugar == "yes" {
		fmt.Printf("and Sweetend.")
	}

	url := custUrl + "confirmOrder"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("Will get your coffee ready!!")
	time.Sleep(10 * time.Second)

}

func TakeCoffee() {

	fmt.Println("Your coffee is ready. Enjoy!")

	url := custUrl + "takeCoffee"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	time.Sleep(2 * time.Second)

}

// func checkCustomer() {

// 	var customer map[string]string

// 		url :="http://localhost:1330/dequeueOrderQueue"
// 	res, err :=http.Get(url)
// 	if err !=nil{
// 		panic(err)
// 	}
// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err !=nil{
// 		panic(err)
// 	}

// 	err = json.Unmarshal(body, &customer)
// 	if err !=nil{
// 		panic(err)
// 	}

// 	fmt.Println(customer["ip"], customer["port"], customer["name"])

// 	time.Sleep(3*time.Second)

// }

func main() {

	checkCustomer()

}
