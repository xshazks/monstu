package wahttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func PostStructtoAPI(wamsg interface{}, url string) (result interface{}) {
	mJson, _ := json.Marshal(wamsg)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(mJson))
	if err != nil {
		fmt.Println("Could not make POST request to whatsauth")
	}
	fmt.Println(resp)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error unmarshaling data from request.")
	}
	json.Unmarshal([]byte(body), &result)
	return result
}
