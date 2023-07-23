package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("python", "test_py/test.py", "-w car")
	//fmt.Println(cmd.Args)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	test := string(out)
	slice := strings.Split(test, " ")
	fmt.Println(slice, len(slice), cap(slice))
	slices := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slices, len(slices), cap(slices))

}
