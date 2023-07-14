package main

import (
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
	// "github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	file := excelize.NewFile()

	data := [][]interface{}{
		{"Ad", "Soyad", "Yaş"},
		{"Ahmet", "Yılmaz", 28},
		{"Ayşe", "Demir", 32},
		{"Mehmet", "Kara", 25},
	}

	sheetName := "Sheet1"
	for row, rowData := range data {
		for col, cellData := range rowData {
			cell := excelize.ToAlphaString(col+1) + fmt.Sprint(row+1)
			file.SetCellValue(sheetName, cell, cellData)
		}
	}

	err := file.SaveAs("veriler.xlsx")
	if err != nil {
		fmt.Println("Excel dosyası kaydedilirken bir hata oluştu:", err)
		os.Exit(1)
	}

	fmt.Println("Veriler başarıyla Excel tablosuna aktarıldı.")
}
