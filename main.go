package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Item struct {
	id    int
	name  string
	price int
	stock int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_store")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return db, nil
}

// Menampilkan data barang dengan parameter status isSale
func showItem(isSale int) {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select id, name, stock, price from items where isSale = ?", isSale)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Item
	for rows.Next() {
		var item = Item{}
		var err = rows.Scan(&item.id, &item.name, &item.stock, &item.price)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, item)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("ID\tNama Barang\t\tStok\tHarga")
	for _, item := range result {
		fmt.Printf("%d\t%s\t\t%d\tRp. %d\n", item.id, item.name, item.stock, item.price)
	}

}

// Mencari data barang menggunakan keyword id & nama barang
func searchItem() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var key string
	fmt.Print("Masukkan kata kunci ID atau nama barang : ")
	fmt.Scan(&key)

	newKey := "%" + key + "%"
	rows, err := db.Query("select id, name, stock, price from items where id like ? or name like ?", newKey, newKey)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Item
	for rows.Next() {
		var item = Item{}
		var err = rows.Scan(&item.id, &item.name, &item.stock, &item.price)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, item)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, item := range result {
		fmt.Printf("%d\t%s\t\t%d\tRp. %d\n", item.id, item.name, item.stock, item.price)
	}
}

// Menambahkan data barang baru
func insertItem() {
	var nama string
	var stok, harga int
	fmt.Println("Tambah Item Barang :")
	fmt.Print("Nama Barang :")
	fmt.Scan(&nama)
	fmt.Print("Stok Barang :")
	fmt.Scan(&stok)
	fmt.Print("Harga Barang :")
	fmt.Scan(&harga)

	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var isSale = 1
	_, err = db.Exec("insert into items values (?,?,?,?,?)", nil, nama, stok, harga, isSale)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Tambah Barang Sukses!")
}

// Update data barang berdasarkan id barang
func updateItem() {
	var nama, inputId string
	var stok, harga int
	showItem(1)
	for {
		fmt.Print("Masukkan ID Barang : ")
		fmt.Scan(&inputId)
		inputID, errC := strconv.Atoi(inputId)
		if errC != nil {
			fmt.Println("Input tidak valid!")
			continue
		}

		fmt.Print("Nama Barang :")
		fmt.Scan(&nama)
		fmt.Print("Stok Barang :")
		fmt.Scan(&stok)
		fmt.Print("Harga Barang :")
		fmt.Scan(&harga)

		var db, err = connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		_, err = db.Exec("update items set name = ?, stock = ?, price = ? where id = ?", nama, stok, harga, inputID)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Update Barang Sukses!")
		break
	}
}

// Arsip data barang berdasarkan id barang
func disableItem() {
	var inputId string
	showItem(1)
	for {
		fmt.Print("Masukkan ID Barang : ")
		fmt.Scan(&inputId)
		inputID, errC := strconv.Atoi(inputId)
		if errC != nil {
			fmt.Println("Input tidak valid!")
			continue
		}

		var db, err = connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		isSale := 2
		_, err = db.Exec("update items set isSale = ? where id = ?", isSale, inputID)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Arsip Barang Sukses!")
		break
	}
}

// Menampilkan data barang berdasarkan id barang
func enableItem() {
	var inputId string
	showItem(2)
	for {
		fmt.Print("Masukkan ID Barang : ")
		fmt.Scan(&inputId)
		inputID, errC := strconv.Atoi(inputId)
		if errC != nil {
			fmt.Println("Input tidak valid!")
			continue
		}

		var db, err = connect()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		isSale := 1
		_, err = db.Exec("update items set isSale = ? where id = ?", isSale, inputID)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Tampilkan Barang Sukses!")
		break
	}
}

func main() {
	for {
		fmt.Println("Sistem Informasi Barang 'Toko Arf'")
		fmt.Println("1. Lihat Daftar Barang")
		fmt.Println("2. Cari Barang")
		fmt.Println("3. Tambah Barang")
		fmt.Println("4. Update Barang")
		fmt.Println("5. Arsipkan Barang")
		fmt.Println("6. Tampilkan Barang")
		fmt.Println("0. Exit")
		fmt.Print("Pilih (0-5) : ")
		var selectMenu string
		fmt.Scan(&selectMenu)
		selectedMenu, err := strconv.Atoi(selectMenu)
		if err != nil {
			fmt.Println("Input tidak valid!")
			continue
		} else if selectedMenu == 0 {
			break
		} else if selectedMenu == 1 {
			fmt.Println("======================================================")
			showItem(1)
			fmt.Println("======================================================")
		} else if selectedMenu == 2 {
			fmt.Println("======================================================")
			searchItem()
			fmt.Println("======================================================")
		} else if selectedMenu == 3 {
			fmt.Println("======================================================")
			insertItem()
			fmt.Println("======================================================")
		} else if selectedMenu == 4 {
			fmt.Println("======================================================")
			updateItem()
			fmt.Println("======================================================")
		} else if selectedMenu == 5 {
			fmt.Println("======================================================")
			disableItem()
			fmt.Println("======================================================")
		} else if selectedMenu == 6 {
			fmt.Println("======================================================")
			enableItem()
			fmt.Println("======================================================")
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
