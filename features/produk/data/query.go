package data

import (
	"errors"
	"project/kutsuya/features/produk"

	"gorm.io/gorm"
)

type productData struct {
	db *gorm.DB
}

func New(db *gorm.DB) produk.DataInterface {
	return &productData{
		db: db,
	}

}

func (repo *productData) Select_AllProduk() ([]produk.Core, error) {
	var all_ProdData []Produk
	tx := repo.db.Find(&all_ProdData)

	if tx.Error != nil {
		return nil, tx.Error
	}

	produk_List := toCoreList(all_ProdData)
	return produk_List, nil
}

func (repo *productData) InsertProduk(newProduk produk.Core) (int, error) {
	newUser := fromCore(newProduk)

	tx := repo.db.Create(&newUser)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil

}

func (repo *productData) UpdateDataProduk(dataProduk produk.Core, id_produk int) (int, error) {
	var dataUpdate Produk
	dataProduk.ID = uint(id_produk)

	tx_OldData := repo.db.First(&dataUpdate, dataProduk.ID)

	if tx_OldData.Error != nil {
		return -1, tx_OldData.Error
	}

	if dataProduk.Nama_Produk != "" {
		dataUpdate.Nama_Produk = dataProduk.Nama_Produk
	}

	if dataProduk.Ukuran != 0 {
		dataUpdate.Ukuran = dataProduk.Ukuran
	}

	if dataProduk.Merk != "" {
		dataUpdate.Merk = dataProduk.Merk
	}

	if dataProduk.Warna != "" {
		dataUpdate.Warna = dataProduk.Warna
	}

	if dataProduk.Gender_Pengguna != "" {
		dataUpdate.Gender_Pengguna = dataProduk.Gender_Pengguna
	}

	if dataProduk.Harga != 0 {
		dataUpdate.Harga = dataProduk.Harga
	}

	if dataProduk.Desksripsi != "" {
		dataUpdate.Desksripsi = dataProduk.Desksripsi
	}

	tx_newData := repo.db.Save(&dataUpdate)

	if tx_newData.Error != nil {
		return -1, tx_newData.Error
	}

	if tx_newData.RowsAffected != 1 {
		return 0, errors.New("zero row affected, fail update")
	}

	return int(tx_newData.RowsAffected), nil

}

func (repo *productData) SelectProdukById(id_produk int) (produk.Core, error) {
	var dataProduk Produk
	dataProduk.ID = uint(id_produk)

	tx := repo.db.First(&dataProduk)

	if tx.Error != nil {
		return produk.Core{}, tx.Error
	}

	dataProdukCore := dataProduk.toCore()
	return dataProdukCore, nil

}
