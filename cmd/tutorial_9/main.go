// more relevant example
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_CHICKEN_PRICE float32 = 5 // maximum price for chicken
const MAX_TOFU_PRICE float32 = 3    // maximum price for tofu

func main() {
	var chickenChannel = make(chan string)                                 // create a channel for strings
	var tofuChannel = make(chan string)                                    // create a channel for strings
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"} // list of websites

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel) // start a goroutine to check chicken prices
		go checkTofuPrices(websites[i], tofuChannel)       // start a goroutine to check tofu prices
	}
	sendMessage(chickenChannel, tofuChannel) // start the function to send messages
}

func checkChickenPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice < MAX_CHICKEN_PRICE {
			c <- website
			break
		}
	}
}

func checkTofuPrices(website string, c chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice < MAX_TOFU_PRICE {
			c <- website
			break
		}
	}
}

func sendMessage(chickenChannel, tofuChannel chan string) {
	select {
	case chickenWebsite := <-chickenChannel: // receive from chicken channel
		fmt.Printf("Text sent: Deal on chicken is available at %v\n", chickenWebsite)
	case tofuWebsite := <-tofuChannel: // receive from tofu channel
		fmt.Printf("Email sent: Deal on tofu is available at %v\n", tofuWebsite)
	}
}
