package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var userId string
	fmt.Println("Enter User ID")
	fmt.Scan(&userId)
	url := "https://reqres.in/api/users/" + userId
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error handling Code for Get")
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		var data map[string]interface{}
		err := json.Unmarshal([]byte(body), &data)
		if err != nil {
			panic(err)
		}
		fmt.Println("String ---> ", string(body))
		fmt.Println("Map ---> ", data["data"])
	}

}
