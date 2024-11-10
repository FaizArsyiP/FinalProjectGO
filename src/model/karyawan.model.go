package model

type Karyawan struct {
	ID           string `bson:"_id,omitempty"`
	Nama         string `bson:"Nama"`
	NIK          uint64 `bson:"NIK"`
	Pendidikan   string `bson:"Pendidikan"`
	TanggalMasuk string `bson:"TanggalMasuk"`
	StatusKerja  string `bson:"StatusKerja"`
}
