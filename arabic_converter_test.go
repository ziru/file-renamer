package main_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	fr "github.com/ziru/file-renamer"
)

func TestConvertChineseNumberToArabicNumber(t *testing.T) {
	tests := [][]string{
		{"一百二十", "120"},
		{"三千五百二十", "3520"},
		{"四", "4"},
		{"十一", "11"},
		{"十", "10"},
		{"二十五", "25"},
		{"一百零八", "108"},
		{"一百一十", "110"},
		{"一百一十六", "116"},
		{"一百五十六", "156"},
	}

	req := require.New(t)
	for _, pair := range tests {
		req.Equal(pair[1], fr.ConvertChineseNumberToArabicNumber(pair[0]))
	}
}

func TestConvertSingleComponent(t *testing.T) {
	req := require.New(t)

	req.Equal(1, fr.ConvertSingleComponent("一"))
	req.Equal(2, fr.ConvertSingleComponent("二"))
	req.Equal(3, fr.ConvertSingleComponent("三"))
	req.Equal(4, fr.ConvertSingleComponent("四"))
	req.Equal(5, fr.ConvertSingleComponent("五"))
	req.Equal(6, fr.ConvertSingleComponent("六"))
	req.Equal(7, fr.ConvertSingleComponent("七"))
	req.Equal(8, fr.ConvertSingleComponent("八"))
	req.Equal(9, fr.ConvertSingleComponent("九"))
	req.Equal(0, fr.ConvertSingleComponent("零"))
}
