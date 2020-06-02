package main

import (
	"fmt"
	"html/template"
	"net/http"

	// "io/ioutil"
	"encoding/xml"
	"log"
)

var nimData = []byte(`
<nimindex>
	<nim>
		<nama>Adi</nama>
		<jurusan>Informatika</jurusan>
		<kelas>IF-41-05</kelas>
		<angkatan>2017</angkatan>
	</nim>
	<nim>
		<nama>Arya</nama>
		<jurusan>Industri</jurusan>
		<kelas>TI-41-12</kelas>
		<angkatan>2017</angkatan>
	</nim>
	<nim>
		<nama>Surya</nama>
		<jurusan>Industri</jurusan>
		<kelas>TI-41-12</kelas>
		<angkatan>2017</angkatan>
	</nim>
	<nim>
		<nama>Andika</nama>
		<jurusan>Informatika</jurusan>
		<kelas>IF-41-10</kelas>
		<angkatan>2017</angkatan>
	</nim>
</nimindex>
`)

type Info struct {
	Nama     []string `xml:"nim>nama"`
	Jurusan  []string `xml:"nim>jurusan"`
	Kelas    []string `xml:"nim>kelas"`
	Angkatan []string `xml:"nim>angkatan"`
}

type DetailInfo struct {
	Jurusan  string
	Kelas    string
	Angkatan string
}

type info struct {
	Judul  string
	Detail map[string]DetailInfo
}

func new(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa,  is near!")
}
func baru_bray(w http.ResponseWriter, r *http.Request) {
	bytes := nimData
	data_mahasiswa := make(map[string]DetailInfo)
	var s Info
	if err := xml.Unmarshal(bytes, &s); err != nil {
		log.Fatal(err)
	}
	for idx, _ := range s.Nama {
		data_mahasiswa[s.Nama[idx]] = DetailInfo{s.Jurusan[idx], s.Kelas[idx], s.Angkatan[idx]}
	}
	p := info{Judul: "Data Mahasiswa", Detail: data_mahasiswa}
	t, _ := template.ParseFiles("baru.html")
	fmt.Println(t.Execute(w, p))
}

func main() {
	http.HandleFunc("/", baru_bray)
	http.HandleFunc("/new/", new)
	http.ListenAndServe(":8000", nil)
}
