package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/FaizArsyiP/FINALPROJECT/src/db"
	"github.com/FaizArsyiP/FINALPROJECT/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Karyawan struct {
	ID           int    `json:"_id,omitempty"`
	Nama         string `json:"Nama"`
	NIK          uint64 `json:"NIK"`
	Pendidikan   string `json:"Pendidikan"`
	TanggalMasuk string `json:"TanggalMasuk"`
	StatusKerja  string `json:"StatusKerja"`
}

func KaryawanHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		db, err := db.DBConnection()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		daftarKaryawan := db.MongoDB.Collection("Karyawan")
		projection := bson.D{
			{Key: "Nama", Value: 1},
			{Key: "TanggalMasuk", Value: 1},
			{Key: "StatusKerja", Value: 1},
		}

		curEm, err := daftarKaryawan.Find(context.TODO(), bson.D{}, options.Find().SetProjection(projection))

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		var Karyawan model.Karyawan
		var KaryawanList []model.Karyawan

		for curEm.Next(context.TODO()) {
			err := curEm.Decode(&Karyawan)
			if err != nil {
				http.Error(w, "Error decoding book data", http.StatusInternalServerError)
				return
			}
			KaryawanList = append(KaryawanList, Karyawan)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(KaryawanList)
		return

	case "POST":
		var data model.Karyawan
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		db, err := db.DBConnection()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer db.MongoDB.Client().Disconnect(context.TODO())

		daftarKaryawan := db.MongoDB.Collection("Karyawan")

		_, err = daftarKaryawan.InsertOne(context.TODO(), data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Karyawan berhasil ditambahkan"))

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
}
