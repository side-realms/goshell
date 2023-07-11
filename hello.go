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
	os.Chdir(inputs[1])
	dir, err := os.Getwd()
	fmt.Println("\x1b[32m", "(　³ω³ )< "+dir)
	return err
}

func SuperUser(inputs []string) error {
	cmd := exec.Command("su")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Err", err)
	} else {
		if string(out) == "" {
			fmt.Println("OUT:success")
		} else {
			fmt.Println("OUT:", string(out))
		}
	}
	return err
}

func SuDo(inputs []string) error {
	if inputs[1] == "apt" {
		inputs[1] = "apt-get"
	}
	cmd := exec.Command("sudo", inputs[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Err", err)
	} else {
		if string(out) == "" {
			fmt.Println("OUT:success")
		} else {
			fmt.Println("OUT:", string(out))
		}
	}
	return err
}

func exe(input string) error {
	input = strings.TrimRight(input, "\n")
	inputs := strings.Split(input, " ")

	if inputs[0] == "cd" {
		err := ChangeDirectory(inputs)
		return err
	}

	if inputs[0] == "su" {
		err := SuperUser(inputs)
		return err
	}

	if inputs[0] == "sudo" {
		err := SuDo(inputs)
		return err
	}

	if len(inputs) == 1 {
		cmd := exec.Command(inputs[0])
		result, err := cmd.Output()
		fmt.Println(string(result))
		return err
	} else {
		cmd := exec.Command(inputs[0], inputs[1:]...)
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
