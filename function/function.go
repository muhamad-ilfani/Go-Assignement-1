package function

import (
	m "Assignment1/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func PrintData(num int) {
	var dataAll m.Datas
	/* Open JSON File*/
	jsonFile, _ := os.Open("./database/data.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	/*Parsing*/
	json.Unmarshal([]byte(byteValue), &dataAll)

	/*Show Data*/
	ShowData(num, dataAll)

}

func ShowData(num int, data m.Datas) {
	dataPrint := map[string]interface{}{
		"Nama":      data.List[num].Nama,
		"Alamat":    data.List[num].Alamat,
		"Pekerjaan": data.List[num].Pekerjaan,
		"Alasan":    data.List[num].Alasan,
	}
	fmt.Println(dataPrint)
}
