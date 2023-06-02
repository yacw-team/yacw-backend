package Game

import "fmt"

func IsValidResult(result map[string]interface{}) bool {
	if result["story"] == nil || result["choice"] == nil || result["round"] == nil {
		return false
	}

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

	return true
}
