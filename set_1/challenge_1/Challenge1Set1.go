package main

import (
	//	"bufio"
	"encoding/base64"
        "encoding/hex"
	"fmt"
	"log"
	//	"os"
)

func main() {
	var input string
	//reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text in Hex to be converted to base64: ")
	//text, err := reader.ReadString('\n')
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error occurred\n")
		log.Fatal(err)
	} else {
		fmt.Println(input)
                hexDecode, _ := hex.DecodeString(input)
		base64Input, _ := encodeBase64(hexDecode)
                fmt.Println("base64 == ", base64Input)
	}
}

func encodeBase64(input []byte) (string, error) {
	//byteString := []byte(input)
	str := base64.StdEncoding.EncodeToString(input)
	return str, nil
}
