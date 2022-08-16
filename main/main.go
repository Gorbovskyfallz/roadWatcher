package main

import "fmt"

func main() {

	kek := new(RegHand)

	kek.NetworkStatus.ModemNetCheck("com", "80")

	fmt.Println(kek.NetworkStatus)

}
