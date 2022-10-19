package analysis

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func CheckDuplicates() (bool, error) {

	fileName := Filename
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	text := string(bs)

	//remove date and time logs
	re := regexp.MustCompile(`\d{1,4}([.\-/])\d{1,2}([.\-/])\d{1,4}`)
	filtered_text := re.ReplaceAllString(text, "")
	re = regexp.MustCompile(`\d{2}:\d{2}:\d{2}`)
	filtered_text = re.ReplaceAllString(filtered_text, "")
	filtered_list := strings.Split(filtered_text, "\n")
	first_row := strings.TrimSpace(filtered_list[0])
	first_id_list := strings.Split(first_row, " ")
	first_eater_id := first_id_list[0]

	//foodItemId => eaterId
	mapDuplicate := make(map[string]string, 0)
	for _, rows := range filtered_list {
		if len(rows) > 0 {
			trim_row := strings.TrimSpace(rows)
			list := strings.Split(trim_row, " ")
			if first_eater_id == list[0] {
				if _, ok := mapDuplicate[list[1]]; !ok {
					mapDuplicate[list[1]] = list[0]
				} else {
					return true, fmt.Errorf("duplicate eater_Id %v with menuId %v", list[0], list[1])
				}
			} else {
				mapDuplicate = make(map[string]string, 0)
				first_eater_id = list[0]
				mapDuplicate[list[1]] = list[0]
			}
		}
	}
	return false, nil
}
