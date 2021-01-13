package actions

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func div(task string) string {
	reDiv := regexp.MustCompile(`[0-9.]+\/[0-9.]+`)
	matchDiv := reDiv.FindAllString(task, 1)

	if len(matchDiv) > 0 {
		taskDiv := strings.Join(matchDiv, "")
		taskDivSplit := strings.Split(taskDiv, "/")
		num1Float, _ := strconv.ParseFloat(taskDivSplit[0], 64)
		num2Float, _ := strconv.ParseFloat(taskDivSplit[1], 64)
		if num2Float == 0 {
			print("Division by Zero isn't allowed!")
			os.Exit(3)
		} else {
			res := num1Float / num2Float
			resString := strconv.FormatFloat(res, 'f', 2, 64)

			// insert div result back into the task
			task = strings.Replace(task, taskDiv, resString, -1)
			fmt.Println(task)
		}
	}
	return task
}

func mul(task string) string {

	reMul := regexp.MustCompile(`[0-9.]+\*[0-9.]+`)
	matchMul := reMul.FindAllString(task, 1)

	if len(matchMul) > 0 {
		taskMul := strings.Join(matchMul, "")
		taskMulSplit := strings.Split(taskMul, "*")
		num1Float, _ := strconv.ParseFloat(taskMulSplit[0], 64)
		num2Float, _ := strconv.ParseFloat(taskMulSplit[1], 64)
		res := num1Float * num2Float
		resString := strconv.FormatFloat(res, 'f', 2, 64)

		// insert div result back into the task
		task = strings.Replace(task, taskMul, resString, -1)
		fmt.Println(task)
	}

	return task
}

func add(task string) (string, int) {

	reAdd := regexp.MustCompile(`[0-9.]+\+[0-9.]+`)
	matchAdd := reAdd.FindAllString(task, 1)
	taskCountAdd := 0

	if len(matchAdd) > 0 {
		taskAdd := strings.Join(matchAdd, "")
		taskAddSplit := strings.Split(taskAdd, "+")
		num1Float, _ := strconv.ParseFloat(taskAddSplit[0], 64)
		num2Float, _ := strconv.ParseFloat(taskAddSplit[1], 64)
		res := num1Float + num2Float
		resString := strconv.FormatFloat(res, 'f', 2, 64)

		// insert div result back into the task
		task = strings.Replace(task, taskAdd, resString, -1)
		fmt.Println(task)

		taskCountAdd = strings.Count(task, "+")
	}
	return task, taskCountAdd
}

func sub(task string) (string, int) {

	reSub := regexp.MustCompile(`[0-9.]+\-[0-9.]+`)
	matchSub := reSub.FindAllString(task, 1)
	taskCountSub := 0

	if len(matchSub) > 0 {
		taskSub := strings.Join(matchSub, "")
		taskAddSplit := strings.Split(taskSub, "-")
		num1Float, _ := strconv.ParseFloat(taskAddSplit[0], 64)
		num2Float, _ := strconv.ParseFloat(taskAddSplit[1], 64)
		res := num1Float - num2Float
		resString := strconv.FormatFloat(res, 'f', 2, 64)

		// insert result back into the task
		task = strings.Replace(task, taskSub, resString, -1)
		fmt.Println(task)

		taskCountSub = strings.Count(task, "-")
	}
	return task, taskCountSub
}

func findIndex(task string) ([]int, []int, []int, []int) {
	indexMul := findTheSymbol(task, "*")
	indexDiv := findTheSymbol(task, "/")
	indexPlu := findTheSymbol(task, "+")
	indexMin := findTheSymbol(task, "-")
	return indexMul, indexDiv, indexPlu, indexMin
}

func findTheSymbol(task string, findThing string) []int {
	var foundIndexes []int
	for i := 0; i < len(task); i++ {
		if findThing == string(task[i]) {
			foundIndexes = append(foundIndexes, i)
		}
	}
	return foundIndexes
}

func calc(task string) string {
	taskCountAdd := strings.Count(task, "+")
	taskCountSub := strings.Count(task, "-")

	indexMul, indexDiv, _, _ := findIndex(task)

	// search if / or * first
	i := 0
	for len(indexMul) != 0 || len(indexDiv) != 0 {
		if (len(indexMul) == len(indexDiv)) && (len(indexMul) != 0 && len(indexDiv) != 0) {
			indexMul, indexDiv, _, _ = findIndex(task)
			if indexDiv[i] > indexMul[i] {
				// Mul
				task = mul(task)
				indexMul, indexDiv, _, _ = findIndex(task)
			} else if indexDiv[i] < indexMul[i] {
				// Div
				task = div(task)
				indexMul, indexDiv, _, _ = findIndex(task)
			}
		} else if len(indexMul) > 0 {
			// Mul
			task = mul(task)
			indexMul, indexDiv, _, _ = findIndex(task)
		} else if len(indexDiv) > 0 {
			// Div
			task = div(task)
			indexMul, indexDiv, _, _ = findIndex(task)
		}
	}

	for taskCountAdd != 0 {
		// Addition
		task, taskCountAdd = add(task)
	}

	for taskCountSub != 0 {
		// Subtraction
		task, taskCountSub = sub(task)
	}

	return task
}
