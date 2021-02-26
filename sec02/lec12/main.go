// Section 02 - Lecture 12 : Packages - Part 2
package main

import (
	"fmt"

	_ "github.com/golang-excercise/sec02/lec12/logger"
	_ "github.com/golang-excercise/sec02/lec12/math"
	log1 "github.com/sirupsen/logrus"
)

func init() {
	fmt.Println("main.init() - 1 from main.go")
}

func init() {
	fmt.Println("main.init() - 2 from main.go")
}
func main() {
	log1.Warn("my error message")
	log1.Error("my error message")
}