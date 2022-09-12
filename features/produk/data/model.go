package data

import (
	"project/kutsuya/features/produk"
	userModel "project/kutsuya/features/user/data"

	"gorm.io/gorm"
)

type Produk struct {
	gorm.Model
	Nama_Produk     string
	Ukuran          int
	Merk            string
	Warna           string
	Gender_Pengguna string
	Harga           string
	Desksripsi      string
	Url_Image       string
	User            userModel.User
}

func fromCore(dataCore produk.Core) Produk {
	return Produk{
		Nama_Produk:     dataCore.Nama_Produk,
		Ukuran:          dataCore.Ukuran,
		Merk:            dataCore.Merk,
		Warna:           dataCore.Warna,
		Gender_Pengguna: dataCore.Gender_Pengguna,
		Harga:           dataCore.Harga,
		Desksripsi:      dataCore.Desksripsi,
		Url_Image:       dataCore.Url_Image,
	}

}

func (dataProduk *Produk) toCore() produk.Core {
	return produk.Core{
		ID:              dataProduk.ID,
		Nama_Produk:     dataProduk.Nama_Produk,
		Ukuran:          dataProduk.Ukuran,
		Merk:            dataProduk.Merk,
		Warna:           dataProduk.Warna,
		Gender_Pengguna: dataProduk.Gender_Pengguna,
		Harga:           dataProduk.Harga,
		Desksripsi:      dataProduk.Desksripsi,
		Url_Image:       dataProduk.Url_Image,
	}
}

func toCoreList(dataProduk []Produk) []produk.Core {
	var dataCore []produk.Core

	for key := range dataProduk {
		dataCore = append(dataCore, dataProduk[key].toCore())

	}

	return dataCore

}
