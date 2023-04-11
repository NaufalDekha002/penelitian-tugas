package module

import (
	"context"
	"fmt"
	"os"

	model "github.com/NaufalDekha002/penelitian-tugas/model"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "tugas5_db",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertPeneliti(db *mongo.Database, col string, nama string, phone_number string, jabatan string, jadwalpenelitian string) (InsertedID interface{}) {
	var peneliti model.Peneliti
	peneliti.Nama = nama
	peneliti.Phone_number = phone_number
	peneliti.Jabatan = jabatan
	peneliti.JadwalPenelitian = jadwalpenelitian
	return InsertOneDoc(db, col, peneliti)
}

func InsertLab(db *mongo.Database, col string, nama string, kategori string) (InsertedID interface{}) {
	var lab model.Lab
	lab.Nama = nama
	lab.Kategori = kategori
	return InsertOneDoc(db, col, lab)
}

func InsertHasilTeliti(db *mongo.Database, col string, nama string) (InsertedID interface{}) {
	var hasilteliti model.HasilTeliti
	hasilteliti.Nama = nama
	return InsertOneDoc(db, col, hasilteliti)
}

func InsertJadwal(db *mongo.Database, col string, durasi string, jam_masuk string, jam_keluar string, hari []string) (InsertedID interface{}) {
	var jadwal model.Jadwal
	jadwal.Durasi = durasi
	jadwal.Jam_masuk = jam_masuk
	jadwal.Jam_keluar = jam_keluar
	jadwal.Hari = hari
	return InsertOneDoc(db, col, jadwal)
}

func InsertPresensi(db *mongo.Database, col string, location string, phone_number string, checkin string, biodata model.Peneliti) (InsertedID interface{}) {
	var presensi model.Presensi
	presensi.Location = location
	presensi.Phone_number = phone_number
	presensi.Checkin = checkin
	presensi.Biodata = biodata
	return InsertOneDoc(db, col, presensi)
}

func GetPresensiFromPhoneNumber(phone_number string, db *mongo.Database, col string) (absen model.Presensi) {
	presensi := db.Collection(col)
	filter := bson.M{"phone_number": phone_number}
	err := presensi.FindOne(context.TODO(), filter).Decode(&absen)
	if err != nil {
		fmt.Printf("GetPenelitiFromJabatan: %v\n", err)
	}
	return absen
}

func GetPenelitiFromJabatan(jabatan string, db *mongo.Database, col string) (staf model.Peneliti) {
	peneliti := db.Collection(col)
	filter := bson.M{"jabatan": jabatan}
	err := peneliti.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("GetPenelitiFromJabatan: %v\n", err)
	}
	return staf
}

func GetHasilTelitiFromNama(nama string, db *mongo.Database, col string) (objek_teliti model.HasilTeliti) {
	hasilteliti := db.Collection(col)
	filter := bson.M{"location": nama}
	err := hasilteliti.FindOne(context.TODO(), filter).Decode(&objek_teliti)
	if err != nil {
		fmt.Printf("GetPresensiFromHasilteliti: %v\n", err)
	}
	return objek_teliti
}

func GetaLabFromKategori(kategori string, db *mongo.Database, col string) (gedung model.Lab) {
	lab := db.Collection(col)
	filter := bson.M{"lab.nama": kategori}
	err := lab.FindOne(context.TODO(), filter).Decode(&gedung)
	if err != nil {
		fmt.Printf("GetaLabTelitiFromKategori: %v\n", err)
	}
	return gedung
}
func GetPenelitiFromPhoneNumber(phone_number string, db *mongo.Database, col string) (staf model.Presensi) {
	peneliti := db.Collection(col)
	filter := bson.M{"phone_number": phone_number}
	err := peneliti.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("GetPenelitiFromPhoneNumber: %v\n", err)
	}
	return staf
}

func GetAllPresensiFromPeneliti(checkin string, db *mongo.Database, col string) (data []model.Presensi) {
	peneliti := db.Collection(col)
	filter := bson.M{"checkin": checkin}
	cursor, err := peneliti.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAllPresensi(db *mongo.Database, col string) (data []model.Presensi) {
	peneliti := db.Collection(col)
	filter := bson.M{}
	cursor, err := peneliti.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
