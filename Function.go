package main

import (
	_"bytes"
	"fmt"
	"github.com/magiconair/properties"
	"github.com/spf13/viper"
	_"github.com/tinyhubs/properties"
    "github.com/plandem/xlsx"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

)
//var ClientID string = GetString("ClientID")
var CountryName = GetPropValue("countryname") + ".xlsx"

type Conf struct {
	Key interface{}
	Value interface{}
}

func Config(key,  value string)  {

	ur := Path() + "/data.json"

	viper.SetConfigFile(ur)
	//// Searches for config file in given paths and read it
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("viper+++++++++++>>>>>>>>>>>>>          ", ur)
	}
	viper.AutomaticEnv()
	//
	//fmt.Println("viper+++++++++++>>>>>>>>>>>>>          ", viper.Get("CountryName"))
	//viper.BindEnv("port", "PORT")

	viper.Get(key)
	viper.Set(key, value)

	viper.WriteConfig()
}




func GetPropValue(key string) string {

	p := properties.MustLoadFile(Path()+"/testData.properties", properties.UTF8)
	value := p.MustGet(key)
	fmt.Println("GetPropValue -------->>>>>>>      ", value)
	return value

}

//Open Properties file------------------------------
func SetPropValue(key, value string)  {

	p := properties.MustLoadFile(Path()+"/testData.properties", properties.UTF8)
	p.Set(key, value)
	p.SetValue("key", value)
	fmt.Println("SetValue -------->>>>>>>      ", key)
}


//work
func GetRow(ce int, ro int, apiSheetName string, ApiTestName string) (map[int][]string) {
	//Open XLSX file
	xl, err := xlsx.Open(CountryName)
	if err != nil {
		fmt.Println("xlsx not opened")
	}
	getNumTestLine := GetNumTestLine(apiSheetName, ApiTestName)  //  поиск начала теста
	fmt.Println("getNumTestLine   ------    ", getNumTestLine)
	sheet := xl.SheetByName(apiSheetName)

	apiListMap := make(map[int][]string)

	for rIdx := 0; rIdx < 100; rIdx++ {

		cell := sheet.Cell(ro, getNumTestLine)
		cell2 := sheet.Cell(ro+1, getNumTestLine)
		//fmt.Println("xxxxxxxxxxx   ------    ",ro, getNumTestLine)
		//fmt.Println("xxxxxxxxxxx   ------    ",ro+1, getNumTestLine)
		//fmt.Println("qqqqqq   ------    ",apiSheetName ,cell.Value() , cell2.Value())

		if cell.Value() == "" { break } //if cell == nil { break }

		cellNew := Replace(cell2.String())
		apiListMap[rIdx] = []string{cell.Value(), cellNew}
		getNumTestLine++
	}
	return apiListMap
}


////work
//func GetRow(ce int, ro int, apiSheetName string, ApiTestName string) (map[int][]string) {
//	//Open XLSX file
//	xl, err := xlsx.Open(CountryName)
//	if err != nil {
//	fmt.Println("xlsx not opened")
//}
//	getNumTestLine := GetNumTestLine(apiSheetName, ApiTestName)
//
//
//
//	sheet := xl.SheetByName(apiSheetName)
//
//	apiListMap := make(map[int][]string)
//	//_, totalRows := sheet.Dimension()
//		for rIdx := 0; rIdx < 100; rIdx++ {
//		cell := sheet.Cell(getNumTestLine, rIdx+ro)
//		cell2 := sheet.Cell(getNumTestLine+1, rIdx+ro)
//			fmt.Println("xxxxxxxxxxx   ------    ",getNumTestLine)
//			fmt.Println("qqqqqq   ------    ",apiSheetName ,cell.Value() , cell2)
//
//		if cell.Value() == "" { break } //if cell == nil { break }
//		apiListMap[rIdx] = []string{cell.Value(), cell2.Value()}
//}
//	return apiListMap
//}


// WORK
func GetNumTestLine(apiSheetName string, apiTestName string)  int {
	var getNumTestLine int
	//Open XLSX file
	xl := XlsxOpen()

	sheet := xl.SheetByName(apiSheetName)
	for cIdx := 0; cIdx < 100; cIdx++ {
		cell := sheet.Cell(0, cIdx)
		if cell.Value() == apiTestName {getNumTestLine = cIdx
		break } //if cell == nil { break }
	}

	value := getNumTestLine+1
	xl.Close()
	return value
}


//-------GetColumn----------------------------------------------------
func GetColumn( column int, apiSheetName string, apiTestName string)  (map[int][]string) { //
	//Open XLSX file
	xl := XlsxOpen()
	//if err != nil {
	//	fmt.Println("xlsx not opened")
	//}

	getNumTestLine := GetNumTestLine(apiSheetName, apiTestName)
	sheet := xl.SheetByName(apiSheetName)
	apiListMap := make(map[int][]string)

	for cIdx := 0; cIdx < 200; cIdx++ {
		cell := sheet.Cell(cIdx+column, getNumTestLine)
		if cell.Value() == "" { break } //if cell == nil { break }
		cellNew := Replace(cell.String())			//  Подмена переменных
		apiListMap[cIdx] = []string{cellNew}
		fmt.Println(cellNew)
	}




	//apiListMap := make(map[int][]string)
	////_, totalRows := sheet.Dimension()
	//	for cIdx := 0; cIdx < 100; cIdx++ {
	//		cell := sheet.Cell(cIdx+ro, ce)
	//		if cell.Value() == "" { break } //if cell == nil { break }
	//		//fmt.Println("cell 2-3", cell)
	//		apiListMap[cIdx] = []string{cell.Value()}
	//		fmt.Printf(apiListMap[cIdx][0])
	//	}
	//
	//
	//for key, value := range apiListMap { // Order not specified
	//	fmt.Println(key, value)
	//}
	return apiListMap
}








//------------------------------------------------------------
func GetHeader(cf int, kz int, ap string) map[string]string {
	//Open XLSX file
	xl := XlsxOpen()

	sheet := xl.SheetByName(ap)
	apiListMap := make(map[string]string)

	//totalCols, totalRows := sheet.Dimension()
	var tem string
	var cell string
	for zIdx := cf; zIdx < cf+2; zIdx++ {

		for fIdx := kz; fIdx < fIdx+2; fIdx++ {
			tem = cell

			cell = sheet.Cell(fIdx, zIdx).Value()
			if cell == "" || len(cell) <= 1 {fIdx = 1000
			break}
			if cell != "" && tem != "" {apiListMap[tem] = cell}
			if tem == "" && cell == "" {zIdx = 1000
				fIdx = 1000
				break}

		}

	}



	for key, value := range apiListMap { // Order not specified
		fmt.Println(key, value)
	}
	return apiListMap
}























func XlsxOpen() *xlsx.Spreadsheet {
	//Open XLSX file
	xl, err := xlsx.Open(CountryName)
	if err != nil {
		fmt.Println("xlsx not opened")
	}
	return xl
}

//var path string
func Path() string {
	path, err := os.Getwd()
	if err != nil {
		//	log.Println(err)
	}
	return path
}




//work
//func GetRow2(ce int, ro int, apiSheetName string, ApiTestName string) (map[int][]string) {
//	//Open XLSX file
//	xl, err := xlsx.Open(CountryName)
//	if err != nil {
//		fmt.Println("xlsx not opened")
//	}
//	getNumTestLine := GetNumTestLine(apiSheetName, ApiTestName)
//	fmt.Println("getNumTestLine   ------    ", getNumTestLine)
//	sheet := xl.SheetByName(apiSheetName)
//
//	apiListMap := make(map[int][]string)
//
//	for rIdx := 0; rIdx < 100; rIdx++ {
//
//
//		cell := sheet.Cell(ro, getNumTestLine)
//		cell2 := sheet.Cell(ro+1, getNumTestLine)
//		fmt.Println("xxxxxxxxxxx   ------    ",ro, getNumTestLine)
//		fmt.Println("xxxxxxxxxxx   ------    ",ro+1, getNumTestLine)
//		fmt.Println("qqqqqq   ------    ",apiSheetName ,cell.Value() , cell2.Value())
//
//		if cell.Value() == "" { break } //if cell == nil { break }
//		apiListMap[rIdx] = []string{cell.Value(), cell2.Value()}
//		getNumTestLine++
//	}
//	return apiListMap
//}

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

func GenerateData(generate map[int][]string)  {
	generateTestData2 := strings.Split(generate[0][0], ";")
	fmt.Println("generateTestData2  - >    ", len(generateTestData2))

	// Выполнение теста - перебираем список Requests из вкладки
	for rIdx := 0; rIdx < len(generateTestData2); rIdx++ {
		//fmt.Println(header[rIdx][0],header[rIdx][1])
		if string(generateTestData2[rIdx]) == "" { break } //if cell == nil { break }
		fmt.Println("generateTestData2  - ><><><>    ", generateTestData2[rIdx])
		generateMap := strings.Split(generateTestData2[rIdx], "##")
		fmt.Println("generateMap  - ><><><>>>>>>>>>>>    ", generateMap[0])
		fmt.Println("generateMap  - ><><><>>>>>>>>>>>    ", generateMap[1])


		//SetPropValue(generateMap[0],generateMap[1])
	}

}