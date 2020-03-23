package types

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"log"
)

func lineToOWA(line []string) *Question{
	score, _ := strconv.ParseInt(line[2], 10, 64)
	question := &Question{Statement: line[0], Answer: line[1], Score: int(score)}
	return question
}

func parseCSV(csvFilePath string, questionBank *QuestionBank) error {

	csvFile, err := os.Open(csvFilePath)

	if err != nil {
		return err
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		question := lineToOWA(line)
		_ = questionBank.AddQuestion(question)
	}

	return nil
}

func clearScreen() {
	os := runtime.GOOS
	// if linux or mac`
	if (os == "linux" || os == "darwin"){
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if (os == "windows") {
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		log.Fatal("Unsupported Operating System! Can't clear your screen")
	}
	return nil
}

