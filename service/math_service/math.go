package math_service

import (
	"fmt"
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
		response = ms.exercise1(limit)
	case 2:
		response = ms.exercise2(limit)
	case 3:
		response = ms.exercise3(limit)
	case 8:
		response = ms.exercise8(limit)
	case 9:
		response = ms.exercise9(limit)
	case 10:
		response = ms.exercise10(limit)
	case 11:
		response = ms.exercise11(limit)
	case 12:
		response = ms.exercise12(limit)
	case 13:
		response = ms.exercise13(limit)
	case 14:
		response = ms.exercise14(limit)
	case 15:
		response = ms.exercise15(limit)
	case 16:
		response = ms.exercise16(limit)
	case 17:
		response = ms.exercise17(limit)
	}
	return
}

// additionExerciseLessThanFive 5以内的加法
func (ms *MathService) exercise1(limit int) (computeList []MathComputeResult) {
	var computeData MathComputeResult
	for i := 0; i < limit; i++ {
		a, b := utils.GenerateRandomTwoNumber(5)
		computeData.NumberOne = a
		computeData.NumberTwo = b
		computeData.Symbol = "+"
		computeData.Result = a + b
		computeList = append(computeList, computeData)
	}
	return
}

// subtractionUpTen 10以内的减法
func (ms *MathService) exercise2(limit int) (computeList []MathComputeResult) {
	var computeData MathComputeResult
	for i := 0; i < limit; i++ {
		a, b := utils.GenerateRandomTwoNumber(10)

		if a > b {
			a, b = b, a
		}

		computeData.NumberOne = b
		computeData.NumberTwo = a
		computeData.Symbol = "-"
		computeData.Result = b - a
		computeList = append(computeList, computeData)
	}
	return
}

func (ms *MathService) exercise3(limit int) (computeList []MathComputeResult) {
	var computeData MathComputeResult
	for i := 0; i < limit; i++ {
		a, b := utils.GenerateRandomTwoNumber(10)

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

func (ms *MathService) exercise8(limit int) (computeList []MathComputeResult) {
	var computeData MathComputeResult
	rand.Seed(time.Now().UnixNano())
	results := make(map[string]struct{}) // 用于存储唯一的结果
	count := 0

	for count < limit {
		// 随机选择操作类型（加法或减法）
		isAddition := rand.Intn(2) == 0

		var key string
		var result int

		if isAddition {
			// 生成加法表达式
			addend1 := rand.Intn(20) + 1
			addend2 := rand.Intn(20) + 1
			result = addend1 + addend2

			// 确保加法结果大于 10
			if result > 10 {
				key = fmt.Sprintf("%d + %d = %d", addend1, addend2, result)
				computeData.NumberOne = addend1
				computeData.NumberTwo = addend2
				computeData.Symbol = "＋"
				computeData.Result = result
			}
		} else {
			// 生成减法表达式
			minuend := rand.Intn(20) + 1
			subtrahend := rand.Intn(minuend) // 确保被减数大于减数
			result = minuend - subtrahend

			// 确保减法结果大于 0
			if result > 0 {
				key = fmt.Sprintf("%d - %d = %d", minuend, subtrahend, subtrahend)
				computeData.NumberOne = minuend
				computeData.NumberTwo = subtrahend
				computeData.Symbol = "－"
				computeData.Result = subtrahend
			}
		}

		// 检查并存储唯一结果
		if key != "" {
			if _, exists := results[key]; !exists {
				results[key] = struct{}{}
				computeList = append(computeList, computeData)
				count++
			}
		}
	}
	return
}

func (ms *MathService) exercise9(limit int) (computeList []MathComputeResult) {
	var computeData MathComputeResult
	rand.Seed(time.Now().UnixNano())
	results := make(map[string]struct{}) // 用于存储唯一的结果
	count := 0

	for count < limit {
		// 生成 1 到 100 之间的两个加数
		addend1 := rand.Intn(100) + 1
		addend2 := rand.Intn(100) + 1

		// 计算和
		sum := addend1 + addend2

		// 确保结果大于 10
		if sum > 10 {
			key := fmt.Sprintf("%d + %d = %d", addend1, addend2, sum)
			if _, exists := results[key]; !exists {
				computeData.NumberOne = addend1
				computeData.NumberTwo = addend2
				computeData.Symbol = "＋"
				computeData.Result = sum
				results[key] = struct{}{}
				computeList = append(computeList, computeData)
				count++
			}
		}
	}
	return
}

func (ms *MathService) exercise10(limit int) (computeList []MathComputeResult) {
	var computeData MathComputeResult
	rand.Seed(time.Now().UnixNano())
	results := make(map[string]struct{}) // 用于存储唯一的结果
	count := 0

	for count < limit {
		// 生成 1 到 100 之间的被减数和减数
		minuend := rand.Intn(100) + 1
		subtrahend := rand.Intn(100) + 1

		// 计算差值
		difference := minuend - subtrahend

		// 确保结果大于 0
		if difference > 0 {
			key := fmt.Sprintf("%d - %d = %d", minuend, subtrahend, difference)
			if _, exists := results[key]; !exists {
				computeData.NumberOne = minuend
				computeData.NumberTwo = subtrahend
				computeData.Symbol = "－"
				computeData.Result = difference
				results[key] = struct{}{}
				computeList = append(computeList, computeData)
				count++
			}
		}
	}
	return
}

func (ms *MathService) exercise11(limit int) (computeList []MathComputeResult) {
	var computeData MathComputeResult
	rand.Seed(time.Now().UnixNano())
	results := make(map[string]struct{}) // 用于存储唯一的结果
	count := 0

	for count < limit {
		// 随机选择操作类型（加法或减法）
		isAddition := rand.Intn(2) == 0

		var key string
		var result int

		if isAddition {
			// 生成加法表达式
			addend1 := rand.Intn(100) + 1
			addend2 := rand.Intn(100) + 1
			result = addend1 + addend2

			// 确保加法结果大于 10
			if result > 10 {
				key = fmt.Sprintf("%d + %d = %d", addend1, addend2, result)
				computeData.NumberOne = addend1
				computeData.NumberTwo = addend2
				computeData.Symbol = "＋"
				computeData.Result = result
			}
		} else {
			// 生成减法表达式
			minuend := rand.Intn(100) + 1
			subtrahend := rand.Intn(minuend) // 确保被减数大于减数
			result = minuend - subtrahend

			// 确保减法结果大于 0
			if result > 0 {
				key = fmt.Sprintf("%d - %d = %d", minuend, subtrahend, result)
				computeData.NumberOne = minuend
				computeData.NumberTwo = subtrahend
				computeData.Symbol = "－"
				computeData.Result = result
			}
		}

		// 检查并存储唯一结果
		if key != "" {
			if _, exists := results[key]; !exists {
				results[key] = struct{}{}
				computeList = append(computeList, computeData)
				count++
			}
		}
	}
	return
}

func (ms *MathService) exercise12(limit int) (computeList []MathComputeResult) {
	var computeData MathComputeResult
	rand.Seed(time.Now().UnixNano())
	results := make(map[string]struct{}) // 用于存储唯一的结果
	count := 0

	for count < limit {
		// 生成 1 到 5 之间的两个因数
		factor1 := rand.Intn(5) + 1
		factor2 := rand.Intn(5) + 1

		// 计算乘积
		product := factor1 * factor2

		// 确保结果不为 1
		if product != 1 {
			key := fmt.Sprintf("%d × %d = %d", factor1, factor2, product)
			if _, exists := results[key]; !exists {
				computeData.NumberOne = factor1
				computeData.NumberTwo = factor2
				computeData.Symbol = "✖️"
				computeData.Result = factor1 * factor2
				results[key] = struct{}{}
				computeList = append(computeList, computeData)
				count++
			}
		}
	}
	return
}

func (ms *MathService) exercise13(limit int) (computeList []MathComputeResult) {
	var computeData MathComputeResult
	rand.Seed(time.Now().UnixNano())
	results := make(map[string]struct{}) // 用于存储唯一的结果
	count := 0

	for count < limit {
		// 生成 6 到 9 之间的两个因数
		factor1 := rand.Intn(4) + 6
		factor2 := rand.Intn(4) + 6

		// 计算乘积
		product := factor1 * factor2

		// 确保结果不为 1
		if product != 1 {
			key := fmt.Sprintf("%d × %d = %d", factor1, factor2, product)
			if _, exists := results[key]; !exists {
				computeData.NumberOne = factor1
				computeData.NumberTwo = factor2
				computeData.Symbol = "✖️"
				computeData.Result = factor1 * factor2
				results[key] = struct{}{}
				computeList = append(computeList, computeData)
				count++
			}
		}
	}
	return
}

func (ms *MathService) exercise14(limit int) (computeList []MathComputeResult) {
	var computeData MathComputeResult
	rand.Seed(time.Now().UnixNano())
	results := make(map[string]struct{}) // 用于存储唯一的结果
	count := 0

	for count < limit {
		// 生成 1 到 10 之间的两个因数
		factor1 := rand.Intn(10) + 1
		factor2 := rand.Intn(10) + 1

		// 计算乘积
		product := factor1 * factor2

		// 确保结果不为 1
		if product != 1 {
			key := fmt.Sprintf("%d × %d = %d", factor1, factor2, product)
			if _, exists := results[key]; !exists {
				computeData.NumberOne = factor1
				computeData.NumberTwo = factor2
				computeData.Symbol = "✖️"
				computeData.Result = factor1 * factor2
				results[key] = struct{}{}
				computeList = append(computeList, computeData)
				count++
			}
		}
	}
	return
}

func (ms *MathService) exercise15(limit int) (computeList []MathComputeResult) {
	var computeData MathComputeResult
	rand.Seed(time.Now().UnixNano())
	results := make(map[string]struct{}) // 用于存储唯一的结果
	count := 0

	for count < limit {
		// 生成 2 到 5 之间的除数
		divisor := rand.Intn(4) + 2

		// 生成 2 到 5 之间的被除数，确保能被除数整除
		// 并且确保结果不为 1
		dividend := divisor * (rand.Intn(4) + 2)
		result := dividend / divisor

		if dividend%divisor == 0 && result != 1 {
			key := fmt.Sprintf("%d ÷ %d = %d", dividend, divisor, result)
			if _, exists := results[key]; !exists {

				computeData.NumberOne = dividend
				computeData.NumberTwo = divisor
				computeData.Symbol = "➗"
				computeData.Result = dividend / divisor

				results[key] = struct{}{}

				computeList = append(computeList, computeData)
				count++
			}
		}
	}
	return
}

func (ms *MathService) exercise16(limit int) (computeList []MathComputeResult) {
	var computeData MathComputeResult
	rand.Seed(time.Now().UnixNano())
	results := make(map[string]struct{}) // 用于存储唯一的结果
	count := 0

	for count < limit {
		// 生成 6 到 9 之间的除数
		divisor := rand.Intn(4) + 6

		// 生成 6 到 9 之间的被除数，确保能被除数整除
		// 并且确保结果不为 1
		dividend := divisor * (rand.Intn(4) + 2)
		result := dividend / divisor

		if dividend%divisor == 0 && result != 1 {
			key := fmt.Sprintf("%d ÷ %d = %d", dividend, divisor, result)
			if _, exists := results[key]; !exists {

				computeData.NumberOne = dividend
				computeData.NumberTwo = divisor
				computeData.Symbol = "➗"
				computeData.Result = dividend / divisor

				results[key] = struct{}{}

				computeList = append(computeList, computeData)
				count++
			}
		}
	}
	return
}

func (ms *MathService) exercise17(limit int) (computeList []MathComputeResult) {
	var computeData MathComputeResult
	rand.Seed(time.Now().UnixNano())
	results := make(map[string]struct{}) // 用于存储唯一的结果
	count := 0

	for count < limit {
		// 生成 2 到 10 之间的除数（确保除数不能为 1）
		divisor := rand.Intn(9) + 2

		// 生成 2 到 10 之间的被除数，确保能被除数整除
		// 并且确保结果不为 1
		dividend := divisor * (rand.Intn(10/divisor) + 2)
		result := dividend / divisor

		if dividend%divisor == 0 && result != 1 {
			key := fmt.Sprintf("%d ÷ %d = %d", dividend, divisor, result)
			if _, exists := results[key]; !exists {
				computeData.NumberOne = dividend
				computeData.NumberTwo = divisor
				computeData.Symbol = "➗"
				computeData.Result = dividend / divisor

				results[key] = struct{}{}

				computeList = append(computeList, computeData)
				count++
			}
		}
	}
	return
}
