package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"parking_lot/lib"
	"strings"
)

func main() {
	args := os.Args

	var res string
	cmd := lib.GetCommandExecuter()
	if len(args) == 2 {

		path := "/usr/src/fixures/" + args[1]

		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			//fmt.Println(scanner.Text())

			command := strings.Split(scanner.Text(), " ")
			//spew.Dump(command)
			res, _ = cmd.Execute(command[0], command[1:]...)
			//res, _ = cmd.Execute(command...)
			fmt.Println(res)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		return
	}

	res, _ = cmd.Execute("create_parking_lot", "6")
	fmt.Println("Result :", res)

	res, _ = cmd.Execute("park", "KA-01-HH-1234", "White")
	fmt.Println("Result :", res)
	res, _ = cmd.Execute("park", "KA-01-HH-9999", "White")
	fmt.Println("Result :", res)
	res, _ = cmd.Execute("park", "KA-01-BB-0001", "Black")
	fmt.Println("Result :", res)
	res, _ = cmd.Execute("park", "KA-01-HH-7777", "Red")
	fmt.Println("Result :", res)
	res, _ = cmd.Execute("park", "KA-01-HH-2701", "Blue")
	fmt.Println("Result :", res)
	res, _ = cmd.Execute("park", "KA-01-HH-3141", "Black")
	fmt.Println("Result :", res)

	res, _ = cmd.Execute("leave", "4")
	fmt.Println("Result :", res)

	res, _ = cmd.Execute("status")
	fmt.Println("Result :", res)

	res, _ = cmd.Execute("park", "KA-01-P-3331", "White")
	fmt.Println("Result :", res)

	res, _ = cmd.Execute("park", "DL-12-AA-9999", "White")
	fmt.Println("Result :", res)

	res, _ = cmd.Execute("registration_numbers_for_cars_with_colour", "White")
	fmt.Println("Result :", res)

	res, _ = cmd.Execute("slot_numbers_for_cars_with_colour", "White")
	fmt.Println("Result :", res)

	res, _ = cmd.Execute("slot_number_for_registration_number", "KA-01-HH-3141")
	fmt.Println("Result :", res)

	res, _ = cmd.Execute("slot_number_for_registration_number", "MH-04-AY-1111")
	fmt.Println("Result :", res)
	fmt.Println("Executing done")
}
