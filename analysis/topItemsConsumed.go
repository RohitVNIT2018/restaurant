package analysis

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

type FoodItemFreq struct {
	foodItemFreq string
	freq         int
}

func (f FoodItemFreq) String() string {
	return fmt.Sprintf("%s %d", f.foodItemFreq, f.freq)
}

type byFreq []FoodItemFreq

func (a byFreq) Len() int           { return len(a) }
func (a byFreq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFreq) Less(i, j int) bool { return a[i].freq < a[j].freq }

var FoodList = make(map[string]string)
var Filename = "restaurant.log"

func init() {
	// initialize menu with ids
	FoodList["101"] = "Steam Rice"
	FoodList["102"] = "Jira Rice"
	FoodList["103"] = "Plan Rice"
	FoodList["201"] = "Tava Roti"
	FoodList["202"] = "Nan Roti"
	FoodList["203"] = "Tandoori Roti"
	FoodList["301"] = "Masala Brinjal Curry"
	FoodList["302"] = "Paneer Curry"
	FoodList["303"] = "Egg Curry"
	FoodList["304"] = "Chiken Curry"
	FoodList["305"] = "Daal Makhni"
	FoodList["306"] = "Daal Tadka"
}
func TopItemsConsumed(n int) [][]string {

	fileName := Filename
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	text := string(bs)

	//filter log data
	re := regexp.MustCompile(`\d{1,4}([.\-/])\d{1,2}([.\-/])\d{1,4}`)
	filtered_text := re.ReplaceAllString(text, "")
	re = regexp.MustCompile(`\d{2}:\d{2}:\d{2}`)
	filtered_text = re.ReplaceAllString(filtered_text, "")
	re = regexp.MustCompile("[0-9][0-9][0-9]")
	matches := re.FindAllString(filtered_text, -1)

	foundIds := make(map[string]int)
	for _, match := range matches {
		foundIds[match]++
	}

	var foodItemFreqs []FoodItemFreq
	for k, v := range foundIds {
		foodItemFreqs = append(foodItemFreqs, FoodItemFreq{k, v})
	}

	sort.Sort(sort.Reverse(byFreq(foodItemFreqs)))

	i := 0
	count := 0
	result := [][]string{}
	for count < n {
		item := strings.Split(foodItemFreqs[i].String(), " ")
		if itemName, found := FoodList[string(item[0])]; found {
			result = append(result, []string{itemName, item[1]})
			count++
		}
		i++
	}
	return result
}
