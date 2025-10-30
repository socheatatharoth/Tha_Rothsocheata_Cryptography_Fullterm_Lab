package main
import (
	"encoding/base64"
	"fmt"
	"encoding/hex"
)
func Stringtobin (input string) string{
	bin := ""
	for i := 0; i< len(input); i++{
		bin += fmt.Sprint("%08b", input[i])
	}
	return bin
}
func main(){
	var input string
	fmt.Print("Enter the text:")
	fmt.Scan(&input)

	bin := Stringtobin(input)
	fmt.Printf("The binary representation of '%s' is: %s\n", input, bin)

	hex := hex.EncodeToString([]byte(input))
	fmt.Printf("The hex representation of %s is %s\n", input, hex)

	base64 := base64.StdEncoding.EncodeToString([]byte(input))
	fmt.Printf("The base64 representation of %s is %s\n", input, base64)
}