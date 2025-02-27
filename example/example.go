package main

import (
	"github.com/eurozulu/colout"
	"log"
	"os"
	"strings"
)

func main() {
	text := "abc123,22/12/2025, Joe, Soap, ¢233.55"
	cw := colout.ColumnWriter{
		Columns: []colout.Column{
			{
				Name:      "id",
				Alignment: colout.Right,
				Width:     10,
			},
			{
				Name:  "created",
				Width: 10,
			},
			{
				Name:  "first-name",
				Width: 25,
			},
			{
				Name:  "surname",
				Width: 25,
			},
			{
				Name:      "amount",
				Width:     10,
				Alignment: colout.Centre,
			},
		},
		ColumnSpacer: "  ",
		Out:          os.Stdout,
	}
	if _, err := cw.WriteStrings(cw.ColumnNames()); err != nil {
		log.Fatal(err)
	}
	if _, err := cw.WriteStrings(strings.Split(text, ",")); err != nil {
		log.Fatal(err)
	}
}
