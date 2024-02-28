package main

import (
	"fmt"
	"github.com/goburrow/modbus"
	"log"
	"time"
)

//func simon() {
//	var client *modbus.ModbusClient
//
//	client, err := modbus.NewClient(
//		&modbus.ClientConfiguration{
//			URL:     "tcp://109.167.241.225:601",
//			Timeout: 1 * time.Second,
//		})
//	if err != nil {
//		fmt.Println("Error with server: ", err)
//		return
//	}
//
//	err = client.Open()
//	if err != nil {
//		fmt.Println("Error with opening client: ", err)
//		return
//	}
//	fmt.Println("time      : ", time.Now().UTC())
//	regs := make([]uint16, 2)
//	regs, err = client.ReadRegisters(100*16, 2, modbus.HOLDING_REGISTER)
//	if err != nil {
//		fmt.Println("Error with reading regs: ", err)
//		return
//	}
//	regs2 := make([]byte, 10)
//	regs2, err = client.ReadBytes(100*16+2, 10, modbus.HOLDING_REGISTER)
//	fmt.Println("first: ", regs[0])
//	fmt.Println("second: ", regs[1])
//	fmt.Println("string: ", string(regs2))
//}

func main() {
	// Создаем TCP клиент для Modbus
	handler := modbus.NewTCPClientHandler("109.167.241.225:601")

	// Устанавливаем таймаут для подключения
	handler.Timeout = 5 * time.Second

	// Создаем клиент
	client := modbus.NewClient(handler)

	// Подключаемся к устройству
	err := handler.Connect()
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer handler.Close()

	// Номер студента
	var studentNumber uint16 = 16

	// Адрес регистра, к которому мы обращаемся
	startAddress := studentNumber * 100

	// Читаем два регистра (16 бит каждый)
	results, err := client.ReadHoldingRegisters(startAddress, 2)
	if err != nil {
		log.Fatalf("Failed to read holding registers: %v", err)
	}

	// Выводим результаты
	fmt.Printf("Registers at address %v: %v\n", startAddress, results)

	// Читаем строку ASCII (8 бит)
	startAddress += 2                                                 // Сдвигаемся на 2 регистра вперед
	asciiResults, err := client.ReadHoldingRegisters(startAddress, 5) // 5 регистров для 10 символов ASCII
	if err != nil {
		log.Fatalf("Failed to read ASCII string: %v", err)
	}

	// Выводим строку ASCII
	fmt.Printf("ASCII string at address %v: %v\n", startAddress, string(asciiResults))

}
