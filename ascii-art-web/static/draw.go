package static

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Draw(str, fs string) (string, int) {

	var res string
	lines, err := scanLines("./static/fonts/" + fs + ".txt")
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

		tmp, err := convToAscii(word, lines)
		if err != 1 {
			t := time.Now()
			fmt.Println(t.Format("3:4:5pm"), "Error: Invalid symbols")
			return tmp, 400
		}
		res += tmp
		res += "\n"
	}
	t := time.Now()
	fmt.Println(t.Format("3:4:5pm"), "Operation complete.")
	return res, 1
}

func convToAscii(str string, lines []string) (string, int) {

	var res string
	for _, let := range str {
		if let < 32 || let > 126 {
			if let != 13 && let != 10 {
				return "400 Bad request", 400
			}
		}
	}

	for i := 1; i <= 8; i++ {
		var output string

		for _, letter := range str {

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
	return res, 1
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
