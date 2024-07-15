package main

import (
	"fmt"
)

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса
	doTask(value1, value2, operation)
}

func doTask(value1, value2, operation interface{}) {
	switch value1.(type) {
	case float64:
		switch value2.(type) {
		case float64:
			switch operation.(type) {
			case string:
				switch operation {
				case "+":
					fmt.Printf("%.4f", value1.(float64)+value2.(float64))
				case "-":
					fmt.Printf("%.4f", value1.(float64)-value2.(float64))
				case "/":
					fmt.Printf("%.4f", value1.(float64)/value2.(float64))
				case "*":
					fmt.Printf("%.4f", value1.(float64)*value2.(float64))
				default:
					fmt.Println("неизвестная операция")
					return
				}
			}
		default:
			fmt.Printf("value=%v: %T", value2, value2)
			return
		}
	default:
		fmt.Printf("value=%v: %T", value1, value1)
		return
	}
}
