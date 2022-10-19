package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	stats "task/analysis"
	order "task/generateIds"
)

func placeOrder() {
	file, e := os.OpenFile("restaurant.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Fatalln("Fail to open file")
	}

	log.SetOutput(file)

	eater_id := order.Generate_eater_id()
	foodMenu := order.Generate_foodmenu()

	for _, foodMenuId := range foodMenu {
		log.Print(eater_id, foodMenuId)
	}

	defer file.Close()
}

func main() {
	fmt.Println("Welcome to Restaurent")
	fileName := "restaurant.log"

	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	//just to fill logs in a file
	if len(bs) < 1 {
		for i := 0; i < 20; i++ {
			placeOrder()
		}
	}

	itemsList := stats.TopItemsConsumed(3)
	fmt.Println("Top 3 items consumed are :")
	for count, list := range itemsList {
		fmt.Println(count+1, ") ", list[0], "consumed ", list[1], " times")
	}

	ok, e := stats.CheckDuplicates()
	if ok {
		fmt.Println("Error :", e)
	} else {
		fmt.Println("No duplicates in log ")
	}
}
