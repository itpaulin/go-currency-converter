package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetRates()  {
	endpoint := API_URL + "latest.json?app_id=" + APP_ID

	response, err := http.Get(endpoint)

	if err != nil {
		log.Fatalf("Occured an error: %s\n", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Occured an error: %s\n", err)

	}

	var data map[string]interface{}
	errJsonParse := json.Unmarshal(body, &data)
	if errJsonParse != nil {
		log.Fatalf("Occured an error when decoding JSON: %s\n", errJsonParse)
	}

	//TODO pick specific rates from the data
	rates := data["rates"].(map[string]interface{})

	fmt.Println(rates["BRL"])
	


}