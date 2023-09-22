package main

import (
	"21.09/user"
	"21.09/user/student"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var st user.User = &student.Student{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Student name: ")

	name, _ := reader.ReadString('\n')

	st.SetName(name)

	fmt.Println(st.GetName())
}
