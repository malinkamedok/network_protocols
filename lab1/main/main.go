package main

import (
	"fmt"
	"io"
	"net/http"
)

func sendAndGet() {
	client := http.Client{Transport: &http.Transport{DisableCompression: true}}
	//client := http.Client{}

	url := "http://109.167.241.225:8001/http_example/give_me_five?student=16&wday=3"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Set("REQUEST_AGENT", "ITMO student")
	req.Header.Set("COURSE", "Protocols")
	req.Header.Set("User-Agent", "")

	fmt.Println("Запрос: ", req)

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Ответ: ", response)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Текст ответа: ", string(body))
}

func postman() {
	url := "http://109.167.241.225:8001/http_example/give_me_five?wday=3&student=16"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("REQUEST_AGENT", "ITMO student")
	req.Header.Add("COURSE", "Protocols")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

func main() {
	sendAndGet()
}
