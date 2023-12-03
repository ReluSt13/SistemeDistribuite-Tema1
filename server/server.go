package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/cmplx"
	"math/rand"
	"net"
	"net/rpc"
	"server/util"
	"sort"
	"strings"
)

type Server struct{}

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

func (s *Server) Exercitiul1(args *Exercitiul1Args, reply *[]string) error {
	log.Printf("Client has called Exercitiul1 with this data: %v\n", args.Words)
	*reply = processExercitiul1(args.Words)
	log.Printf("Server sends: %s back to the client.\n", *reply)
	return nil
}

func (s *Server) Exercitiul2(args *Exercitiul2Args, reply *int) error {
	log.Printf("Client has called Exercitiul2 with this data: %v\n", args.Words)
	*reply = processExercitiul2(args.Words)
	log.Printf("Server sends: %d perfect squares found.\n", *reply)
	return nil
}

func (s *Server) Exercitiul3(args *Exercitiul3Args, reply *int) error {
	log.Printf("Client has called Exercitiul3 with this data: %v\n", args.Numbers)
	*reply = processExercitiul3(args.Numbers)
	log.Printf("Server sends: %d back to the client.\n", *reply)
	return nil
}

func (s *Server) Exercitiul4(args *Exercitiul4Args, reply *float64) error {
	log.Printf("Client has called Exercitiul4 with this data: %v\n", args.Numbers)
	*reply = processExercitiul4(args.Numbers)
	log.Printf("Server sends: %f back to the client.\n", *reply)
	return nil
}

func (s *Server) Exercitiul5(args *Exercitiul5Args, reply *[]int) error {
	log.Printf("Client has called Exercitiul5 with this data: %v\n", args.Strings)
	*reply = processExercitiul5(args.Strings)
	log.Printf("Server sends: %v back to the client.\n", *reply)
	return nil
}

func (s *Server) Exercitiul6(args *Exercitiul6Args, reply *[]string) error {
	log.Printf("Client has called Exercitiul6 with this data: %v\n", args.Strings)
	*reply = processExercitiul6(args.Strings, args.Shift, args.Direction)
	log.Printf("Server sends: %v back to the client.\n", *reply)
	return nil
}

func (s *Server) Exercitiul7(args *Exercitiul7Args, reply *string) error {
	log.Printf("Client has called Exercitiul7 with this data: %s\n", args.EncodedString)
	*reply = processExercitiul7(args.EncodedString)
	log.Printf("Server sends: %s back to the client.\n", *reply)
	return nil
}

func (s *Server) Exercitiul8(args *Exercitiul8Args, reply *int) error {
	log.Printf("Client has called Exercitiul8 with this data: %v\n", args.Numbers)
	*reply = processExercitiul8(args.Numbers)
	log.Printf("Server sends: %d back to the client.\n", *reply)
	return nil
}

func (s *Server) Exercitiul9(args *Exercitiul9Args, reply *int) error {
	log.Printf("Client has called Exercitiul9 with this data: %v\n", args.Strings)
	*reply = processExercitiul9(args.Strings)
	log.Printf("Server sends: %d back to the client.\n", *reply)
	return nil
}

func (s *Server) Exercitiul10(args *Exercitiul10Args, reply *int) error {
	log.Printf("Client has called Exercitiul10 with this data: %v\n", args.Strings)
	*reply = processExercitiul10(args.Strings)
	log.Printf("Server sends: %d back to the client.\n", *reply)
	return nil
}

func (s *Server) Exercitiul11(args *Exercitiul11Args, reply *int) error {
	log.Printf("Client has called Exercitiul11 with this data: %v, %d\n", args.Numbers, args.K)
	*reply = processExercitiul11(args.Numbers, args.K)
	log.Printf("Server sends: %d back to the client.\n", *reply)
	return nil
}

func (s *Server) Exercitiul12(args *Exercitiul12Args, reply *int) error {
	log.Printf("Client has called Exercitiul12 with this data: %v\n", args.Numbers)
	*reply = processExercitiul12(args.Numbers)
	log.Printf("Server sends: %d back to the client.\n", *reply)
	return nil
}

func (s *Server) Exercitiul13(args *Exercitiul13Args, reply *[]float64) error {
	log.Printf("Client has called Exercitiul13 with this data: %v\n", args.Numbers)
	*reply = processExercitiul13(args.Numbers)
	log.Printf("Server sends: %v back to the client.\n", *reply)
	return nil
}

func (s *Server) Exercitiul14(args *Exercitiul14Args, reply *[]string) error {
	log.Printf("Client has called Exercitiul14 with this data: %v\n", args.Strings)
	*reply = processExercitiul14(args.Strings)
	log.Printf("Server sends: %v back to the client.\n", *reply)
	return nil
}

func (s *Server) Exercitiul15(args *Exercitiul15Args, reply *[]string) error {
	log.Printf("Client has called Exercitiul15 with this data: %v\n", args.Characters)
	*reply = processExercitiul15(args.Characters)
	log.Printf("Server sends: %v back to the client.\n", *reply)
	return nil
}

func processExercitiul1(words []string) []string {
	log.Println("Server is processing the data...")
	var output []string

	if len(words) > 0 {
		wordLength := len(words[0])

		for i := 0; i < wordLength; i++ {
			var wordBuilder strings.Builder
			for _, word := range words {
				wordBuilder.WriteByte(word[i])
			}
			output = append(output, wordBuilder.String())
		}
	}

	return output
}

func processExercitiul2(words []string) int {
	var count int
	for _, word := range words {
		numbers := util.ExtractNumbers(word)
		for _, num := range numbers {
			if util.IsPerfectSquare(num) {
				count++
			}
		}
	}
	return count
}

func processExercitiul3(numbers []int) int {
	log.Println("Server is processing the data...")
	sum := 0

	for _, number := range numbers {
		reversed := util.ReverseNumber(number)
		sum += reversed
	}

	return sum
}

func processExercitiul4(numbers []int) float64 {
	log.Println("Server is processing the data...")
	sum := 0
	count := 0
	a := numbers[0]
	b := numbers[1]

	for _, number := range numbers[2:] {
		sumDigits := util.SumDigits(number)
		if sumDigits >= a && sumDigits <= b {
			sum += number
			count++
		}
	}

	if count == 0 {
		return 0
	}

	return float64(sum) / float64(count)
}

func processExercitiul5(strings []string) []int {
	log.Println("Server is processing the data...")
	var numbers []int

	for _, str := range strings {
		if util.IsBinary(str) {
			number := util.BinaryToBase10(str)
			numbers = append(numbers, number)
		}
	}

	return numbers
}

func processExercitiul6(strings []string, shift int, direction string) []string {
	log.Println("Server is processing the data...")
	var encodedStrings []string

	for _, str := range strings {
		encoded := util.CaesarCipher(str, shift, direction)
		encodedStrings = append(encodedStrings, encoded)
	}

	return encodedStrings
}

func processExercitiul7(encoded string) string {
	log.Println("Server is processing the data...")
	return util.Decode(encoded)
}

func processExercitiul8(numbers []int) int {
	log.Println("Server is processing the data...")
	count := 0

	for _, number := range numbers {
		if util.IsPrime(number) {
			count += util.CountDigits(number)
		}
	}

	return count
}

func processExercitiul9(strings []string) int {
	log.Println("Server is processing the data...")
	count := 0

	for _, str := range strings {
		if util.HasEvenVowelsAtEvenPositions(str) {
			count++
		}
	}

	return count
}

func processExercitiul10(strings []string) int {
	log.Println("Server is processing the data...")
	var numbers []int

	for _, str := range strings {
		numbers = append(numbers, util.ExtractNumbers(str)...)
	}

	if len(numbers) == 0 {
		return 0
	}

	gcd := numbers[0]
	for _, number := range numbers[1:] {
		gcd = util.GCD(gcd, number)
	}

	return gcd
}

func processExercitiul11(numbers []int, k int) int {
	log.Println("Server is processing the data...")
	sum := 0

	for _, number := range numbers {
		rotated := util.RotateRight(number, k)
		sum += rotated
	}

	return sum
}

func processExercitiul12(numbers []int) int {
	log.Println("Server is processing the data...")
	sum := 0

	for _, number := range numbers {
		modified := util.DoubleFirstDigit(number)
		sum += modified
	}

	return sum
}

func processExercitiul13(numbers []int) []float64 {
	log.Println("Server is processing the data...")
	a, b := numbers[0], numbers[1]
	var absValues []float64

	for i := 2; i < len(numbers); i += 2 {
		complexNumber := complex(float64(numbers[i]), float64(numbers[i+1]))
		absValue := cmplx.Abs(complexNumber)
		if absValue < float64(a) || absValue > float64(b) {
			absValues = append(absValues, absValue)
		}
	}

	sort.Float64s(absValues)
	return absValues
}

func processExercitiul14(strings []string) []string {
	log.Println("Server is processing the data...")
	var validPasswords []string

	for _, str := range strings {
		if util.IsValidPassword(str) {
			validPasswords = append(validPasswords, str)
		}
	}

	return validPasswords
}

func processExercitiul15(characters []string) []string {
	log.Println("Server is processing the data...")
	var passwords []string

	// Generate a random number of passwords between 1 and 10
	numPasswords := rand.Intn(10) + 1

	for i := 0; i < numPasswords; i++ {
		// Generate a password of random length between 5 and 15
		length := rand.Intn(11) + 5
		password := util.GeneratePassword(characters, length)
		passwords = append(passwords, password)
	}

	return passwords
}

func handleMsgConn(c net.Conn) {
	defer c.Close()

	reader := bufio.NewReader(c)
	clientName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("Client %s connected\n", strings.TrimSpace(clientName))

	_, err = c.Write([]byte("Server accepted connection of client " + strings.TrimSpace(clientName) + "\n"))
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Printf("Client %s disconnected\n", strings.TrimSpace(clientName))
			} else {
				fmt.Println(err)
			}
			break
		}
	}
}

func main() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg_ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("Server started...")
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)

		msgConn, err := msg_ln.Accept()
		if err != nil {
			continue
		}
		go handleMsgConn(msgConn)
	}
}
