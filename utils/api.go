package adv

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var SaveCity string = ""
var SaveC []string

func Save() {

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", "https://base.reactivaservicios.com/api/database/fields/table/412/", nil)

	// Headers
	req.Header.Add("Authorization", "Token LvPrh0HN4Ll6N0ZwIRxrRgiYVDSBbtWN")

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	//fmt.Println("response Status : ", resp.Status)
	//fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))

	fmt.Println(SaveCity)
}
