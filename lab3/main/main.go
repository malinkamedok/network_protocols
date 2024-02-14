package main

import (
	"fmt"
	"net"
	"strings"
)

var msg string = "3Student 1 2024-02-14-15-47-49,635 0,017876 Student 2 2024-02-14-15-47-49,635 0,201170 Student 3 2024-02-14-15-47-49,635 0,852571 Student 4 2024-02-14-15-47-49,635 0,690648 Student 5 2024-02-14-15-47-49,635 0,678021 Student 6 2024-02-14-15-47-49,635 0,743635 Student 7 2024-02-14-15-47-49,635 0,246680 Student 8 2024-02-14-15-47-49,635 0,215070 Student 9 2024-02-14-15-47-49,635 0,787391 Student10 2024-02-14-15-47-49,635 0,264731 Student11 2024-02-14-15-47-49,635 0,221968 Student12 2024-02-14-15-47-49,635 0,923678 Student13 2024-02-14-15-47-49,635 0,074869 Student14 2024-02-14-15-47-49,635 0,661677 Student15 2024-02-14-15-47-49,635 0,649666 Student16 2024-02-14-15-47-49,635 0,779076 Student17 2024-02-14-15-47-49,635 0,975931 Student18 2024-02-14-15-47-49,635 0,841937 Student19 2024-02-14-15-47-49,635 0,562847 Student20 2024-02-14-15-47-49,635 0,706869 Student21 2024-02-14-15-47-49,635 0,639830 Student22 2024-02-14-15-47-49,635 0,797632 Student23 2024-02-14-15-47-49,635 0,205685 Student24 2024-02-14-15-47-49,6"

func main() {
	conn, err := net.Dial("tcp", "109.167.241.225:6340")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected")

	buffer := make([]byte, 1024)

	var bufferResult []string

	for i := 0; i < 10; i++ {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		bufferResult = append(bufferResult, string(buffer[:n]))

		// Process and use the data (here, we'll just print it)
		//fmt.Printf("Received: %s\n", buffer[:n])
	}

	var result []string

	//result1 := string(buffer)

	for i := 0; i < 10; i++ {

		result1 := bufferResult[i]

		if strings.Contains(result1, "Student16") {
			//fmt.Println(strings.Index(result1, "Student16"))
			//fmt.Println(result1[646:])
			before, _, found := strings.Cut(result1[646:], " Student17")
			if found {
				fmt.Println(before)
				result = append(result, before)
			} else {
				fmt.Println("not found")
			}
		}
	}

	fmt.Println(result)

	//fmt.Println(result1)
}

//3Student 1 2024-02-14-15-47-49,635 0,017876 Student 2 2024-02-14-15-47-49,635 0,201170 Student 3 2024-02-14-15-47-49,635 0,852571 Student 4 2024-02-14-15-47-49,635 0,690648 Student 5 2024-02-14-15-47-49,635 0,678021 Student 6 2024-02-14-15-47-49,635 0,743635 Student 7 2024-02-14-15-47-49,635 0,246680 Student 8 2024-02-14-15-47-49,635 0,215070 Student 9 2024-02-14-15-47-49,635 0,787391 Student10 2024-02-14-15-47-49,635 0,264731 Student11 2024-02-14-15-47-49,635 0,221968 Student12 2024-02-14-15-47-49,635 0,923678 Student13 2024-02-14-15-47-49,635 0,074869 Student14 2024-02-14-15-47-49,635 0,661677 Student15 2024-02-14-15-47-49,635 0,649666 Student16 2024-02-14-15-47-49,635 0,779076 Student17 2024-02-14-15-47-49,635 0,975931 Student18 2024-02-14-15-47-49,635 0,841937 Student19 2024-02-14-15-47-49,635 0,562847 Student20 2024-02-14-15-47-49,635 0,706869 Student21 2024-02-14-15-47-49,635 0,639830 Student22 2024-02-14-15-47-49,635 0,797632 Student23 2024-02-14-15-47-49,635 0,205685 Student24 2024-02-14-15-47-49,6
