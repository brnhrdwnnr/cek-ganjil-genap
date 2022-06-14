package cekganjilgenap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableCekGanjilGenap(t *testing.T) {
	tests := [] struct {
		name string
		expected string
		param []int
	}{
			{
				name: "01",
				expected: "ganjil, genap, ganjil",
				param: []int{1, 2, 3},
			},
			{
				name: "02",
				expected: "ganjil, genap, ganjil, genap",
				param: []int{1, 2, 3, 4},
			},
			{
				name: "03",
				expected: "ganjil, ganjil, ganjil, ganjil, ganjil, ganjil, ganjil",
				param: []int{7, 7, 7, 7, 7, 7, 7},
			},
			{
				name: "04",
				expected: "genap, genap, genap",
				param: []int{2, 2, 2},
			},
	}
	assert := assert.New(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CekGanjilGenap(test.param...)
			assert.Equal(test.expected, result, "they should be equal")
		})
	}
}

func BenchmarkCekGanjilGenapTable(b *testing.B) {
	tests := [] struct {
		name string
		param []int
	}{
			{
				name: "01",
				param: []int{1, 2, 3, 4, 5},
			},
			{
				name: "02",
				param: []int{1, 2, 3, 4, 5, 6},
			},
			{
				name: "03",
				param: []int{1, 2, 3, 4, 5, 6, 7},
			},
			{
				name: "04",
				param: []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
	}
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CekGanjilGenap(test.param...)
			}
		})
	}
}