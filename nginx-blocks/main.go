package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// NginxBlock struct
type NginxBlock struct {
	StartLine   string
	EndLine     string
	AllContents string
	// split lines by \n on AllContents,
	// use make to create *[],
	// first create make([]*Type..)
	// then use &var to make it *
	// AllLines          *[]*string
	// NestedBlocks      []*NginxBlock
	TotalBlocksInside int
}

func main() {
	data, err := ioutil.ReadFile("nginx.conf")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	blocks := ""
	// blockActive := false
	var blockArray []string
	fmt.Println("Contents of file:")
	fmt.Println("length : ", len(string(data)))

	blockcharAt := make(map[int]int)
	blockCount := 0
	lineNumber := 1
	var previousLine []int
	var nginxBlocks []NginxBlock
	for i, c := range data {
		// fmt.Printf(string(c))
		if string(c) == "\n" {
			lineNumber++
		}
		if string(c) == "{" {
			// fmt.Println(lineNumber)
			previousLine = append(previousLine, lineNumber)
			blockCount++
			blockcharAt[blockCount] = i + 1
		}
		if string(c) == "}" {
			// fmt.Printf("block between : %v - %v", previousLine[len(previousLine)-1], lineNumber)
			// fmt.Println()

			start := strconv.Itoa(previousLine[len(previousLine)-1])
			end := strconv.Itoa(lineNumber)
			b := NginxBlock{StartLine: string(start), EndLine: end}

			previousLine = previousLine[:len(previousLine)-1]
			tempBlock := blocks + string(c)
			blockNow := tempBlock[blockcharAt[blockCount]-1:]
			b.AllContents = blockNow
			blockArray = append(blockArray, blockNow)
			nginxBlocks = append(nginxBlocks, b)
			blockCount--
		}
		blocks = blocks + string(c)
	}
	// fmt.Println("length block array ====>", len(blockArray))
	// for _, block := range blockArray {
	// 	fmt.Println("======>")
	// 	numberofBlocks := 0
	// 	fmt.Println(block)
	// 	for _, c := range block {
	// 		if string(c) == "}" {
	// 			numberofBlocks++
	// 		}
	// 	}
	// 	numberofBlocks--
	// 	fmt.Println("Number of nested block : ", numberofBlocks)
	// 	fmt.Println("======>")
	// }

	for _, bl := range nginxBlocks {
		fmt.Println("=============>")
		fmt.Println("startline : ", bl.StartLine)
		fmt.Println("endline   : ", bl.EndLine)
		fmt.Println(bl.AllContents)
		fmt.Println("Total Blocks   : ", bl.TotalBlocksInside)
		fmt.Println("=============>")
		fmt.Println()
		fmt.Println()
	}
}
