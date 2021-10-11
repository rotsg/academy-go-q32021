package universities

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/rotsg/academy-go-q32021/model"
)

const api = "http://universities.hipolabs.com/search?country=Mexico"

func GetUniversities() {
	resp, err := http.Get(api)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var jsonData []model.University
	err = json.Unmarshal(bodyBytes, &jsonData)
	if err != nil {
		fmt.Println(err)
	}

	csvFile, err := os.Create("./data/universities.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	for index, university := range jsonData {
		var row []string
		row = append(row, strconv.Itoa(index+1))
		row = append(row, university.Name)
		row = append(row, university.Country)
		row = append(row, university.WebPage[0])
		writer.Write(row)
	}
	writer.Flush()
}
