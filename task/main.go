package main

import (
	"fmt"
	"strings"
)

type Battery string

func (b *Battery) String() string {
	newBattery := "["
	for i := 0; i < strings.Count(string(*b), "0"); i++ {
		newBattery += " "
	}
	for i := 0; i < strings.Count(string(*b), "1"); i++ {
		newBattery += "X"
	}
	newBattery += "]"
	*b = Battery(newBattery)
	return string(*b)
}

func main() {
	var batteryForTest Battery
	_, _ = fmt.Scan(&batteryForTest)
	batteryForTest.String()
	fmt.Print(batteryForTest)
}
