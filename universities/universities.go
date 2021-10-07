package universities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rotsg/academy-go-q32021/model"
)

const api = "http://universities.hipolabs.com/search?country=Mexico"

func GetUniversities() {
	resp, err := http.Get(api)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var todoStruct model.University
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct %+v\n", todoStruct)
}
