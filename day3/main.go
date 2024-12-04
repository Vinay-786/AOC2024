package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Unable to read file")
	}
	defer file.Close()

	var tempbuf bytes.Buffer
	io.Copy(&tempbuf, file)

	buffer := tempbuf.String()
	fmt.Println(buffer)

	var answer int = 0
	var enable bool = true
	for i := 0; i < len(buffer); i++ {
		if i <= len(buffer)-4 && buffer[i:i+4] == "do()" {
			enable = true
		}
		if i <= len(buffer)-7 && buffer[i:i+7] == "don't()" {
			enable = false
		}
		if enable && buffer[i] == 'm' {
			// check for mul
			if buffer[i+1] == 'u' && buffer[i+2] == 'l' && buffer[i+3] == '(' {

				i += 4 // skip 4 char mul(
				x, n, _ := parseNextNumber(buffer, i)
				fmt.Println("value of x: ", x)
				i += n

				if i < len(buffer) && buffer[i] == ',' {
					i++

					y, n, _ := parseNextNumber(buffer, i)
					fmt.Println("value of y: ", y)
					i += n

					if i < len(buffer) && buffer[i] == ')' {
						fmt.Printf("X * Y: %d, %d\n", x, y)
						if x != -1 && y != -1 {
							answer += x * y
						}
					}
				}
			}
		}
	}

	fmt.Println("Answer: ", answer)
}

func parseNextNumber(buffer string, start int) (int, int, error) {
	var digits strings.Builder
	i := start

	for i < len(buffer) {
		if unicode.IsDigit(rune(buffer[i])) {
			digits.WriteByte(buffer[i])
			i++
		} else {
			break
		}
	}

	if digits.Len() == 0 {
		return -1, 0, fmt.Errorf("no digits found")
	}

	num, err := strconv.Atoi(digits.String())
	if err != nil {
		return -1, 0, err
	}

	return num, i - start, nil
}
