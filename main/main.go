package main

import "fmt"

func main() {

	kek := new(RegHand)

	kek.NetworkStatus.GeneralNetCheck("com", "80")

	fmt.Println(kek.NetworkStatus)

}
