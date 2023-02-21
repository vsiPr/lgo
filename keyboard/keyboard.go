//package keyboard gets user input
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//function GetFloat gets float numbers from user input
//it returns float-pointing number and error
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
