package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/FaizArsyiP/FINALPROJECT/src/db"
	"github.com/FaizArsyiP/FINALPROJECT/src/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Buku struct {
	ID      primitive.ObjectID `json:"_id,omitempty"`
	Judul   string             `json:"Judul"`
	Penulis string             `json:"Penulis"`
	Tahun   uint16             `json:"Tahun"`
	Stok    uint8              `json:"Stok"`
	Harga   uint64             `json:"Harga"`
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		id := vars["id"]

		db, err := db.DBConnection()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer db.MongoDB.Client().Disconnect(context.TODO())

		daftarBuku := db.MongoDB.Collection("BookList")

		if id != "" {

			objID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				http.Error(w, "Invalid book ID format", http.StatusBadRequest)
				return
			}

			var buku model.Buku
			err = daftarBuku.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&buku)
			if err != nil {
				http.Error(w, "Buku tidak ditemukan", http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(buku)
			return
		}

		projection := bson.D{
			{Key: "Judul", Value: 1},
			{Key: "Penulis", Value: 1},
			{Key: "Harga", Value: 1},
		}
		curBook, err := daftarBuku.Find(context.TODO(), bson.D{}, options.Find().SetProjection(projection))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer curBook.Close(context.TODO())

		var buku model.Buku
		var bukuList []model.Buku

		for curBook.Next(context.TODO()) {
			err := curBook.Decode(&buku)
			if err != nil {
				http.Error(w, "Error decoding book data", http.StatusInternalServerError)
				return
			}
			bukuList = append(bukuList, buku)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bukuList)
		return

	case "POST":
		var data model.Buku
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

		daftarBuku := db.MongoDB.Collection("BookList")

		_, err = daftarBuku.InsertOne(context.TODO(), data)

		if err != nil {
			http.Error(w, "Failed to add book", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Buku berhasil ditambahkan"))
		return

	case "PUT":
		vars := mux.Vars(r)
		id := vars["id"]

		var updateData struct {
			Stok  uint8  `json:"Stok"`
			Harga uint64 `json:"Harga"`
		}
		err := json.NewDecoder(r.Body).Decode(&updateData)
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

		daftarBuku := db.MongoDB.Collection("BookList")

		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			http.Error(w, "Invalid book ID format", http.StatusBadRequest)
			return
		}

		_, err = daftarBuku.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{
			"$set": bson.M{
				"Stok":  updateData.Stok,
				"Harga": updateData.Harga,
			},
		})
		if err != nil {
			http.Error(w, "Failed to update book", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Buku berhasil diperbarui"))
		return

	case "DELETE":

		vars := mux.Vars(r)
		id := vars["id"]

		db, err := db.DBConnection()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer db.MongoDB.Client().Disconnect(context.TODO())

		daftarBuku := db.MongoDB.Collection("BookList")

		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			http.Error(w, "Invalid book ID format", http.StatusBadRequest)
			return
		}

		_, err = daftarBuku.DeleteOne(context.TODO(), bson.M{"_id": objID})
		if err != nil {
			http.Error(w, "Failed to delete book", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Buku berhasil dihapus"))
		return

	default:
		// Method not allowed
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}
