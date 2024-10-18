package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	//input := bufio.NewScanner(os.Stdin)
	//fields := []string{"Hostname", "Username", "Password"}
	//args := make([]string, 3)
	//count := 0
	//fmt.Println(fields[count])
	//for input.Scan() && count < 2 {
	//	args[count] = input.Text()
	//	count++
	//	fmt.Println(fields[count])
	//}
	//cmd := exec.Command("ssh", fmt.Sprintf("%s@%s", args[1], args[0]))
	//cmd.Stdout = os.Stdout
	//if err := cmd.Run(); err != nil {
	//	log.Print(err)
	//}
	//fmt.Println("End")
	conn, err := net.Dial("tcp", "139.162.27.190:22")
	result := make([]byte, 300)
	conn.Read(result)
	fmt.Println(string(result))
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(conn)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Fprintln(conn, "ls")
	}
}
