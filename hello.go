package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func ChangeDirectory(inputs []string) error {
	if len(inputs) < 2 {
		return errors.New("Add path!")
	}
	return os.Chdir(inputs[1])
}

func SuperUser(inputs []string) {
	cmd := exec.Command("su")
	cmd.Stdin = strings.NewReader("ubuntu" + "\n")
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("\x1b[31m", "(#`-´)<ENTER PASSWORD!!!!! : ")
	//pass, _ := reader.ReadString('\n')

	//cmd := exec.Command(inputs[0])
	//cmd.Stdin = strings.NewReader("ubuntu" + "\n")

	//return cmd.Run()
}

func exe(input string) error {
	input = strings.TrimRight(input, "\n")
	inputs := strings.Split(input, " ")

	if inputs[0] == "cd" {
		err := ChangeDirectory(inputs)
		return err
	}

	if inputs[0] == "su" {
		SuperUser(inputs)
	}

	if len(inputs) == 1 {
		cmd := exec.Command(inputs[0])
		result, err := cmd.Output()
		fmt.Println(string(result))
		return err
	} else {
		cmd := exec.Command(inputs[0], inputs[1])
		result, err := cmd.Output()
		fmt.Println(string(result))
		return err
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		//dir, err := os.Getwd()
		user, err := user.Current()

		if err != nil {
			panic(err)
		}

		if user.Uid == "0" {
			fmt.Print("\x1b[31m", "(#`-´)< ")
		} else {
			fmt.Print("\x1b[32m", "(#`-´)< ")
		}
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = exe(input); err != nil {
			fmt.Fprintln(os.Stdout, "ERROR")
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
