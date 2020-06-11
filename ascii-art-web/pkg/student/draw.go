package student

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Draw(str, fs string) (string, int) {

	var res string
	lines, err := scanLines("./pkg/fonts/" + fs + ".txt")
	if err != nil {
		t := time.Now()
		fmt.Println(t.Format("3:4:5pm"), "Error:", err)
		return "500 Internal server error", 500
	}

	arr := strings.Split(str, "\n")

	if str == "" {
		return "", 1
	}

	for _, word := range arr {

		res += convToAscii(word, lines)
		res += "\n"
	}
	return res, 1
}

func convToAscii(str string, lines []string) string {

	var res string
	var clearStr string
	for _, let := range str {
		if (let >= 32 && let <= 126) || (let == 13 || let == 10) {
			clearStr += string(let)
		}
	}

	if len(clearStr) > 1000 {
		clearStr = "Nice try, but no"
	}

	for i := 1; i <= 8; i++ {
		var output string

		for _, letter := range clearStr {

			if letter != ' ' {
				readFrom := (int(letter-32) * 9) + i
				for index, line := range lines {
					if index == readFrom {
						output += line
						break
					}
				}
			} else {
				output += "      "
			}
		}
		res += output
		res += "\n"
	}
	return res
}

func scanLines(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	defer file.Close()
	return lines, nil
}
