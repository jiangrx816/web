package math_service

import (
	"github.com/jiangrx816/gopkg/server/api"
	"math/rand"
	"time"
	"web/utils"
)

type MathComputeResult struct {
	NumberOne int    `json:"number_one"`
	NumberTwo int    `json:"number_two"`
	Symbol    string `json:"symbol"`
	Result    int    `json:"result"`
}

// MakeComputeList
/**
level 对应的类型
limit 每次生成的数量
1 5以内的加法
2 10以内的减法
3 10以内的加减
4 20以内加法（不进位）
5 20以内减法（不退位）
6 20以内加法（进位）
7 20以内减法（退位）
8 20以内加减
9 100以内加法
10 100以内减法
11 100以内加减
12 1到5乘法
13 6到9乘法
14 10以内乘法
15 2到5除法
16 6到9除法
17 10以内除法
*/
func (ms *MathService) MakeComputeList(level, limit int) (response []MathComputeResult, apiErr api.Error) {
	switch level {
	case 1:
		response = ms.additionExerciseLessThanFive(5, limit)
	case 2:
		response = ms.subtractionUpTen(10, limit)
	case 3:
		response = ms.plusOrMinusTen(10, limit)
	case 4:
		response = ms.additionExerciseLessThanFive(5, limit)
	case 5:
		response = ms.additionExerciseLessThanFive(5, limit)
	case 6:
		response = ms.additionExerciseLessThanFive(5, limit)
	case 7:
		response = ms.additionExerciseLessThanFive(5, limit)
	case 8:
		response = ms.additionExerciseLessThanFive(5, limit)
	case 9:
		response = ms.additionExerciseLessThanFive(5, limit)
	case 10:
		response = ms.additionExerciseLessThanFive(5, limit)
	case 11:
		response = ms.additionExerciseLessThanFive(5, limit)
	case 12:
		response = ms.additionExerciseLessThanFive(5, limit)
	case 13:
		response = ms.additionExerciseLessThanFive(5, limit)
	case 14:
		response = ms.additionExerciseLessThanFive(5, limit)
	case 15:
		response = ms.additionExerciseLessThanFive(5, limit)
	case 16:
		response = ms.additionExerciseLessThanFive(5, limit)
	case 17:
		response = ms.additionExerciseLessThanFive(5, limit)
	}
	return
}

// additionExerciseLessThanFive 5以内的加法
func (ms *MathService) additionExerciseLessThanFive(max, limit int) (computeList []MathComputeResult) {
	var computeData MathComputeResult
	for i := 0; i < limit; i++ {
		a, b := utils.GenerateRandomTwoNumber(max)
		computeData.NumberOne = a
		computeData.NumberTwo = b
		computeData.Symbol = "+"
		computeData.Result = a + b
		computeList = append(computeList, computeData)
	}
	return
}

// subtractionUpTen 10以内的减法
func (ms *MathService) subtractionUpTen(max, limit int) (computeList []MathComputeResult) {
	var computeData MathComputeResult
	for i := 0; i < limit; i++ {
		a, b := utils.GenerateRandomTwoNumber(max)

		if a > b {
			a, b = b, a
		}

		computeData.NumberOne = a
		computeData.NumberTwo = b
		computeData.Symbol = "-"
		computeData.Result = b - a
		computeList = append(computeList, computeData)
	}
	return
}

// plusOrMinusTen 10以内的加减
func (ms *MathService) plusOrMinusTen(max, limit int) (computeList []MathComputeResult) {
	var computeData MathComputeResult
	for i := 0; i < limit; i++ {
		a, b := utils.GenerateRandomTwoNumber(max)

		rand.Seed(time.Now().UnixNano()) // 设置随机种子
		op := rand.Intn(2)               // 随机选择 0 或 1, 用于决定加法或减法
		if op == 1 && a > b {
			a, b = b, a
			computeData.NumberOne = b
			computeData.NumberTwo = a
			computeData.Symbol = "-"
			computeData.Result = b - a
		} else {
			computeData.NumberOne = a
			computeData.NumberTwo = b
			computeData.Symbol = "+"
			computeData.Result = a + b
		}

		computeList = append(computeList, computeData)
	}
	return
}
