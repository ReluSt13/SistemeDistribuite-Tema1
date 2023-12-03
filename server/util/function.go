package util

import (
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

func ExtractNumbers(s string) []int {
	re := regexp.MustCompile("[0-9]+")
	matches := re.FindAllString(s, -1)
	numbers := make([]int, len(matches))
	for i, match := range matches {
		numbers[i], _ = strconv.Atoi(match)
	}
	return numbers
}

func IsPerfectSquare(n int) bool {
	root := int(math.Sqrt(float64(n)))
	return root*root == n
}

func ReverseNumber(n int) int {
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n = n / 10
	}
	return reversed
}

func SumDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}

func IsBinary(s string) bool {
	for _, r := range s {
		if r != '0' && r != '1' {
			return false
		}
	}
	return true
}

func BinaryToBase10(s string) int {
	number, _ := strconv.ParseInt(s, 2, 64)
	return int(number)
}

func CaesarCipher(s string, shift int, direction string) string {
	shifted := ""
	for _, r := range s {
		if 'a' <= r && r <= 'z' {
			base := 'a'
			if direction == "RIGHT" {
				shifted += string((r-base+int32(shift))%26 + base)
			} else {
				shifted += string((r-base-int32(shift)+26)%26 + base)
			}
		} else if 'A' <= r && r <= 'Z' {
			base := 'A'
			if direction == "RIGHT" {
				shifted += string((r-base+int32(shift))%26 + base)
			} else {
				shifted += string((r-base-int32(shift)+26)%26 + base)
			}
		} else {
			shifted += string(r)
		}
	}
	return shifted
}

func Decode(encoded string) string {
	decoded := ""
	i := 0
	for i < len(encoded) {
		j := i
		for j < len(encoded) && '0' <= encoded[j] && encoded[j] <= '9' {
			j++
		}
		count, _ := strconv.Atoi(encoded[i:j])
		decoded += strings.Repeat(string(encoded[j]), count)
		i = j + 1
	}
	return decoded
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func CountDigits(n int) int {
	count := 0
	for n > 0 {
		count++
		n = n / 10
	}
	return count
}

func HasEvenVowelsAtEvenPositions(s string) bool {
	vowels := 0
	for i, r := range s {
		if i%2 == 0 {
			if isVowel(r) {
				vowels++
			} else {
				return false
			}
		}
	}
	return vowels%2 == 0
}

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func RotateRight(n, k int) int {
	digits := strconv.Itoa(n)
	k = k % len(digits)
	rotated := digits[len(digits)-k:] + digits[:len(digits)-k]
	result, _ := strconv.Atoi(rotated)
	return result
}

func DoubleFirstDigit(n int) int {
	digits := strconv.Itoa(n)
	firstDigit := string(digits[0])
	modified, _ := strconv.Atoi(firstDigit + digits)
	return modified
}

func IsValidPassword(s string) bool {
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(s)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(s)
	hasDigit := regexp.MustCompile(`\d`).MatchString(s)
	hasSymbol := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(s)

	return hasLower && hasUpper && hasDigit && hasSymbol
}

func GeneratePassword(characters []string, length int) string {
	var password strings.Builder

	for i := 0; i < length; i++ {
		// Append a random character from the array to the password
		index := rand.Intn(len(characters))
		password.WriteString(characters[index])
	}

	return password.String()
}
