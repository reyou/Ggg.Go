package main

import (
	"fmt"

	magazine "./magazine"
)

func main() {
	address := magazine.Address{Street: "123 Oak St", City: "Omaha", State: "NE", PostalCode: "68111"}
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.HomeAddress = address
	fmt.Println(subscriber.HomeAddress)
	fmt.Printf("%#v\n", subscriber.HomeAddress)
}
