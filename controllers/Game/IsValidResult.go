package Game

import "fmt"

// 检查生成的格式是否符合预期
func IsValidResult(result map[string]interface{}) bool {
	// 检查字段是否存在且不为空
	if result["story"] == nil || result["choice"] == nil || result["round"] == nil {
		return false
	}

	// 进一步检查字段值的类型和合法性
	if _, ok := result["story"].(string); !ok {

		return false
	}

	choices, ok := result["choice"].([]interface{})
	if !ok || len(choices) == 0 {

		return false
	}

	validKeys := map[string]bool{
		"A": true,
		"B": true,
		"C": true,
		"D": true,
	}

	for _, choice := range choices {
		choiceMap, ok := choice.(map[string]interface{})
		if !ok || len(choiceMap) != 1 {

			return false
		}

		for key := range choiceMap {
			if !validKeys[key] {

				return false
			}
		}
	}

	if round, ok := result["round"].(float64); !ok || round <= 0 {

		return false
	}

	// 根据需要添加其他验证规则

	return true
}
