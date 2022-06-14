package cekganjilgenap

func CekGanjilGenap(num int) string {
	var result string
	if num % 2 == 0 {
		result = "genap"
	} else {
		result = "ganjil"
	}
	return result
}