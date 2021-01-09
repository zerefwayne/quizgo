package types

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func lineToOWA(line []string) *Question {
	score, _ := strconv.ParseInt(line[2], 10, 64)
	question := &Question{Statement: line[0], Answer: line[1], Score: int(score)}
	return question
}

func parseStrings(csv [][]string, questionBank *QuestionBank) error {
	for _, line := range csv {
		question := lineToOWA(line)
		err := questionBank.AddQuestion(question)

		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}

func getCSVFileContents(filePath string) ([][]string, error) {
	csvContents, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(csvContents)
	csvs, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}
	return csvs, nil
}

func getJSONFileContents(filePath string) ([]byte, error) {
	// Open our jsonFile
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue, nil
}
