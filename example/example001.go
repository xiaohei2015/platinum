package main

import (
	"fmt"

	"github.com/xiaohei2015/platinum"
)

func main() {
	fmt.Println("Now:", platinum.Now())
	fmt.Println("Right now:", platinum.Now().ToDateTimeString())
	fmt.Println("Date:", platinum.CreateFromDate(2018, 10, 31))
}
