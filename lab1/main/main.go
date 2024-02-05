package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := http.Client{}

	url := "http://109.167.241.225:8001/http_example/give_me_five?student=16&wday=2"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Set("REQUEST_AGENT", "ITMO student")
	req.Header.Set("COURSE", "Protocols")

	fmt.Println("Запрос: ", req)

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Ответ: ", response)
}
