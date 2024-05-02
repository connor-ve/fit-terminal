/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
)

type log struct {
	Date    string `json:"date"`
	Workout []struct {
		Exercise int `json:exercise`
		Set      struct {
		} `json:set`
	} `json:"workout"`
}

func main() {
	fmt.Println("Hello")
}

func pullWorkouts() {

}
