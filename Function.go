package main

import (
	_ "bytes"
	"fmt"
	"github.com/plandem/xlsx"
	"github.com/plandem/xlsx/format/styles"
	"github.com/spf13/viper"
	_ "github.com/tidwall/gjson"
	_ "github.com/tinyhubs/properties"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)
var CountryName = GetPropValue("countryname") + ".xlsx"

const (
	startRowTest = 5
	startColumn = 2
	startBodyCol = 12
	)


type Cell struct {
	columnNum int
	rowNum int
}
var Dop_parseRes []string
var valueTemp string
var newelement string



//-------GetColumn----------------------------------------------------
func GetColumn(column int, rownum int, a Data)  (map[int][]string) { //
	a.empty = false
	getNumTestLine := rownum
	apiListMap := make(map[int][]string)

	for cIdx := 0; cIdx < 3000; cIdx++ {
		cell := a.sheet.Cell(cIdx+column, getNumTestLine)
		if cell.Value() == "" {
			a.empty = true
			break
		}
		cellNew := Replace(cell.String())			//  Подмена переменных
		apiListMap[cIdx] = []string{cellNew}
		//fmt.Println(cellNew)
	}

	return apiListMap
}


func (a Data) GetCell(i int, z int) string {
	return a.sheet.Cell(i, z).Value()
}

func (a Data) SetSell(i int, z int, y string)  {
	FillColor := "#ffffff"
	if a.Error.errorCount > 0 && i == 7 {
		FillColor = "#ff0000"
	}
	if a.Error.errorCount == 0 && i == 7 {
		FillColor = "#7CFC00"
	}

	cell := a.sheet.Cell(i, z)
	cell.SetValue(y)
	cell.SetStyles(styles.New(
		styles.Font.Bold,
		styles.Font.Color("#000000"),
		styles.Fill.Type(styles.PatternTypeSolid),
		styles.Fill.Color(FillColor),
		//styles.Border.Color("#009000"),
		//styles.Border.Type(styles.BorderStyleMedium),
	))

}

func GetPropValue(key string) string {
	ur := Path() + "/data.json"
	viper.SetConfigFile(ur)
	viper.ReadInConfig()
	viper.AutomaticEnv()
	//viper.WriteConfig()
	get := viper.GetString(key)

	//fmt.Println("Такой переменной нет в файле - data.json ->  ", key, err)

	if get == "" {
		msg := GetPropValue("countryname") + "\nТест остановлен \n Такой переменной нет в файле - data.json ->  " + key
		telegram(msg)
		os.Exit(0)
	}

	return get
}

func SetPropValue(key, value string)  {
	ur := Path() + "/data.json"
	viper.SetConfigFile(ur)
	viper.ReadInConfig()
	viper.AutomaticEnv()
	viper.Set(key, string(value))
	viper.WriteConfig()

}

//work
func GetRow2(ce int, ro int, apiSheetName string, rownum int) (map[int][]string) {
	getNumTestLine := rownum
	//Open XLSX file
	xl, err := xlsx.Open(CountryName)
	if err != nil {
		fmt.Println("GetRow not opened")
	}

	//getNumTestLine := GetNumTestLine(apiSheetName, ApiTestName)  //  поиск начала теста
	fmt.Println("getNumTestLine   ------    ", getNumTestLine)
	sheet := xl.SheetByName(apiSheetName)
	apiListMap := make(map[int][]string)
	for rIdx := 0; rIdx < 3000; rIdx++ {

		cell := sheet.Cell(ro, getNumTestLine)
		cell2 := sheet.Cell(ro+1, getNumTestLine)
		if cell.Value() == "" { break } //if cell == nil { break }
		cellNew := Replace(cell2.String())
		apiListMap[rIdx] = []string{cell.Value(), cellNew}
		getNumTestLine++
	}

	return apiListMap
}


//work
func GetRow(ce int, ro int, rownum int, a Data) (map[int][]string) {
	getNumTestLine := rownum

	apiListMap := make(map[int][]string)

	for rIdx := 0; rIdx < 3000; rIdx++ {

		cell := a.sheet.Cell(ro, getNumTestLine)
		cell2 := a.sheet.Cell(ro+1, getNumTestLine)
		if cell.Value() == "" { break } //if cell == nil { break }
		cellNew := Replace(cell2.String())
		apiListMap[rIdx] = []string{cell.Value(), cellNew}
		getNumTestLine++
	}

	return apiListMap
}

// WORK
func GetNumTestLine(apiSheetName string, apiTestName string)  int {
	var getNumTestLine int
	xl, err := xlsx.Open(CountryName)
	if err != nil {
		fmt.Println("GetNumTestLine not opened")
	}
	var er string
	sheet := xl.SheetByName(apiSheetName)
	for cIdx := 0; cIdx < 3000; cIdx++ {
		cell := sheet.Cell(0, cIdx)
		er = cell.Value()
		if cell.Value() == apiTestName {getNumTestLine = cIdx
		break }
	}

	value := getNumTestLine+1

	if er == "" {
		fmt.Println("Ошибка. В вкладке APIList есть ссылка на не существующий тест")
		os.Exit(0)
	}

	return value
}

//var path string
func Path() string {
	path, err := os.Getwd()
	if err != nil {
	}
	return path
}


func Random(str string) string {
	var letter[]rune
	 data := strings.Split(str, "##")
	i, _ := strconv.Atoi(data[1])
	if i == 0 { i = 6 }

	switch {
	case strings.Contains(str, "RanString"):
		letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	case strings.Contains(str, "RanInt"):
		letter = []rune("0123456789")
	case strings.Contains(str, "RanStrInt"):
		letter = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	default:
	}

	rand.Seed(time.Now().UnixNano())
	b := make([]rune,i)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func (a Data) GenerateData(generate map[int][]string)  {
	fmt.Println("generate  - ><><><>    ", generate[0])

	if a.empty {

		generateTestData2 := strings.Split(generate[0][0], ";")

		// Выполнение теста - перебираем список Requests из вкладки
		for rIdx := 0; rIdx < len(generateTestData2); rIdx++ {

			if string(generateTestData2[rIdx]) == "" {
				break
			} //if cell == nil { break }

			generateMap := strings.Split(generateTestData2[rIdx], "::")
			fmt.Println("generateMap  - ><><><>>>>>>>>>>>    ", generateMap[0])
			fmt.Println("generateMap  - ><><><>>>>>>>>>>>    ", generateMap[1])
			SetPropValue(generateMap[0], Replace(generateMap[1]))
			}
	}

}



