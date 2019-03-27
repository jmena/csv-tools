package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
)

func main() {

	input := csv.NewReader(bufio.NewReader(os.Stdin))

	header, err := input.Read()

	if err == io.EOF {
		return
	} else if err != nil {
		log.Fatal(err)
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	out.WriteString("[")
	printComma := false
	for {
		line, err := input.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(record)
		jsonObj := make(map[string]string)
		for i := 0; i < len(header); i++ {
			jsonObj[header[i]] = line[i]
		}

		jsonStr, err := json.Marshal(jsonObj)
		if err != nil {
			log.Fatal(err)
		}
		if printComma {
			out.WriteString(",")
		} else {

			printComma = true
		}
		out.Write(jsonStr)

	}
	out.WriteString("]")
}
