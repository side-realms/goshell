package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func exe(input string) {
	input = strings.TrimRight(input, "\n")
	out, err := exec.Command(input).Output()
	if err != nil {
		fmt.Fprintln(os.Stdout, "ERROR:")
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Fprintln(os.Stdout, string(out))
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		exe(input)
		//fmt.Fprintln(os.Stdout, input+"keroro")
	}
}
