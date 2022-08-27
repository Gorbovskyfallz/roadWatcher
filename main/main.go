package main

import (
	"fmt"
	"github.com/moby/sys/mountinfo"
)

func main() {
	kek, _ := mountinfo.Mounted("/media/passed3")
	fmt.Println(kek)
}
