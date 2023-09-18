package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	result, err := exec.Command("go", "run", "../Investigo/investigo.go", "fastfir").Output()
	if err != nil {
		print(err)
	}
	parsed := string(result)
	array := strings.Split(parsed, "fastfir")
	array1 := []string{}
	for _, site := range array {
		array1 = append(array1, strings.Split(site, ":")[0])
	}
	final := array1[3:]
	for index, website := range array1 {
		final = append(final, string(website[index])[4:])
	}
	fmt.Print(final)
}
