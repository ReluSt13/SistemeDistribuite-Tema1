package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/rpc"
	"os"
	"strconv"
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

func parseToIntSlice(s string) []int {
	strs := strings.Split(s, ",")
	ints := make([]int, len(strs))
	for i, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalf("Invalid number: %s", str)
		}
		ints[i] = num
	}
	return ints
}

func parseToWordsShiftDirection(params string) ([]string, int, string) {
	paramsSlice := strings.Split(params, ",")
	if len(paramsSlice) < 4 {
		log.Fatalf("Not enough parameters for Exercitiul6")
	}
	words := []string{paramsSlice[0], paramsSlice[1]}
	shift, err := strconv.Atoi(paramsSlice[2])
	if err != nil {
		log.Fatalf("Invalid shift parameter: %s", paramsSlice[2])
	}
	direction := paramsSlice[3]
	return words, shift, direction
}

func parseToNumbersAndDigit(params string) ([]int, int) {
	paramsSlice := strings.Split(params, ",")
	if len(paramsSlice) < 2 {
		log.Fatalf("Not enough parameters for Exercitiul11")
	}
	numbers := make([]int, len(paramsSlice)-1)
	for i := 0; i < len(paramsSlice)-1; i++ {
		num, err := strconv.Atoi(paramsSlice[i])
		if err != nil {
			log.Fatalf("Invalid number parameter: %s", paramsSlice[i])
		}
		numbers[i] = num
	}
	digit, err := strconv.Atoi(paramsSlice[len(paramsSlice)-1])
	if err != nil {
		log.Fatalf("Invalid digit parameter: %s", paramsSlice[len(paramsSlice)-1])
	}
	return numbers, digit
}

type Config struct {
	MaxArraySize  int `json:"maxArraySize"`
	MaxStringSize int `json:"maxStringSize"`
}

func readConfig() Config {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func main() {
	// Define command-line flags
	exerciseName := flag.String("exercise", "", "The name of the exercise to call")
	clientName := flag.String("client", "", "The name of the client")
	params := flag.String("params", "", "The parameters for the exercise, comma-separated")
	flag.Parse()

	// read config
	config := readConfig()

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
	_, err = msg_conn.Write([]byte(*clientName + "\n"))
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
					return
				}
				break
			}
			// print the message trimmed
			log.Println(strings.TrimSpace(msg))
		}
	}()

	switch *exerciseName {
	case "Exercitiul1":
		var reply []string
		var words []string
		if *params == "" {
			words = []string{"casa", "masa", "trei", "tanc", "4321"}
		} else {
			words = strings.Split(*params, ",")
		}
		if len(words) > config.MaxArraySize {
			log.Fatalf("Array size is larger than the maximum allowed size of %d", config.MaxArraySize)
		}
		args := &Exercitiul1Args{Words: words}
		err = rpc_conn.Call("Server.Exercitiul1", args, &reply)
		if err != nil {
			fmt.Println(err)
		} else {
			log.Printf("Server responded to client %s with: %s\n", *clientName, reply)
		}
	case "Exercitiul2":
		var reply int
		var words []string
		if *params == "" {
			words = []string{"abd4g5", "1sdf6fd", "fd2fdsf5"}
		} else {
			words = strings.Split(*params, ",")
		}
		if len(words) > config.MaxArraySize {
			log.Fatalf("Array size is larger than the maximum allowed size of %d", config.MaxArraySize)
		}
		args := &Exercitiul2Args{Words: words}
		err = rpc_conn.Call("Server.Exercitiul2", args, &reply)
		if err != nil {
			fmt.Println(err)
		} else {
			log.Printf("Server responded to client %s with: %d\n", *clientName, reply)
		}
	case "Exercitiul3":
		var reply int
		var numbers []int
		if *params == "" {
			numbers = []int{12, 13, 14}
		} else {
			numbers = parseToIntSlice(*params)
		}
		if len(numbers) > config.MaxArraySize {
			log.Fatalf("Array size is larger than the maximum allowed size of %d", config.MaxArraySize)
		}
		args := &Exercitiul3Args{Numbers: numbers}
		err = rpc_conn.Call("Server.Exercitiul3", args, &reply)
		if err != nil {
			fmt.Println(err)
		} else {
			log.Printf("Server responded to client %s with: %d\n", *clientName, reply)
		}
	case "Exercitiul4":
		var reply float64
		var numbers []int
		if *params == "" {
			numbers = []int{2, 10, 11, 39, 32, 80, 84}
		} else {
			numbers = parseToIntSlice(*params)
		}
		if len(numbers) > config.MaxArraySize {
			log.Fatalf("Array size is larger than the maximum allowed size of %d", config.MaxArraySize)
		}
		args := &Exercitiul4Args{Numbers: numbers}
		err = rpc_conn.Call("Server.Exercitiul4", args, &reply)
		if err != nil {
			fmt.Println(err)
		} else {
			log.Printf("Server responded to client %s with: %f\n", *clientName, reply)
		}
	case "Exercitiul5":
		var reply []int
		var words []string
		if *params == "" {
			words = []string{"2dasdas", "12", "dasdas", "1010", "101"}
		} else {
			words = strings.Split(*params, ",")
		}
		if len(words) > config.MaxArraySize {
			log.Fatalf("Array size is larger than the maximum allowed size of %d", config.MaxArraySize)
		}
		args := &Exercitiul5Args{Strings: words}
		err = rpc_conn.Call("Server.Exercitiul5", args, &reply)
		if err != nil {
			fmt.Println(err)
		} else {
			log.Printf("Server responded to client %s with: %v\n", *clientName, reply)
		}
	case "Exercitiul6":
		var reply []string
		var words []string
		var shift int
		var direction string
		if *params == "" {
			words = []string{"abcdef", "uvwxyz"}
			shift = 3
			direction = "LEFT"
		} else {
			words, shift, direction = parseToWordsShiftDirection(*params)
		}
		if len(words) > config.MaxArraySize {
			log.Fatalf("Array size is larger than the maximum allowed size of %d", config.MaxArraySize)
		}
		args := &Exercitiul6Args{Strings: words, Shift: shift, Direction: direction}
		err = rpc_conn.Call("Server.Exercitiul6", args, &reply)
		if err != nil {
			fmt.Println(err)
		} else {
			log.Printf("Server responded to client %s with: %s\n", *clientName, reply)
		}
	case "Exercitiul7":
		var reply string
		var encodedString string
		if *params == "" {
			encodedString = "1G11o1L"
		} else {
			encodedString = *params
		}
		if len(encodedString) > config.MaxStringSize {
			log.Fatalf("String size is larger than the maximum allowed size of %d", config.MaxStringSize)
		}
		args := &Exercitiul7Args{EncodedString: encodedString}
		err = rpc_conn.Call("Server.Exercitiul7", args, &reply)
		if err != nil {
			fmt.Println(err)
		} else {
			log.Printf("Server responded to client %s with: %s\n", *clientName, reply)
		}
	case "Exercitiul8":
		var reply int
		var numbers []int
		if *params == "" {
			numbers = []int{23, 17, 15, 3, 18}
		} else {
			numbers = parseToIntSlice(*params)
		}
		if len(numbers) > config.MaxArraySize {
			log.Fatalf("Array size is larger than the maximum allowed size of %d", config.MaxArraySize)
		}
		args := &Exercitiul8Args{Numbers: numbers}
		err = rpc_conn.Call("Server.Exercitiul8", args, &reply)
		if err != nil {
			fmt.Println(err)
		} else {
			log.Printf("Server responded to client %s with: %d\n", *clientName, reply)
		}
	case "Exercitiul9":
		var reply int
		var words []string
		if *params == "" {
			words = []string{"mama", "iris", "bunica", "ala"}
		} else {
			words = strings.Split(*params, ",")
		}
		if len(words) > config.MaxArraySize {
			log.Fatalf("Array size is larger than the maximum allowed size of %d", config.MaxArraySize)
		}
		args := &Exercitiul9Args{Strings: words}
		err = rpc_conn.Call("Server.Exercitiul9", args, &reply)
		if err != nil {
			fmt.Println(err)
		} else {
			log.Printf("Server responded to client %s with: %d\n", *clientName, reply)
		}
	case "Exercitiul10":
		var reply int
		var words []string
		if *params == "" {
			words = []string{"24", "16", "8", "aaa", "bbb"}
		} else {
			words = strings.Split(*params, ",")
		}
		if len(words) > config.MaxArraySize {
			log.Fatalf("Array size is larger than the maximum allowed size of %d", config.MaxArraySize)
		}
		args := &Exercitiul10Args{Strings: words}
		err = rpc_conn.Call("Server.Exercitiul10", args, &reply)
		if err != nil {
			fmt.Println(err)
		} else {
			log.Printf("Server responded to client %s with: %d\n", *clientName, reply)
		}
	case "Exercitiul11":
		var reply int
		var numbers []int
		var digit int
		if *params == "" {
			numbers = []int{1234, 3456, 4567}
			digit = 2
		} else {
			numbers, digit = parseToNumbersAndDigit(*params)
		}
		if len(numbers) > config.MaxArraySize {
			log.Fatalf("Array size is larger than the maximum allowed size of %d", config.MaxArraySize)
		}
		args := &Exercitiul11Args{Numbers: numbers, K: digit}
		err = rpc_conn.Call("Server.Exercitiul11", args, &reply)
		if err != nil {
			fmt.Println(err)
		} else {
			log.Printf("Server responded to client %s with: %d\n", *clientName, reply)
		}
	case "Exercitiul12":
		var reply int
		var numbers []int
		if *params == "" {
			numbers = []int{23, 43, 26, 74}
		} else {
			numbers = parseToIntSlice(*params)
		}
		if len(numbers) > config.MaxArraySize {
			log.Fatalf("Array size is larger than the maximum allowed size of %d", config.MaxArraySize)
		}
		args := &Exercitiul12Args{Numbers: numbers}
		err = rpc_conn.Call("Server.Exercitiul12", args, &reply)
		if err != nil {
			fmt.Println(err)
		} else {
			log.Printf("Server responded to client %s with: %d\n", *clientName, reply)
		}
	case "Exercitiul13":
		var reply []float64
		var numbers []int
		if *params == "" {
			numbers = []int{3, 10, 3, 4, 5, 6}
		} else {
			numbers = parseToIntSlice(*params)
		}
		if len(numbers) > config.MaxArraySize {
			log.Fatalf("Array size is larger than the maximum allowed size of %d", config.MaxArraySize)
		}
		args := &Exercitiul13Args{Numbers: numbers}
		err = rpc_conn.Call("Server.Exercitiul13", args, &reply)
		if err != nil {
			fmt.Println(err)
		} else {
			log.Printf("Server responded to client %s with: %v\n", *clientName, reply)
		}
	case "Exercitiul14":
		var reply []string
		var words []string
		if *params == "" {
			words = []string{"Ceva12!@", "asa212", "dasdas"}
		} else {
			words = strings.Split(*params, ",")
		}
		if len(words) > config.MaxArraySize {
			log.Fatalf("Array size is larger than the maximum allowed size of %d", config.MaxArraySize)
		}
		args := &Exercitiul14Args{Strings: words}
		err = rpc_conn.Call("Server.Exercitiul14", args, &reply)
		if err != nil {
			fmt.Println(err)
		} else {
			log.Printf("Server responded to client %s with: %v\n", *clientName, reply)
		}
	case "Exercitiul15":
		var reply []string
		var characters []string
		if *params == "" {
			characters = []string{"a", "b", "e", "3", "!", "A"}
		} else {
			characters = strings.Split(*params, ",")
		}
		if len(characters) > config.MaxArraySize {
			log.Fatalf("Array size is larger than the maximum allowed size of %d", config.MaxArraySize)
		}
		args := &Exercitiul15Args{Characters: characters}
		err = rpc_conn.Call("Server.Exercitiul15", args, &reply)
		if err != nil {
			fmt.Println(err)
		} else {
			log.Printf("Server responded with: %v\n", reply)
		}
	default:
		log.Fatalf("Unknown exercise: %s", *exerciseName)
	}

	msg_conn.Close()
}
