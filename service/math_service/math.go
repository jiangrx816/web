package math_service

import (
	"github.com/jiangrx816/gopkg/server/api"
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
