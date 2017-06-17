package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func stringToJSON(input string) []byte {
	output, _ := json.Marshal(input)
	return output
}

func JSONToString(jsonString []byte) string {
	var output interface{}
	if err := json.Unmarshal(jsonString, &output); err != nil {
		panic(err)
	}
	return output.(string)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type a string to encode in JSON")
	text, _ := reader.ReadString('\n')
	jsonString := stringToJSON(text)
	fmt.Printf("String successfully encoded: %v\n", jsonString)
	decodedString := JSONToString(jsonString)
	fmt.Printf("Here's the decoded version back: %s", decodedString)
}
