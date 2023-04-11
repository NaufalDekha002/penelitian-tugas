package NPM

import (
	"fmt"

	"testing"

	"github.com/NaufalDekha002/penelitian-tugas/model"
	"github.com/NaufalDekha002/penelitian-tugas/module"
)

func TestInsertPeneliti(t *testing.T) {
	nama := "Salsabila Irbah"
	phone_number := "083543242546"
	jabatan := "Chemical Analyst"
	JadwalPenelitian := "Data Field Cross Check"

	hasil := module.InsertPeneliti(module.MongoConn, "peneliti", nama, phone_number, jabatan, JadwalPenelitian)
	fmt.Println(hasil)
}

func TestInsertLab(t *testing.T) {
	nama := "Data Maintenance Room"
	kategori := "Gedung"

	hasil := module.InsertLab(module.MongoConn, "Lab", nama, kategori)
	fmt.Println(hasil)
}

func TestInsertJadwal(t *testing.T) {
	durasi := "2 Jam"
	jam_masuk := "08.00"
	jam_keluar := "10.00"
	hari := []string{"senin"}
	hasil := module.InsertJadwal(module.MongoConn, "jadwal", durasi, jam_masuk, jam_keluar, hari)
	fmt.Println(hasil)
}

func TestInsertHasilTeliti(t *testing.T) {
	nama := "Chemical Project"
	hasil := module.InsertHasilTeliti(module.MongoConn, "hasil teliti", nama)
	fmt.Println(hasil)
}

func TestInsertPresensi(t *testing.T) {
	location := "Gedung 300"
	phone_number := "083543242546"
	checkin := "Masuk"
	biodata := model.Peneliti{
		Nama:             "Salsabila Irbah",
		Phone_number:     "083543242546",
		Jabatan:          "Head Chemical Analyst",
		JadwalPenelitian: "Data Field Cross Check",
	}
	absen := module.InsertPresensi(module.MongoConn, "hasil", location, phone_number, checkin, biodata)
	fmt.Println(absen)
}

func TestGetPenelitiFromPhoneNumber(t *testing.T) {
	phone_number := "083543242546"
	staf := module.GetPenelitiFromPhoneNumber(phone_number, module.MongoConn, "peneliti")
	fmt.Println(staf)
}

func TestGetHasilTelitiFromNama(t *testing.T) {
	nama := "Chemical Project"
	objek_teliti := module.GetHasilTelitiFromNama(nama, module.MongoConn, "hasil teliti")
	fmt.Println(objek_teliti)
}

func TestGetPresensiFromPhoneNumber(t *testing.T) {
	phone_number := "Gedung 300"
	absen := module.GetHasilTelitiFromNama(phone_number, module.MongoConn, "hasil")
	fmt.Println(absen)
}
