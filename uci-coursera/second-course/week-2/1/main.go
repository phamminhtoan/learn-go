package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type UserPrompt int

const (
	RegexFloat          string = `([\d]+\.?\d*)`
	MaxUserPromptLength int    = 50
	QuitCode            string = "X"
)

const (
	ReadAcceleration UserPrompt = iota
	ReadVelocity
	ReadDisplacement
	ReadTime
	ShowStringValue
	RegexFound
	RegexNotFound
	MaxIntegerCount
)

func GetUserPromptPadded(up UserPrompt, str string, strFmt string) string {
	var paddedStr = str + strings.Repeat(" ", (MaxUserPromptLength-len(str))) + strFmt
	return paddedStr
}

func GetUserPrompt(up UserPrompt) string {
	var strReturnVal string
	switch up {
	case ReadAcceleration:
		strReturnVal = GetUserPromptPadded(up, "Please enter acceleration ('X' to Exit):", "")
	case ReadVelocity:
		strReturnVal = GetUserPromptPadded(up, "Please enter initial velocity ('X' to Exit):", "")
	case ReadDisplacement:
		strReturnVal = GetUserPromptPadded(up, "Please enter initial displacement ('X' to Exit):", "")
	case ReadTime:
		strReturnVal = GetUserPromptPadded(up, "Please enter time ('X' to Exit):", "")
	case RegexFound:
		strReturnVal = GetUserPromptPadded(up, "Valid Input!", "")
	case RegexNotFound:
		strReturnVal = GetUserPromptPadded(up, "Invalid Input!", "")
	}
	return strReturnVal
}

func readUserInput(prompt UserPrompt, regex string) (float64, error) {
	var inputVal float64
	var err error
	re := regexp.MustCompile(regex)
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Printf(GetUserPrompt(prompt))
	for scanner.Scan() {
		var strVal string = scanner.Text()
		if strVal == QuitCode {
			os.Exit(0)
		}
		match := re.Match([]byte(strVal))
		inputVal, err = strconv.ParseFloat(strVal, 64)
		if match == true && err == nil {
			return inputVal, err
		}
		fmt.Println(GetUserPrompt(RegexNotFound))
		fmt.Printf(GetUserPrompt(prompt))
	}
	return inputVal, err
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (a*(math.Pow(t, 2)*0.5) + (v0 * t) + (s0))
	}
	return fn
}

func main() {
	prompts := []UserPrompt{ReadAcceleration, ReadVelocity, ReadDisplacement, ReadTime}
	var acceleration, initialVelocity, initialDisplacement, time float64
	var err error

	for _, prompt := range prompts {
		switch prompt {
		case ReadAcceleration:
			acceleration, err = readUserInput(prompt, RegexFloat)
		case ReadVelocity:
			initialVelocity, err = readUserInput(prompt, RegexFloat)
		case ReadDisplacement:
			initialDisplacement, err = readUserInput(prompt, RegexFloat)
		case ReadTime:
			time, err = readUserInput(prompt, RegexFloat)
		}
	}

	if err == nil {
		fmt.Printf("acceleration: %f\n", acceleration)
		fmt.Printf("velocity:     %f\n", initialVelocity)
		fmt.Printf("displacement: %f\n", initialDisplacement)
		fmt.Printf("time:         %f\n", time)
	}
	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	fmt.Printf("Displacement: %f\n", fn(time))

}
