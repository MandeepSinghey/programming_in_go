package main

import (
	"fmt"
	"os"
)

//create a slice of slice for holding the numbers
//  9 rows and 7 columns
var bigDigits = [][]string{
	{
		" 000 ",
		"0   0",
		"0   0",
		"0   0",
		"0   0",
		"0   0",
		" 000 "},
	{
		"   1 ",
		" 111 ",
		"   1 ",
		"   1 ",
		"   1 ",
		"   1 ",
		" 1111"},
	//  col1,     col2,       col3,      col4,     col5,    col6,   col7
	//{" 2222 ", " 2  2 ", "    2 ", "  2   ", " 2    ", "2     ", "222222"},
	{
		" 2222 ",
		" 2  2 ",
		"    2 ",
		"  2   ",
		" 2    ",
		"2     ",
		"222222"},
	{
		"33333",
		"    3",
		"    3",
		"33333",
		"    3",
		"    3",
		"33333"},
	{
		"4   4",
		"4   4",
		"4   4",
		"44444",
		"    4",
		"    4",
		"    4"},
	{
		"55555",
		"5    ",
		"5    ",
		"55555",
		"    5",
		"    5",
		"55555"},
	{
		"66666",
		"6    ",
		"6    ",
		"66666",
		"6   6",
		"6   6",
		"66666"},
	{
		"77777",
		"7   7",
		"7   7",
		"    7",
		"    7",
		"    7",
		"    7"},
	{
		"88888",
		"8   8",
		"8   8",
		"88888",
		"8   8",
		"8   8",
		"88888"},
	{
		"99999",
		"9   9",
		"9   9",
		"99999",
		"    9",
		"    9",
		"99999"},
}

func main() {
	//check if the user passed the number to be printed
	//if length is just 1 then only file name was used to ru the program
	if len(os.Args) == 1 {
		fmt.Println("Input number after file name to see the ouput")
		//write a function to show a default ouput
	}
	//read the first
	stringOfDigits := os.Args[1]
	//fmt.Println(stringOfDigits)
	// assuming all rows are of same length =7 columns
	// loop over those 7 columns in each row
	for row := range bigDigits[0] {
		line := " "
		// iterate over enterted number
		for column := range stringOfDigits {
			/*
				example,lets say we have the character'3'( which is code point 51),
				we can get its integer value by doing the subtraction '3' - '0'
				(which is code point 48)
				(i.e., 51− 48) which results in an integer (of type byte) of value 3.
			*/
			digit := stringOfDigits[column] - '0'

			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + " "
			}

		}
		fmt.Println(line)

	}

}
