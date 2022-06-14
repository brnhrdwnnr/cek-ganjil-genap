package cekganjilgenap

func CekGanjilGenap(num ... int) string {
	var result string
	for index, value := range num {
		if index != len(num) - 1 && value % 2 == 0  {
			result += "genap, "
		} else if index != len(num) - 1 && value % 2 != 0 {
			result += "ganjil, "
		} else if index == len(num) - 1 && value % 2 == 0  {
			result += "genap"
		} else if index == len(num) - 1 && value % 2 != 0 {
			result += "ganjil"
		}	
	}
	return result
}