// Cafe-Menu — Aplikasi Katalog Menu Digital Cafe
// Tugas Besar Algoritma Pemrograman 2
//
// Anggota Kelompok:
//   - Angelina
//   - Putri
//   - Syifa

package main

import "fmt"

// ===================== KONSTANTA =====================

const MAKS = 100 // kapasitas maksimum array statis

// ===================== TIPE BENTUKAN =====================

// MenuItem menyimpan semua data setiap item menu cafe
type MenuItem struct {
	ID        int
	Nama      string
	Kategori  string
	Harga     float64
	Komposisi string
	Tersedia  bool
}

// KoleksiMenu adalah tipe bentukan yang menyimpan array statis dan jumlah elemen aktif
type KoleksiMenu struct {
	Data   [MAKS]MenuItem
	Jumlah int
}

// ===================== VARIABEL GLOBAL =====================
// Hanya array utama yang boleh menjadi variabel global

var daftarMenu KoleksiMenu

// ===================== HELPER STRING =====================

// toLower mengubah semua huruf kapital menjadi huruf kecil
// Parameter : s — string yang akan diubah
// Return    : string hasil konversi huruf kecil semua
func toLower(s string) string {
	hasil := ""
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			hasil = hasil + string(c+32)
		} else {
			hasil = hasil + string(c)
		}
	}
	return hasil
}

// equalFold membandingkan dua string tanpa memperhatikan besar kecil huruf
// Parameter : a, b — dua string yang dibandingkan
// Return    : true jika sama (case-insensitive), false jika berbeda
func equalFold(a, b string) bool {
	return toLower(a) == toLower(b)
}

// ulangiKarakter mengulang sebuah karakter sebanyak n kali
// Parameter : char — karakter yang diulang; n — jumlah pengulangan
// Return    : string hasil pengulangan
func ulangiKarakter(char string, n int) string {
	hasil := ""
	i := 0
	for i < n {
		hasil = hasil + char
		i++
	}
	return hasil
}

// containsInsensitive memeriksa apakah sub ada di dalam str (case-insensitive)
// Parameter : str — string sumber; sub — substring yang dicari
// Return    : true jika ditemukan, false jika tidak
func containsInsensitive(str, sub string) bool {
	strL := toLower(str)
	subL := toLower(sub)
	lenStr := len(strL)
	lenSub := len(subL)
	if lenSub == 0 {
		return true
	}
	ditemukan := false
	i := 0
	for i <= lenStr-lenSub {
		cocok := true
		j := 0
		for j < lenSub {
			if strL[i+j] != subL[j] {
				cocok = false
			}
			j++
		}
		if cocok {
			ditemukan = true
		}
		i++
	}
	return ditemukan
}

// ===================== HELPER INPUT =====================

// bacaString membaca satu kata dari input pengguna
// Parameter : prompt — teks yang ditampilkan sebelum input
// Return    : string yang diketik pengguna
func bacaString(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scan(&input)
	return input
}

// bacaInt membaca input integer dari pengguna dengan validasi
// Parameter : prompt — teks yang ditampilkan sebelum input
// Return    : integer valid yang diketik pengguna
func bacaInt(prompt string) int {
	valid := false
	val := 0
	for !valid {
		fmt.Print(prompt)
		_, err := fmt.Scan(&val)
		if err == nil {
			valid = true
		} else {
			fmt.Println("  Input harus berupa angka! Coba lagi.")
			var buang string
			fmt.Scan(&buang)
		}
	}
	return val
}

// bacaFloat membaca input float64 dari pengguna dengan validasi
// Parameter : prompt — teks yang ditampilkan sebelum input
// Return    : float64 valid yang diketik pengguna
func bacaFloat(prompt string) float64 {
	valid := false
	val := 0.0
	for !valid {
		fmt.Print(prompt)
		_, err := fmt.Scan(&val)
		if err == nil {
			valid = true
		} else {
			fmt.Println("  Input harus berupa angka! Coba lagi.")
			var buang string
			fmt.Scan(&buang)
		}
	}
	return val
}

// bacaBool membaca pilihan ya/tidak dari pengguna
// Parameter : prompt — teks pertanyaan yang ditampilkan
// Return    : true jika pengguna memilih 'y', false jika 'n'
func bacaBool(prompt string) bool {
	valid := false
	hasil := false
	for !valid {
		fmt.Print(prompt + " (y/n): ")
		var input string
		fmt.Scan(&input)
		input = toLower(input)
		if input == "y" {
			hasil = true
			valid = true
		} else if input == "n" {
			hasil = false
			valid = true
		} else {
			fmt.Println("  Masukkan 'y' untuk Ya atau 'n' untuk Tidak.")
		}
	}
	return hasil
}

// lanjutkan menunggu pengguna menekan Enter sebelum kembali ke menu utama
func lanjutkan() {
	fmt.Print("\nKetik sembarang lalu Enter untuk kembali ke menu utama: ")
	var buang string
	fmt.Scan(&buang)
}

// ===================== HELPER TAMPILAN =====================

// cetakGaris mencetak garis pemisah sepanjang n karakter
// Parameter : char — karakter pembentuk garis; n — panjang garis
func cetakGaris(char string, n int) {
	fmt.Println(ulangiKarakter(char, n))
}

// cetakHeaderSection menampilkan header bernama judul dengan garis pembatas
// Parameter : judul — teks judul bagian yang ditampilkan
func cetakHeaderSection(judul string) {
	fmt.Println()
	cetakGaris("-", 60)
	fmt.Println("  " + judul)
	cetakGaris("-", 60)
}

// statusTersedia mengubah nilai bool ketersediaan menjadi teks
// Parameter : tersedia — status ketersediaan item
// Return    : "TERSEDIA" atau "HABIS"
func statusTersedia(tersedia bool) string {
	teks := "HABIS"
	if tersedia {
		teks = "TERSEDIA"
	}
	return teks
}

// cetakItemMenu menampilkan detail satu item menu ke layar
// Parameter : m — item menu yang akan ditampilkan
func cetakItemMenu(m MenuItem) {
	fmt.Println()
	cetakGaris("-", 58)
	fmt.Printf("  #%02d  %s\n", m.ID, m.Nama)
	fmt.Printf("  Kategori  : %s\n", m.Kategori)
	fmt.Printf("  Harga     : Rp %.0f\n", m.Harga)
	fmt.Printf("  Komposisi : %s\n", m.Komposisi)
	fmt.Printf("  Status    : %s\n", statusTersedia(m.Tersedia))
}

// cetakDaftarDariArray menampilkan semua item dari array statis sementara
// Parameter : arr — array statis; n — jumlah elemen aktif
func cetakDaftarDariArray(arr [MAKS]MenuItem, n int) {
	if n == 0 {
		fmt.Println()
		fmt.Println("  Tidak ada data menu yang ditemukan.")
		return
	}
	i := 0
	for i < n {
		cetakItemMenu(arr[i])
		i++
	}
	fmt.Println()
	cetakGaris("-", 58)
	fmt.Printf("  Menampilkan %d menu\n", n)
}

// ===================== HELPER NEXT ID =====================

// hitungNextID mencari ID terbesar dari koleksi dan mengembalikan ID berikutnya
// Parameter : k — koleksi menu saat ini
// Return    : integer ID berikutnya yang belum dipakai
func hitungNextID(k KoleksiMenu) int {
	maks := 0
	i := 0
	for i < k.Jumlah {
		if k.Data[i].ID > maks {
			maks = k.Data[i].ID
		}
		i++
	}
	return maks + 1
}

/*
Angelina
*/
// ===================== TAMPILAN MENU UTAMA =====================

// cetakMainMenu menampilkan menu utama aplikasi beserta total menu saat ini
// Parameter : totalMenu — jumlah menu yang tersimpan saat ini
func cetakMainMenu(totalMenu int) {
	fmt.Println()
	cetakGaris("=", 60)
	fmt.Println("  CAFE-MENU")
	fmt.Println("  Aplikasi Katalog Menu Digital Cafe")
	fmt.Println("  Tugas Besar Algoritma Pemrograman 2")
	cetakGaris("=", 60)
	fmt.Println()
	fmt.Println("  [1] Tampilkan Semua Menu")
	fmt.Println("  [2] Tambah Menu Baru")
	fmt.Println("  [3] Ubah Data Menu")
	fmt.Println("  [4] Hapus Menu")
	fmt.Println("  [5] Cari Menu  (Sequential / Binary Search)")
	fmt.Println("  [6] Urutkan Menu (Selection / Insertion Sort)")
	fmt.Println("  [7] Statistik Cafe")
	fmt.Println("  [0] Keluar")
	fmt.Println()
	cetakGaris("-", 60)
	fmt.Printf("  Total menu: %d item (kapasitas maks: %d)\n", totalMenu, MAKS)
}

// ===================== CRUD =====================

// tampilkanSemuaMenu menampilkan seluruh isi daftarMenu global
func tampilkanSemuaMenu() {
	cetakHeaderSection("SEMUA MENU")
	cetakDaftarDariArray(daftarMenu.Data, daftarMenu.Jumlah)
}

// tambahMenu menambahkan satu item menu baru ke dalam daftarMenu global
func tambahMenu() {
	cetakHeaderSection("TAMBAH MENU BARU")

	if daftarMenu.Jumlah >= MAKS {
		fmt.Println("  Menu sudah penuh! Kapasitas maksimum tercapai.")
		return
	}

	nama := bacaString("  Nama Menu            : ")
	if nama == "" {
		fmt.Println("  Nama tidak boleh kosong!")
		return
	}

	fmt.Println()
	fmt.Println("  Contoh kategori: coffee | non-coffee | makanan | dessert | snack")
	kategori := toLower(bacaString("  Kategori             : "))
	harga := bacaFloat("  Harga (Rp)           : ")
	komposisi := bacaString("  Komposisi Bahan      : ")
	tersedia := bacaBool("  Tersedia sekarang    ")

	idBaru := hitungNextID(daftarMenu)

	item := MenuItem{
		ID:        idBaru,
		Nama:      nama,
		Kategori:  kategori,
		Harga:     harga,
		Komposisi: komposisi,
		Tersedia:  tersedia,
	}

	daftarMenu.Data[daftarMenu.Jumlah] = item
	daftarMenu.Jumlah++

	fmt.Println()
	fmt.Printf("  Menu '%s' berhasil ditambahkan! (ID: #%02d)\n", nama, idBaru)
}

// cariIdxByID mencari indeks elemen di array berdasarkan ID menggunakan Sequential Search
// Parameter : k — koleksi menu; id — ID yang dicari
// Return    : indeks elemen jika ditemukan, -1 jika tidak ada
func cariIdxByID(k KoleksiMenu, id int) int {
	idx := -1
	ditemukan := false
	i := 0
	for i < k.Jumlah && !ditemukan {
		if k.Data[i].ID == id {
			idx = i
			ditemukan = true
		}
		i++
	}
	return idx
}

// ubahMenu mengubah data satu item menu berdasarkan ID yang dipilih pengguna
// Pencarian ID menggunakan Sequential Search
func ubahMenu() {
	cetakHeaderSection("UBAH DATA MENU")

	if daftarMenu.Jumlah == 0 {
		fmt.Println("  Belum ada menu!")
		return
	}
	cetakDaftarDariArray(daftarMenu.Data, daftarMenu.Jumlah)

	id := bacaInt("  Masukkan ID menu yang ingin diubah: ")

	// Sequential Search untuk mencari posisi berdasarkan ID
	idx := cariIdxByID(daftarMenu, id)
	if idx == -1 {
		fmt.Println("  ID tidak ditemukan!")
		return
	}

	fmt.Println()
	fmt.Printf("  Mengedit: %s\n", daftarMenu.Data[idx].Nama)
	fmt.Println("  (Masukkan '-' untuk melewati field yang tidak ingin diubah)")
	fmt.Println()

	nama := bacaString(fmt.Sprintf("  Nama Menu [%s]: ", daftarMenu.Data[idx].Nama))
	if nama != "-" && nama != "" {
		daftarMenu.Data[idx].Nama = nama
	}

	kategori := bacaString(fmt.Sprintf("  Kategori [%s]: ", daftarMenu.Data[idx].Kategori))
	if kategori != "-" && kategori != "" {
		daftarMenu.Data[idx].Kategori = toLower(kategori)
	}

	fmt.Printf("  Harga [%.0f] (masukkan 0 untuk melewati): ", daftarMenu.Data[idx].Harga)
	var harga float64
	fmt.Scan(&harga)
	if harga > 0 {
		daftarMenu.Data[idx].Harga = harga
	}

	komposisi := bacaString(fmt.Sprintf("  Komposisi [%s]: ", daftarMenu.Data[idx].Komposisi))
	if komposisi != "-" && komposisi != "" {
		daftarMenu.Data[idx].Komposisi = komposisi
	}

	daftarMenu.Data[idx].Tersedia = bacaBool(fmt.Sprintf("  Tersedia [%v]", daftarMenu.Data[idx].Tersedia))

	fmt.Println()
	fmt.Printf("  Data menu ID #%02d berhasil diperbarui!\n", id)
}

// hapusMenu menghapus satu item menu dari daftarMenu berdasarkan ID
// Pencarian menggunakan Binary Search (data diurutkan sementara berdasarkan ID)
func hapusMenu() {
	cetakHeaderSection("HAPUS MENU")

	if daftarMenu.Jumlah == 0 {
		fmt.Println("  Belum ada menu!")
		return
	}
	cetakDaftarDariArray(daftarMenu.Data, daftarMenu.Jumlah)

	id := bacaInt("  Masukkan ID menu yang ingin dihapus: ")

	// Binary Search berdasarkan ID untuk menemukan posisi
	// Buat salinan array diurutkan berdasarkan ID
	var sortedArr [MAKS]MenuItem
	var sortedN int
	sortedArr, sortedN = salinArray(daftarMenu.Data, daftarMenu.Jumlah)
	sortedArr = insertionSortByID(sortedArr, sortedN)

	posBS := binarySearchByID(sortedArr, sortedN, id)
	if posBS == -1 {
		fmt.Println("  ID tidak ditemukan!")
		return
	}

	namaHapus := sortedArr[posBS].Nama

	konfirmasi := bacaBool(fmt.Sprintf("  Yakin menghapus '%s'", namaHapus))
	if konfirmasi {
		// Cari posisi asli di daftarMenu (sequential) lalu geser elemen
		idxAsli := cariIdxByID(daftarMenu, id)
		hapusPadaIndeks(&daftarMenu, idxAsli)
		fmt.Println()
		fmt.Printf("  Menu '%s' berhasil dihapus.\n", namaHapus)
	} else {
		fmt.Println("  Penghapusan dibatalkan.")
	}
}

// hapusPadaIndeks menghapus elemen pada posisi idx dengan menggeser elemen-elemen sesudahnya
// Parameter : k — pointer ke koleksi menu; idx — indeks elemen yang dihapus
func hapusPadaIndeks(k *KoleksiMenu, idx int) {
	i := idx
	for i < k.Jumlah-1 {
		k.Data[i] = k.Data[i+1]
		i++
	}
	k.Jumlah--
}

// ===================== HELPER ARRAY =====================

// salinArray menyalin n elemen pertama dari src ke array baru
// Parameter : src — array sumber; n — jumlah elemen yang disalin
// Return    : salinan array dan jumlah elemennya
func salinArray(src [MAKS]MenuItem, n int) ([MAKS]MenuItem, int) {
	var hasil [MAKS]MenuItem
	i := 0
	for i < n {
		hasil[i] = src[i]
		i++
	}
	return hasil, n
}

// ===================== PENCARIAN =====================

// sequentialSearchKategori mencari semua item dengan kategori tertentu secara berurutan
// Parameter : k — koleksi menu; kategori — kategori yang dicari
// Return    : array hasil pencarian dan jumlah elemen yang cocok
func sequentialSearchKategori(k KoleksiMenu, kategori string) ([MAKS]MenuItem, int) {
	var hasil [MAKS]MenuItem
	jumlah := 0
	i := 0
	for i < k.Jumlah {
		if equalFold(k.Data[i].Kategori, kategori) {
			hasil[jumlah] = k.Data[i]
			jumlah++
		}
		i++
	}
	return hasil, jumlah
}

// sequentialSearchNama mencari semua item yang namanya mengandung kata kunci
// Parameter : k — koleksi menu; kata — kata kunci pencarian nama
// Return    : array hasil pencarian dan jumlah elemen yang cocok
func sequentialSearchNama(k KoleksiMenu, kata string) ([MAKS]MenuItem, int) {
	var hasil [MAKS]MenuItem
	jumlah := 0
	i := 0
	for i < k.Jumlah {
		if containsInsensitive(k.Data[i].Nama, kata) {
			hasil[jumlah] = k.Data[i]
			jumlah++
		}
		i++
	}
	return hasil, jumlah
}

// insertionSortByID mengurutkan array berdasarkan ID secara ascending (untuk Binary Search)
// Parameter : arr — array yang diurutkan; n — jumlah elemen aktif
// Return    : array yang sudah terurut berdasarkan ID
func insertionSortByID(arr [MAKS]MenuItem, n int) [MAKS]MenuItem {
	i := 1
	for i < n {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j].ID > key.ID {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
		i++
	}
	return arr
}

// insertionSortByKategori mengurutkan array berdasarkan Kategori A-Z (untuk Binary Search kategori)
// Parameter : arr — array yang diurutkan; n — jumlah elemen aktif
// Return    : array yang sudah terurut berdasarkan Kategori
func insertionSortByKategori(arr [MAKS]MenuItem, n int) [MAKS]MenuItem {
	i := 1
	for i < n {
		key := arr[i]
		j := i - 1
		for j >= 0 && toLower(arr[j].Kategori) > toLower(key.Kategori) {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
		i++
	}
	return arr
}

// binarySearchByID mencari satu elemen berdasarkan ID menggunakan Binary Search
// Parameter : arr — array terurut berdasarkan ID; n — jumlah elemen; id — ID yang dicari
// Return    : indeks elemen dalam arr jika ditemukan, -1 jika tidak ada
func binarySearchByID(arr [MAKS]MenuItem, n int, id int) int {
	low := 0
	high := n - 1
	foundIdx := -1
	ditemukan := false
	for low <= high && !ditemukan {
		mid := (low + high) / 2
		if arr[mid].ID == id {
			foundIdx = mid
			ditemukan = true
		} else if arr[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return foundIdx
}

// binarySearchKategori mencari semua item dengan kategori tertentu menggunakan Binary Search
// Parameter : k — koleksi menu; kategori — kategori yang dicari
// Return    : array hasil pencarian dan jumlah elemen yang cocok
func binarySearchKategori(k KoleksiMenu, kategori string) ([MAKS]MenuItem, int) {
	// Buat salinan dan urutkan berdasarkan kategori terlebih dahulu
	var sorted [MAKS]MenuItem
	var n int
	sorted, n = salinArray(k.Data, k.Jumlah)
	sorted = insertionSortByKategori(sorted, n)

	// Cari satu posisi dengan binary search
	low := 0
	high := n - 1
	foundIdx := -1
	ditemukan := false
	for low <= high && !ditemukan {
		mid := (low + high) / 2
		katMid := toLower(sorted[mid].Kategori)
		katCari := toLower(kategori)
		if katMid == katCari {
			foundIdx = mid
			ditemukan = true
		} else if katMid < katCari {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if foundIdx == -1 {
		var kosong [MAKS]MenuItem
		return kosong, 0
	}

	// Kumpulkan semua elemen dengan kategori sama di sekitar foundIdx
	var hasil [MAKS]MenuItem
	jumlah := 0

	left := foundIdx
	for left > 0 && equalFold(sorted[left-1].Kategori, kategori) {
		left--
	}
	right := foundIdx
	for right < n-1 && equalFold(sorted[right+1].Kategori, kategori) {
		right++
	}
	i := left
	for i <= right {
		hasil[jumlah] = sorted[i]
		jumlah++
		i++
	}
	return hasil, jumlah
}

// menuCari menampilkan sub-menu pencarian dan menjalankan algoritma yang dipilih pengguna
func menuCari() {
	cetakHeaderSection("CARI MENU")
	fmt.Println()
	fmt.Println("  Cari berdasarkan:")
	fmt.Println("  [1] Kategori  (Sequential Search)")
	fmt.Println("  [2] Kategori  (Binary Search)")
	fmt.Println("  [3] Nama Menu (Sequential Search)")
	fmt.Println()

	pilihan := bacaInt("  Pilih metode (1/2/3): ")

	var hasilArr [MAKS]MenuItem
	var jumlahHasil int
	var metode, kata string

	if pilihan == 1 {
		fmt.Println()
		fmt.Println("  Contoh: coffee, non-coffee, makanan, dessert, snack")
		kata = toLower(bacaString("  Masukkan kategori: "))
		hasilArr, jumlahHasil = sequentialSearchKategori(daftarMenu, kata)
		metode = "Sequential Search (Kategori)"
	} else if pilihan == 2 {
		fmt.Println()
		fmt.Println("  Contoh: coffee, non-coffee, makanan, dessert, snack")
		kata = toLower(bacaString("  Masukkan kategori: "))
		hasilArr, jumlahHasil = binarySearchKategori(daftarMenu, kata)
		metode = "Binary Search (Kategori)"
	} else if pilihan == 3 {
		fmt.Println()
		kata = bacaString("  Masukkan kata kunci nama menu: ")
		hasilArr, jumlahHasil = sequentialSearchNama(daftarMenu, kata)
		metode = "Sequential Search (Nama)"
	} else {
		fmt.Println("  Pilihan tidak valid!")
		return
	}

	fmt.Println()
	fmt.Printf("  Hasil %s — '%s' (%d item)\n", metode, kata, jumlahHasil)
	cetakDaftarDariArray(hasilArr, jumlahHasil)
}

// ===================== PENGURUTAN =====================

// selectionSortHarga mengurutkan array berdasarkan harga
// Parameter : arr — array yang diurutkan; n — jumlah elemen; asc — true=ascending, false=descending
// Return    : array yang sudah terurut
func selectionSortHarga(arr [MAKS]MenuItem, n int, asc bool) [MAKS]MenuItem {
	i := 0
	for i < n-1 {
		pilihanIdx := i
		j := i + 1
		for j < n {
			if asc {
				if arr[j].Harga < arr[pilihanIdx].Harga {
					pilihanIdx = j
				}
			} else {
				if arr[j].Harga > arr[pilihanIdx].Harga {
					pilihanIdx = j
				}
			}
			j++
		}
		arr[i], arr[pilihanIdx] = arr[pilihanIdx], arr[i]
		i++
	}
	return arr
}

// insertionSortHarga mengurutkan array berdasarkan harga dengan metode penyisipan
// Parameter : arr — array yang diurutkan; n — jumlah elemen; asc — true=ascending, false=descending
// Return    : array yang sudah terurut
func insertionSortHarga(arr [MAKS]MenuItem, n int, asc bool) [MAKS]MenuItem {
	i := 1
	for i < n {
		key := arr[i]
		j := i - 1
		if asc {
			for j >= 0 && arr[j].Harga > key.Harga {
				arr[j+1] = arr[j]
				j--
			}
		} else {
			for j >= 0 && arr[j].Harga < key.Harga {
				arr[j+1] = arr[j]
				j--
			}
		}
		arr[j+1] = key
		i++
	}
	return arr
}

// selectionSortNama mengurutkan array berdasarkan nama secara alfabetis
// Parameter : arr — array yang diurutkan; n — jumlah elemen; asc — true=A-Z, false=Z-A
// Return    : array yang sudah terurut
func selectionSortNama(arr [MAKS]MenuItem, n int, asc bool) [MAKS]MenuItem {
	i := 0
	for i < n-1 {
		pilihanIdx := i
		j := i + 1
		for j < n {
			if asc {
				if toLower(arr[j].Nama) < toLower(arr[pilihanIdx].Nama) {
					pilihanIdx = j
				}
			} else {
				if toLower(arr[j].Nama) > toLower(arr[pilihanIdx].Nama) {
					pilihanIdx = j
				}
			}
			j++
		}
		arr[i], arr[pilihanIdx] = arr[pilihanIdx], arr[i]
		i++
	}
	return arr
}

// insertionSortNama mengurutkan array berdasarkan nama dengan metode penyisipan
// Parameter : arr — array yang diurutkan; n — jumlah elemen; asc — true=A-Z, false=Z-A
// Return    : array yang sudah terurut
func insertionSortNama(arr [MAKS]MenuItem, n int, asc bool) [MAKS]MenuItem {
	i := 1
	for i < n {
		key := arr[i]
		j := i - 1
		if asc {
			for j >= 0 && toLower(arr[j].Nama) > toLower(key.Nama) {
				arr[j+1] = arr[j]
				j--
			}
		} else {
			for j >= 0 && toLower(arr[j].Nama) < toLower(key.Nama) {
				arr[j+1] = arr[j]
				j--
			}
		}
		arr[j+1] = key
		i++
	}
	return arr
}

// menuUrut menampilkan sub-menu pengurutan dan menjalankan algoritma yang dipilih
func menuUrut() {
	cetakHeaderSection("URUTKAN MENU")
	fmt.Println()
	fmt.Println("  Urutkan berdasarkan:")
	fmt.Println("  [1] Harga")
	fmt.Println("  [2] Nama (Alfabetis)")
	fmt.Println()
	pilihanKolom := bacaInt("  Pilih kolom (1/2): ")

	if pilihanKolom != 1 && pilihanKolom != 2 {
		fmt.Println("  Pilihan tidak valid!")
		return
	}

	fmt.Println()
	fmt.Println("  Pilih algoritma:")
	fmt.Println("  [1] Selection Sort  — Pilih ekstrem, pindahkan ke depan")
	fmt.Println("  [2] Insertion Sort  — Sisipkan ke posisi yang tepat")
	fmt.Println()
	pilihanAlgo := bacaInt("  Pilih algoritma (1/2): ")

	if pilihanAlgo != 1 && pilihanAlgo != 2 {
		fmt.Println("  Pilihan tidak valid!")
		return
	}

	fmt.Println()
	fmt.Println("  Pilih arah urutan:")
	fmt.Println("  [1] Ascending  (terkecil ke terbesar / A ke Z)")
	fmt.Println("  [2] Descending (terbesar ke terkecil / Z ke A)")
	fmt.Println()
	pilihanArah := bacaInt("  Pilih arah (1/2): ")

	if pilihanArah != 1 && pilihanArah != 2 {
		fmt.Println("  Pilihan tidak valid!")
		return
	}

	asc := pilihanArah == 1

	// Buat salinan agar urutan asli tidak berubah
	var salinan [MAKS]MenuItem
	var n int
	salinan, n = salinArray(daftarMenu.Data, daftarMenu.Jumlah)

	var namaAlgo, namaKolom, namaArah string

	if pilihanAlgo == 1 {
		namaAlgo = "Selection Sort"
		if pilihanKolom == 1 {
			namaKolom = "Harga"
			salinan = selectionSortHarga(salinan, n, asc)
		} else {
			namaKolom = "Nama"
			salinan = selectionSortNama(salinan, n, asc)
		}
	} else {
		namaAlgo = "Insertion Sort"
		if pilihanKolom == 1 {
			namaKolom = "Harga"
			salinan = insertionSortHarga(salinan, n, asc)
		} else {
			namaKolom = "Nama"
			salinan = insertionSortNama(salinan, n, asc)
		}
	}

	if asc {
		namaArah = "Ascending"
	} else {
		namaArah = "Descending"
	}

	fmt.Printf("  %s — %s — %s\n", namaAlgo, namaKolom, namaArah)
	cetakDaftarDariArray(salinan, n)
}

// ===================== STATISTIK =====================

// tampilkanStatistik menampilkan ringkasan data seluruh menu cafe
func tampilkanStatistik() {
	fmt.Println()
	cetakGaris("=", 60)
	fmt.Println("  CAFE-MENU  —  STATISTIK LENGKAP")
	cetakGaris("=", 60)

	if daftarMenu.Jumlah == 0 {
		fmt.Println("  Belum ada data menu.")
		return
	}

	// Kumpulkan data kategori unik (maks 20 kategori berbeda)
	const MAKS_KAT = 20
	var kategoriList [MAKS_KAT]string
	var kategoriCount [MAKS_KAT]int
	var kategoriTotal [MAKS_KAT]float64
	jumlahKat := 0

	var totalHarga float64
	totalTersedia := 0

	i := 0
	for i < daftarMenu.Jumlah {
		m := daftarMenu.Data[i]
		totalHarga = totalHarga + m.Harga
		if m.Tersedia {
			totalTersedia++
		}

		// Cari kategori ini sudah ada atau belum
		idxKat := -1
		k := 0
		ditemukanKat := false
		for k < jumlahKat && !ditemukanKat {
			if equalFold(kategoriList[k], m.Kategori) {
				idxKat = k
				ditemukanKat = true
			}
			k++
		}

		if idxKat == -1 {
			// Kategori baru
			kategoriList[jumlahKat] = m.Kategori
			kategoriCount[jumlahKat] = 1
			kategoriTotal[jumlahKat] = m.Harga
			jumlahKat++
		} else {
			kategoriCount[idxKat] = kategoriCount[idxKat] + 1
			kategoriTotal[idxKat] = kategoriTotal[idxKat] + m.Harga
		}
		i++
	}

	rataRata := totalHarga / float64(daftarMenu.Jumlah)
	totalTidakTersedia := daftarMenu.Jumlah - totalTersedia

	fmt.Println()
	fmt.Printf("  Total Semua Menu  : %d item\n", daftarMenu.Jumlah)
	fmt.Printf("  Kapasitas Tersisa : %d slot\n", MAKS-daftarMenu.Jumlah)
	fmt.Printf("  Rata-rata Harga   : Rp %.0f\n", rataRata)
	fmt.Printf("  Menu Tersedia     : %d item\n", totalTersedia)
	fmt.Printf("  Menu Habis        : %d item\n", totalTidakTersedia)

	fmt.Println()
	fmt.Println("  Detail Per Kategori:")
	cetakGaris("-", 58)
	fmt.Printf("  %-20s  %5s  %s\n", "Kategori", "Total", "Rata-rata Harga")
	cetakGaris("-", 58)

	idx := 0
	for idx < jumlahKat {
		rataKat := kategoriTotal[idx] / float64(kategoriCount[idx])
		fmt.Printf("  %-20s  %5d  Rp %.0f\n", kategoriList[idx], kategoriCount[idx], rataKat)
		idx++
	}
	cetakGaris("-", 58)

	fmt.Println()
	fmt.Println("  Proporsi Menu (Bar Chart):")
	fmt.Println()
	idx = 0
	for idx < jumlahKat {
		persen := float64(kategoriCount[idx]) / float64(daftarMenu.Jumlah) * 100
		barLen := int(persen / 3)
		if barLen < 1 {
			barLen = 1
		}
		bar := ulangiKarakter("#", barLen)
		sisa := ulangiKarakter(".", 33-barLen)
		fmt.Printf("  %-14s [%s%s] %.1f%%\n", kategoriList[idx], bar, sisa, persen)
		idx++
	}
	fmt.Println()
	cetakGaris("=", 60)
}

// ===================== DATA CONTOH =====================

// isiDataContoh mengisi daftarMenu global dengan 10 data awal yang realistis
func isiDataContoh() {
	data := [10]MenuItem{
		{ID: 1, Nama: "Espresso", Kategori: "coffee", Harga: 18000, Komposisi: "espresso-shot,air-panas", Tersedia: true},
		{ID: 2, Nama: "Cappuccino", Kategori: "coffee", Harga: 28000, Komposisi: "espresso,susu-steamed,foam", Tersedia: true},
		{ID: 3, Nama: "Caramel-Latte", Kategori: "coffee", Harga: 35000, Komposisi: "espresso,susu,karamel,whip-cream", Tersedia: true},
		{ID: 4, Nama: "Cold-Brew", Kategori: "coffee", Harga: 30000, Komposisi: "kopi-arabika-cold-brew-12-jam", Tersedia: false},
		{ID: 5, Nama: "Matcha-Latte", Kategori: "non-coffee", Harga: 32000, Komposisi: "matcha-powder,susu-full-cream,gula-tebu", Tersedia: true},
		{ID: 6, Nama: "Teh-Tarik", Kategori: "non-coffee", Harga: 15000, Komposisi: "teh-hitam,susu-kental-manis", Tersedia: true},
		{ID: 7, Nama: "Es-Cokelat", Kategori: "non-coffee", Harga: 22000, Komposisi: "cokelat-premium,susu,es-batu", Tersedia: true},
		{ID: 8, Nama: "Croissant", Kategori: "makanan", Harga: 22000, Komposisi: "tepung-premium,mentega,telur", Tersedia: true},
		{ID: 9, Nama: "Sandwich-Ayam", Kategori: "makanan", Harga: 38000, Komposisi: "roti-sourdough,ayam-panggang,selada,tomat,mayo", Tersedia: false},
		{ID: 10, Nama: "Cheesecake-Matcha", Kategori: "dessert", Harga: 35000, Komposisi: "cream-cheese,matcha,graham-cracker", Tersedia: true},
	}

	i := 0
	for i < 10 {
		daftarMenu.Data[i] = data[i]
		i++
	}
	daftarMenu.Jumlah = 10
}

// ===================== MAIN =====================

func main() {
	isiDataContoh()

	fmt.Println()
	cetakGaris("=", 60)
	fmt.Println("  Selamat datang di Cafe-Menu!")
	fmt.Println("  Aplikasi Katalog Menu Digital Cafe")
	fmt.Println("  Tugas Besar Algoritma Pemrograman 2")
	fmt.Println()
	fmt.Println("  Kelompok:")
	fmt.Println("    - Angelina")
	fmt.Println("    - Putri")
	fmt.Println("    - Syifa")
	cetakGaris("=", 60)

	jalan := true
	for jalan {
		cetakMainMenu(daftarMenu.Jumlah)
		pilihan := bacaInt("  Pilihan Anda: ")

		switch pilihan {
		case 1:
			tampilkanSemuaMenu()
		case 2:
			tambahMenu()
		case 3:
			ubahMenu()
		case 4:
			hapusMenu()
		case 5:
			menuCari()
		case 6:
			menuUrut()
		case 7:
			tampilkanStatistik()
		case 0:
			fmt.Println()
			cetakGaris("=", 60)
			fmt.Println("  Terima kasih telah menggunakan Cafe-Menu!")
			fmt.Println("  Sampai jumpa!")
			cetakGaris("=", 60)
			fmt.Println()
			jalan = false
		default:
			fmt.Println()
			fmt.Println("  Pilihan tidak valid! Silakan pilih angka 0-7.")
		}

		if jalan {
			lanjutkan()
		}
	}
}
