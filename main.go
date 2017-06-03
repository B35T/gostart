package main

import (
	"gostart/src/controllers/commands"
	"fmt"
	"log"
	"os/exec"
)

func main()  {

	var text string
	fmt.Scanf("%v" , &text)
	//fmt.Println(text)
	control(text)
}

func control(text string) {
	if text == "help" || text == "doc" {
		fmt.Println("help or doc")
		fmt.Println("c : Create")
		fmt.Println("pwd")
		fmt.Println("ls")
		fmt.Println("help or doc")

		main()

	} else if text == "c" {

		fmt.Print("Enter folder name : ")
		t := ""
		fmt.Scanf("%v",&t)

		output,err :=  commands.Create(t)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(output)

		fmt.Print("Enter Project name : ")
		fmt.Scanf("%v",&t)
		a, b , _ := commands.GoStartMVC(t)
		fmt.Println(a)
		fmt.Println(b)

		main()

	} else if text == "env" {
		output, _ := commands.GoENV()
		fmt.Println(output)

		main()
	} else if text == "pwd" {
		pwd, _ := commands.GetPWD()
		fmt.Println(pwd)

		main()

	} else if text == "ls" {
		ls := exec.Command("ls")
		i, _ := ls.CombinedOutput()
		fmt.Println(string(i))

		main()
	} else {
		fmt.Println("help")
		main()
	}
}