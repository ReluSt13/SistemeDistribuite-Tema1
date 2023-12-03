package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/rpc"
	"strings"
)

type Exercitiul1Args struct {
	Words []string
}

type Exercitiul2Args struct {
	Words []string
}

type Exercitiul3Args struct {
	Numbers []int
}

type Exercitiul4Args struct {
	Numbers []int
}

type Exercitiul5Args struct {
	Strings []string
}

type Exercitiul6Args struct {
	Strings   []string
	Shift     int
	Direction string
}

type Exercitiul7Args struct {
	EncodedString string
}

type Exercitiul8Args struct {
	Numbers []int
}

type Exercitiul9Args struct {
	Strings []string
}

type Exercitiul10Args struct {
	Strings []string
}

type Exercitiul11Args struct {
	Numbers []int
	K       int
}

type Exercitiul12Args struct {
	Numbers []int
}

type Exercitiul13Args struct {
	Numbers []int
}

type Exercitiul14Args struct {
	Strings []string
}

type Exercitiul15Args struct {
	Characters []string
}

func main() {
	rpc_conn, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	msg_conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	// send client name to server
	_, err = msg_conn.Write([]byte("Client_A\n"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// listen for messages from the server
	go func() {
		reader := bufio.NewReader(msg_conn)
		for {
			msg, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			}
			// print the message trimmed
			log.Println(strings.TrimSpace(msg))
		}
	}()

	var ex1_reply []string
	log.Printf("Calling Exercitiul1 with data: %v\n", []string{"casa", "masa", "trei", "tanc", "4321"})
	err = rpc_conn.Call("Server.Exercitiul1", &Exercitiul1Args{Words: []string{"casa", "masa", "trei", "tanc", "4321"}}, &ex1_reply)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("Server responded with: %s\n", ex1_reply)
	}

	var ex2_reply int
	log.Printf("Calling Exercitiul2 with data: %v\n", []string{"abd4g5", "1sdf6fd", "fd2fdsf5"})
	err = rpc_conn.Call("Server.Exercitiul2", &Exercitiul2Args{Words: []string{"abd4g5", "1sdf6fd", "fd2fdsf5"}}, &ex2_reply)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("Server responded with: %d\n", ex2_reply)
	}

	var ex3_reply int
	log.Printf("Calling Exercitiul3 with data: %v\n", []int{12, 13, 14})
	err = rpc_conn.Call("Server.Exercitiul3", &Exercitiul3Args{Numbers: []int{12, 13, 14}}, &ex3_reply)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("Server responded with: %d\n", ex3_reply)
	}

	var ex4_reply float64
	log.Printf("Calling Exercitiul4 with data: %v\n", []int{2, 10, 11, 39, 32, 80, 84})
	err = rpc_conn.Call("Server.Exercitiul4", &Exercitiul4Args{Numbers: []int{2, 10, 11, 39, 32, 80, 84}}, &ex4_reply)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("Server responded with: %f\n", ex4_reply)
	}

	var ex5_reply []int
	log.Printf("Calling Exercitiul5 with data: %v\n", []string{"2dasdas", "12", "dasdas", "1010", "101"})
	err = rpc_conn.Call("Server.Exercitiul5", &Exercitiul5Args{Strings: []string{"2dasdas", "12", "dasdas", "1010", "101"}}, &ex5_reply)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("Server responded with: %v\n", ex5_reply)
	}

	var ex6_reply []string
	log.Printf("Calling Exercitiul6 with data: %v, %d, %s\n", []string{"abcdef", "uvwxyz"}, 3, "LEFT")
	err = rpc_conn.Call("Server.Exercitiul6", &Exercitiul6Args{Strings: []string{"abcdef", "uvwxyz"}, Shift: 3, Direction: "LEFT"}, &ex6_reply)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("Server responded with: %s\n", ex6_reply)
	}

	var ex7_reply string
	log.Printf("Calling Exercitiul7 with data: %s\n", "1G11o1L")
	err = rpc_conn.Call("Server.Exercitiul7", &Exercitiul7Args{EncodedString: "1G11o1L"}, &ex7_reply)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("Server responded with: %s\n", ex7_reply)
	}

	var ex8_reply int
	log.Printf("Calling Exercitiul8 with data: %v\n", []int{23, 17, 15, 3, 18})
	err = rpc_conn.Call("Server.Exercitiul8", &Exercitiul8Args{Numbers: []int{23, 17, 15, 3, 18}}, &ex8_reply)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("Server responded with: %v\n", ex8_reply)
	}

	var ex9_reply int
	log.Printf("Calling Exercitiul9 with data: %v\n", []string{"mama", "iris", "bunica", "ala"})
	err = rpc_conn.Call("Server.Exercitiul9", &Exercitiul9Args{Strings: []string{"mama", "iris", "bunica", "ala"}}, &ex9_reply)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("Server responded with: %v\n", ex9_reply)
	}

	var ex10_reply int
	log.Printf("Calling Exercitiul10 with data: %v\n", []string{"24", "16", "8", "aaa", "bbb"})
	err = rpc_conn.Call("Server.Exercitiul10", &Exercitiul10Args{Strings: []string{"24", "16", "8", "aaa", "bbb"}}, &ex10_reply)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("Server responded with: %v\n", ex10_reply)
	}

	var ex11_reply int
	log.Printf("Calling Exercitiul11 with data: %v, %d\n", []int{1234, 3456, 4567}, 2)
	err = rpc_conn.Call("Server.Exercitiul11", &Exercitiul11Args{Numbers: []int{1234, 3456, 4567}, K: 2}, &ex11_reply)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("Server responded with: %v\n", ex11_reply)
	}

	var ex12_reply int
	log.Printf("Calling Exercitiul12 with data: %v\n", []int{23, 43, 26, 74})
	err = rpc_conn.Call("Server.Exercitiul12", &Exercitiul12Args{Numbers: []int{23, 43, 26, 74}}, &ex12_reply)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("Server responded with: %v\n", ex12_reply)
	}

	var ex13_reply []float64
	log.Printf("Calling Exercitiul13 with data: %v\n", []int{3, 10, 3, 4, 5, 6})
	err = rpc_conn.Call("Server.Exercitiul13", &Exercitiul13Args{Numbers: []int{3, 10, 3, 4, 5, 6, 15, -19}}, &ex13_reply)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("Server responded with: %v\n", ex13_reply)
	}

	var ex14_reply []string
	log.Printf("Calling Exercitiul14 with data: %v\n", []string{"Ceva12!@", "asa212", "dasdas"})
	err = rpc_conn.Call("Server.Exercitiul14", &Exercitiul14Args{Strings: []string{"Ceva12!@", "asa212", "dasdas"}}, &ex14_reply)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("Server responded with: %v\n", ex14_reply)
	}

	var ex15_reply []string
	log.Printf("Calling Exercitiul15 with data: %v\n", []string{"a", "b", "e", "3", "!", "A"})
	err = rpc_conn.Call("Server.Exercitiul15", &Exercitiul15Args{Characters: []string{"a", "b", "e", "3", "!", "A"}}, &ex15_reply)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("Server responded with: %v\n", ex15_reply)
	}

	msg_conn.Close()
}
