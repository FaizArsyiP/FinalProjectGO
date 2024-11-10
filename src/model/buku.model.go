package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Buku struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Judul   string             `bson:"Judul"`
	Penulis string             `bson:"Penulis"`
	Tahun   uint16             `bson:"Tahun"`
	Stok    uint8              `bson:"Stok"`
	Harga   uint64             `bson:"Harga"`
}
