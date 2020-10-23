package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

func main() {
	// for loop that just displays the list/options
	// list, add, remove, edit, list is numbered
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	tp := user.HomeDir + "/todo.txt"

	list := readData(tp)
	var option int
	for {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
		printList(list)
		fmt.Println("1. New item | 2. Remove item | 3. Edit item | 4. Edit item order | 5. Exit")
		fmt.Scan(&option)
		switch {
		case option == 1:
			list = newItem(list)
		case option == 3:
			list = editItem(list)
		case option == 2:
			list = removeItem(list)
		case option == 4:
			list = changeOrder(list)
		case option == 5:
			return
		default:
			fmt.Println("Please enter a valid option (1-6)")
		}
		writeData(tp, list)
	}
}

func changeOrder(list []string) []string {
	var input string
	fmt.Println("What numbers do you want to swap? (separate with comma")
	printList(list)
	fmt.Scanln(&input)
	sl := strings.Split(input, ",")
	var c []int
	for _, i := range sl {
		j, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("Invalid input")
			return list
		}
		c = append(c, j)
	}

	list[c[0]], list[c[1]] = list[c[1]], list[c[0]]

	// code to insert instead of swap
	// val := list[c[0]]
	// list = append(list[:c[1]+1], list[c[1]:]...)
	// list[c[1]] = val

	return list
}

func removeItem(list []string) []string {
	var input int
	fmt.Println("What item do you want to remove? (invalid input to cancel)")
	printList(list)
	_, err := fmt.Scan(&input)
	if err != nil {
		return list
	}
	return append(list[:input], list[input+1:]...)
}

func editItem(list []string) []string {
	var input int
	var strInput string
	fmt.Println("What item do you want to edit?")
	printList(list)
	fmt.Scan(&input)
	if input > len(list) {
		fmt.Println("Invalid selection!")
		return list
	}
	fmt.Println("What do you want to change it to? (blank to cancel)")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		strInput = scanner.Text()

	}

	if strInput == "" {
		return list
	}

	list[input] = strInput
	return list
}

func newItem(list []string) []string {
	var strInput string
	fmt.Println("What is the new item?")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		strInput = scanner.Text()

	}
	return append(list, strInput)
}

func printList(list []string) {
	fmt.Println("")
	for num, item := range list {
		fmt.Println(num, ". ", item)
	}
	fmt.Println("")
}

func readData(filename string) []string {
	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file")
		fmt.Println("Creating new todo-list")
		return []string{}
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println("Error reading lines")
	}

	var lfix []string

	for _, val := range lines {
		lfix = append(lfix, val[0])
	}

	fmt.Println("Read sucessful\n First line:", lines[0])
	return lfix
}

func writeData(filename string, data []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file!")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	var fix [][]string
	for _, val := range data {
		fix = append(fix, []string{val})
	}

	err = writer.WriteAll(fix)
	if err != nil {
		fmt.Println("Error writing to file!")
	}

}
