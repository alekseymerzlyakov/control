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

var CountryXlsx = "APP/" + Filename + ".xlsx"

const (
	startRowTest = 5
	startColumn  = 2
	startBodyCol = 12
	BeforeColumn = 9
)

type Cell struct {
	columnNum int
	rowNum    int
}

var Dop_parseRes []string
var valueTemp string
var newelement string

//-------GetColumn----------------------------------------------------
func GetColumn(column int, rownum int, a Data) map[int][]string { //
	//a.empty = false
	getNumTestLine := rownum
	apiListMap := make(map[int][]string)

	for cIdx := 0; cIdx < 3000; cIdx++ {
		cell := a.sheet.Cell(cIdx+column, getNumTestLine)
		if cell.Value() == "" {
			//Empty = false
			break
		} else {
			//Empty = true
			cellNew := a.Replace(cell.String()) //  Подмена переменных
			apiListMap[cIdx] = []string{cellNew}
		}
		//fmt.Println(cellNew)
	}

	return apiListMap
}

func (a Data) GetCell(i int, z int) string {
	return a.sheet.Cell(i, z).Value()
}

var FillColor = "#ffffff"

func (a Data) SetSell(i int, z int, y string) {
	FillColor = "#ffffff"
	if a.errorCount >= 1 && i == 7 {
		FillColor = "#ff0000"
		styles.Fill.Color("#ff0000")
	}
	if a.errorCount == 0 && i == 7 {
		FillColor = "#7CFC00"
		styles.Fill.Color("#7CFC00")
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
	a.errorCount = 0
}

func GetPropValue(key string) string {

	ur := Path() + "/Data/" + Filename + ".json"
	viper.SetConfigFile(ur)
	viper.ReadInConfig()
	viper.AutomaticEnv()
	//viper.WriteConfig()
	get := viper.GetString(key)
	//fmt.Println("GetPropValue key --- >>>     ", key)
	//fmt.Println("GetPropValue get --- >>>     ", get)
	if len(get) == 0 {
		msg :=
			"\n function name:   GetPropValue" + "\n\n  Key ->  " + key +
				"\n apitestname:   ->     " + GetPropValue("apitestname") +
				"\nТест остановлен \n Ключа " + key + " нет в файле - /Data/" + Filename + ".json ->  " +
				"\n\n возможная проблема -> API вернуло не корректный ответ и небыли спаршены данные в результате чего было записано пустое значение" +
				" \n надо искать в log console -> " + key + " \nAPI Response \n  GetPropValue(\"last_response\")  \n\n" + "Страна --->   " + GetPropValue("country_name") + "\n"
		telegram(msg)
		os.Exit(0)
	}

	return get
}

func SetPropValue(key, value string) {
	ur := Path() + "/Data/" + Filename + ".json"

	viper.ReadInConfig()
	viper.AutomaticEnv()
	viper.Set(key, string(value))
	viper.SetConfigFile(ur)
	viper.WriteConfig()
	viper.SafeWriteConfig()

}

//work
func (a *Data) GetRow2(ce int, ro int, apiSheetName string, rownum int) map[int][]string {
	getNumTestLine := rownum
	//Open XLSX file
	xl, err := xlsx.Open(CountryXlsx)
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
		if cell.Value() == "" {
			break
		} //if cell == nil { break }
		cellNew := a.Replace(cell2.String())
		apiListMap[rIdx] = []string{cell.Value(), cellNew}
		getNumTestLine++
	}

	return apiListMap
}

//work
func GetRow(ce int, ro int, rownum int, a Data) map[int][]string {
	getNumTestLine := rownum
	apiListMap := make(map[int][]string)
	for rIdx := 0; rIdx < 3000; rIdx++ {
		cell := a.sheet.Cell(ro, getNumTestLine)
		cell2 := a.sheet.Cell(ro+1, getNumTestLine)
		if cell.Value() == "" {
			break
		} //if cell == nil { break }
		cellNew := a.Replace(cell2.String())
		apiListMap[rIdx] = []string{cell.Value(), cellNew}
		getNumTestLine++
	}
	return apiListMap
}

//-----------------------------------------

func GetHeader(ce int, ro int, rownum int, a Data) map[int][]string {
	getNumTestLine := rownum
	apiListMap := make(map[int][]string)
	cell := a.sheet.Cell(ro, getNumTestLine)

	switch {
	case strings.Contains(cell.Value(), "::"):

		parseHeader := a.GetCell(ro, getNumTestLine)
		strHeader_ := strings.Replace(parseHeader, "\n", "", -1)
		strHeader := strings.Split(strHeader_, "##")

		// Выполнение теста - перебираем список Requests из вкладки
		fmt.Println("------------------------- Preparation of Header-------------------------")
		for rIdx := 0; rIdx < len(strHeader); rIdx++ {
			strHeader[rIdx] = strings.Trim(strHeader[rIdx], " ")
			if len(string(strHeader[rIdx])) <= 1 || strHeader[rIdx] == "" {
				break
			}

			if !strings.Contains(strHeader[rIdx], "::") {
				msg := "В тесте    ---- >    " + a.ApiTestName + "\n" + "Не корректно составлен Header - надо проверить разделитель - ::"
				fmt.Println("В тесте    ---- >    ", a.ApiTestName)
				fmt.Println("Не корректно составлен Header - надо проверить разделитель - ::")

				telegram(msg)
				os.Exit(0)
			}

			generateMap := strings.Split(strHeader[rIdx], "::")
			fmt.Println("strHeader  ---- >    ", generateMap[0])
			fmt.Println("strHeader  ---- >    ", generateMap[1])
			cellNew := a.Replace(generateMap[1])
			apiListMap[rIdx] = []string{generateMap[0], cellNew}
		}

	default:
		for rIdx := 0; rIdx < 3000; rIdx++ {
			cell := a.sheet.Cell(ro, getNumTestLine)
			cell2 := a.sheet.Cell(ro+1, getNumTestLine)
			if cell.Value() == "" {
				break
			} //if cell == nil { break }
			cellNew := a.Replace(cell2.String())
			apiListMap[rIdx] = []string{cell.Value(), cellNew}
			getNumTestLine++
		}
	}

	return apiListMap
}

// WORK
func GetNumTestLine(apiSheetName string, apiTestName string) int {
	var getNumTestLine int
	xl, err := xlsx.Open(CountryXlsx)
	if err != nil {
		fmt.Println("GetNumTestLine not opened")
	}
	var er string
	sheet := xl.SheetByName(apiSheetName)
	for cIdx := 0; cIdx < 3000; cIdx++ {
		cell := sheet.Cell(0, cIdx)
		er = cell.Value()
		if cell.Value() == apiTestName {
			getNumTestLine = cIdx
			break
		}
	}

	value := getNumTestLine + 1

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
	letter := []rune{}

	data := strings.Split(str, "##")
	i, _ := strconv.Atoi(data[1])
	if i == 0 {
		i = 6
	}

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
	b := make([]rune, i)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func (a Data) GenerateData(generate map[int][]string) {
	fmt.Println("generate test data --->    ", generate[0])

	if len(generate[0]) != 0 {
		//if generate[0][0] != "" {
		generateTestData2_ := strings.Replace(generate[0][0], "\n", "", -1)
		generateTestData2 := strings.Split(generateTestData2_, ";")

		// Выполнение теста - перебираем список Requests из вкладки
		for rIdx := 0; rIdx < len(generateTestData2); rIdx++ {

			if string(generateTestData2[rIdx]) == "" {
				break
			} //if cell == nil { break }

			generateMap := strings.Split(generateTestData2[rIdx], "::")
			fmt.Println("generateMap  - ><><><>>>>>>>>>>>    ", generateMap[0])
			fmt.Println("generateMap  - ><><><>>>>>>>>>>>    ", generateMap[1])
			SetPropValue(generateMap[0], a.Replace(generateMap[1]))
		}
	}

}
