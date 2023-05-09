package utils

func preNUm(data byte) int {
	var mask byte = 0x80
	var num = 0
	for i := 0; i < 8; i++ {
		if (data & mask) == mask {
			num++
			mask = mask >> 1
		} else {
			break
		}
	}
	return num
}

func isUtf8(data []byte) bool {
	i := 0
	for i < len(data) {
		if (data[i] & 0x80) == 0x00 {
			i++
			continue
		} else if num := preNUm(data[i]); num > 2 {
			i++
			for j := 0; j < num-1; j++ {
				if (data[i] & 0xc0) != 0x80 {
					return false
				}
				i++
			}
		} else {
			return false
		}
	}
	return true
}

func ApiKeyCheck(apiKey string) bool {
	if len(apiKey) == 51 {
		if apiKey[0] == 's' && apiKey[1] == 'k' && apiKey[2] == '-' {
			a := []byte(apiKey)
			if isUtf8(a) == true {
				return true
			}
		}
	}
	return false
}
