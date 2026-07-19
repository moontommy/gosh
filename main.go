package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const version = "0.0.1"

var modulePath = "unknown"

func printBanner() {
	banner := "===============================\n"
	banner += "   _____ _____ _____ _         \n"
	banner += "   |  __ \\  _  /  ___| |       \n"
	banner += "   | |  \\/ |/' \\ `--.| |__     \n"
	banner += "   | | __|  /| |`--. \\ '_ \\    \n"
	banner += "   | |_\\ \\ |_/ /\\__/ / | | |   \n"
	banner += "   \\____/\\___/\\____/|_| |_|    \n"
	banner += "                               \n"
	banner += "   Welcome to G0Sh shell     \n"
	banner += fmt.Sprintf("   version %s        \n", version)
	banner += fmt.Sprintf("   %s        \n", modulePath)
	banner += "===============================\n"
	fmt.Print(banner)
}

func handleBuiltin(args []string) bool {
	switch args[0] {
	case "exit":
		fmt.Print("Goodbye")
		os.Exit(0)
	case "cd":
		os.Chdir(args[1])
		return true
	}

	return false
}

func main() {
	printBanner()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(fmt.Sprintf("G0Sh-%s > ", version))

		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		input = strings.TrimSpace(input)

		// Built-in command
		if handleBuiltin(strings.Fields(input)) {
			continue
		}

		// Run external command
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
