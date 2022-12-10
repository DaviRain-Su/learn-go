package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	number, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
