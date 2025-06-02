package main

import "fmt"

const ebulicaoF = 212.0

func main() {

	tempF := ebulicaoF
	tempC := (tempF - 32) * 5 / 9
	fmt.Println("A temperatura de ebulicao da agua em F =", tempF)
	fmt.Println("A temperatura de ebulicao da agua em C =", tempC)

}
