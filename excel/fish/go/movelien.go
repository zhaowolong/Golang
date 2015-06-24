package main

import (
	"bufio"
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}

	return files, nil
}

// 写每一个*.dat数据
func write(infile string, sheet *xlsx.Sheet, rawNbr int) {
	fmt.Println("正在处理", infile)
	var id int
	start := strings.LastIndexAny(infile, "\\")
	id, _ = strconv.Atoi(infile[start+1 : len(infile)-4])
	fmt.Println(id)
	var cell *xlsx.Cell
	var row *xlsx.Row

	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)

	var indexi int = 0
	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		str := string(line) // Convert []byte to string
		if len(str) < 5 {
			break
		}
		split := str[1 : len(str)-1]
		array := strings.Split(split, ",")
		row = sheet.AddRow()

		cell = row.AddCell()
		cell.SetInt(id)
		indexi = 0
		for _, v := range array {
			if indexi >= rawNbr-1 {
				break
			}
			indexi = indexi + 1
			cell = row.AddCell()
			value, _ := strconv.Atoi(v)
			cell.SetInt(value)
		}
	}
	return
}

func main() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error
	var RawNbr int = 4
	file = xlsx.NewFile()
	sheetName := [2]string{"Big", "Small"}
	datafile := [2]string{"../data/moveline/big", "../data/moveline/small"}

	for i, v := range sheetName {
		sheet = file.AddSheet(v)
		// 每列标题
		text := [4]string{"ID", "X", "Y", "Angle"}
		for i := 0; i < RawNbr; i++ {
			if i == 0 {
				row = sheet.AddRow()
			}
			cell = row.AddCell()
			cell.Value = text[i]
		}

		files, _ := ListDir(datafile[i], ".dat")
		for j := 0; j < len(files); j++ {
			write(files[j], sheet, RawNbr)
		}
	}
	fmt.Println("正在向excel写数据......")
	err = file.Save("../outexcel/单条鱼路径.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
