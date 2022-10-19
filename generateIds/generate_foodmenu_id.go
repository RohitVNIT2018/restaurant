package order

import (
	"math/rand"
	"time"
)

var FoodList = make(map[string]string)

func Generate_foodmenu() []int {

	curryList := []int{301, 302, 303, 304, 305, 306}
	chapatiList := []int{201, 202, 203}
	riceList := []int{101, 102, 103}

	foodMenu := [][]int{curryList, chapatiList, riceList}

	foodMenuIds := make([]int, 0)
	random_source := rand.NewSource(time.Now().UnixNano())
	random_seed := rand.New(random_source)
	for i, l := range foodMenu {
		if i == 0 {
			itemIndex := random_seed.Intn(len(l) - 1)

			itemId := l[itemIndex]
			foodMenuIds = append(foodMenuIds, itemId)
		}
		itemIndex := random_seed.Intn(len(l) - 1)

		itemId := l[itemIndex]
		foodMenuIds = append(foodMenuIds, itemId)

	}

	return foodMenuIds

}
