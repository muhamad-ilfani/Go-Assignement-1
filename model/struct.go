package model

type Data struct {
	Nama      string `json:"Nama"`
	Alamat    string `json:"Alamat"`
	Pekerjaan string `json:"Pekerjaan"`
	Alasan    string `json:"Alasan"`
}

type Datas struct {
	List []Data `json:"Data"`
}
