// Cafe-Menu — Aplikasi Katalog Menu Digital Cafe
// Tugas Besar Algoritma Pemrograman 2

package main

import "fmt"

// ===================== STRUCT =====================

// MenuItem menyimpan semua data setiap item menu cafe
type MenuItem struct {
	ID        int
	Nama      string
	Kategori  string
	Harga     float64
	Komposisi string
	Tersedia  bool
}

// ===================== VARIABEL GLOBAL =====================

var daftarMenu []MenuItem
var nextID = 1

// ===================== HELPER STRING =====================

// toLower mengubah semua huruf kapital menjadi huruf kecil
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
func equalFold(a, b string) bool {
	return toLower(a) == toLower(b)
}

// ulangiKarakter mengulang karakter sebanyak n kali dan mengembalikan hasilnya
func ulangiKarakter(char string, n int) string {
	hasil := ""
	for i := 0; i < n; i++ {
		hasil = hasil + char
	}
	return hasil
}

// ===================== HELPER INPUT =====================

// bacaString membaca satu kata dari input pengguna
func bacaString(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scan(&input)
	return input
}

// bacaInt membaca input integer dari pengguna dengan validasi
func bacaInt(prompt string) int {
	for {
		fmt.Print(prompt)
		var val int
		_, err := fmt.Scan(&val)
		if err == nil {
			return val
		}
		fmt.Println("  Input harus berupa angka! Coba lagi.")
		var buang string
		fmt.Scan(&buang)
	}
}

// bacaFloat membaca input float64 dari pengguna dengan validasi
func bacaFloat(prompt string) float64 {
	for {
		fmt.Print(prompt)
		var val float64
		_, err := fmt.Scan(&val)
		if err == nil {
			return val
		}
		fmt.Println("  Input harus berupa angka! Coba lagi.")
		var buang string
		fmt.Scan(&buang)
	}
}

// bacaBool membaca pilihan ya/tidak dari pengguna
func bacaBool(prompt string) bool {
	for {
		fmt.Print(prompt + " (y/n): ")
		var input string
		fmt.Scan(&input)
		input = toLower(input)
		if input == "y" {
			return true
		} else if input == "n" {
			return false
		}
		fmt.Println("  Masukkan 'y' untuk Ya atau 'n' untuk Tidak.")
	}
}

// lanjutkan menunggu pengguna mengetik sesuatu sebelum kembali ke menu utama
func lanjutkan() {
	fmt.Print("\nKetik sembarang lalu Enter untuk kembali ke menu utama: ")
	var buang string
	fmt.Scan(&buang)
}

// ===================== TAMPILAN =====================

// cetakGaris mencetak garis pemisah sepanjang n karakter
func cetakGaris(char string, n int) {
	fmt.Println(ulangiKarakter(char, n))
}

// cetakHeaderSection menampilkan header untuk setiap bagian menu
func cetakHeaderSection(judul string) {
	fmt.Println()
	cetakGaris("-", 60)
	fmt.Println("  " + judul)
	cetakGaris("-", 60)
}

// cetakMainMenu menampilkan menu utama aplikasi
func cetakMainMenu() {
	fmt.Println()
	cetakGaris("=", 60)
	fmt.Println("  CAFE-MENU")
	fmt.Println("  Aplikasi Katalog Menu Digital Cafe")
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
	fmt.Printf("  Total menu: %d item\n", len(daftarMenu))
}

// ===================== TAMPILAN ITEM MENU =====================

// statusTersedia mengembalikan string status ketersediaan menu
func statusTersedia(tersedia bool) string {
	if tersedia {
		return "TERSEDIA"
	}
	return "HABIS"
}

// cetakItemMenu menampilkan satu item menu dalam format detail
func cetakItemMenu(m MenuItem) {
	fmt.Println()
	cetakGaris("-", 58)
	fmt.Printf("  #%02d  %s\n", m.ID, m.Nama)
	fmt.Printf("  Kategori  : %s\n", m.Kategori)
	fmt.Printf("  Harga     : Rp %.0f\n", m.Harga)
	fmt.Printf("  Komposisi : %s\n", m.Komposisi)
	fmt.Printf("  Status    : %s\n", statusTersedia(m.Tersedia))
}

// cetakDaftarMenu menampilkan semua menu dalam slice secara berurutan
func cetakDaftarMenu(list []MenuItem) {
	if len(list) == 0 {
		fmt.Println()
		fmt.Println("  Tidak ada data menu yang ditemukan.")
		return
	}
	for _, m := range list {
		cetakItemMenu(m)
	}
	fmt.Println()
	cetakGaris("-", 58)
	fmt.Printf("  Menampilkan %d menu\n", len(list))
}

// ===================== CRUD MENU =====================

// tambahMenu menambahkan menu baru ke dalam daftarMenu
func tambahMenu() {
	cetakHeaderSection("TAMBAH MENU BARU")

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

	item := MenuItem{
		ID:        nextID,
		Nama:      nama,
		Kategori:  kategori,
		Harga:     harga,
		Komposisi: komposisi,
		Tersedia:  tersedia,
	}
	daftarMenu = append(daftarMenu, item)
	nextID++

	fmt.Println()
	fmt.Printf("  Menu '%s' berhasil ditambahkan! (ID: #%02d)\n", nama, item.ID)
}

// ubahMenu mengubah data menu yang sudah ada berdasarkan ID yang dipilih
func ubahMenu() {
	cetakHeaderSection("UBAH DATA MENU")

	if len(daftarMenu) == 0 {
		fmt.Println("  Belum ada menu!")
		return
	}
	cetakDaftarMenu(daftarMenu)

	id := bacaInt("  Masukkan ID menu yang ingin diubah: ")

	// mencari posisi menu di dalam slice berdasarkan ID
	idx := -1
	for i, m := range daftarMenu {
		if m.ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println("  ID tidak ditemukan!")
		return
	}

	m := &daftarMenu[idx]
	fmt.Println()
	fmt.Printf("  Mengedit: %s\n", m.Nama)
	fmt.Println("  (Masukkan '-' untuk melewati field yang tidak ingin diubah)")
	fmt.Println()

	// memperbarui setiap field jika pengguna tidak mengetik '-'
	nama := bacaString(fmt.Sprintf("  Nama Menu [%s]: ", m.Nama))
	if nama != "-" {
		m.Nama = nama
	}

	kategori := bacaString(fmt.Sprintf("  Kategori [%s]: ", m.Kategori))
	if kategori != "-" {
		m.Kategori = toLower(kategori)
	}

	fmt.Printf("  Harga [%.0f] (masukkan 0 untuk melewati): ", m.Harga)
	var harga float64
	fmt.Scan(&harga)
	if harga > 0 {
		m.Harga = harga
	}

	komposisi := bacaString(fmt.Sprintf("  Komposisi [%s]: ", m.Komposisi))
	if komposisi != "-" {
		m.Komposisi = komposisi
	}

	m.Tersedia = bacaBool(fmt.Sprintf("  Tersedia [%v]", m.Tersedia))

	fmt.Println()
	fmt.Printf("  Data menu ID #%02d berhasil diperbarui!\n", id)
}

// hapusMenu menghapus menu dari daftarMenu berdasarkan ID
func hapusMenu() {
	cetakHeaderSection("HAPUS MENU")

	if len(daftarMenu) == 0 {
		fmt.Println("  Belum ada menu!")
		return
	}
	cetakDaftarMenu(daftarMenu)

	id := bacaInt("  Masukkan ID menu yang ingin dihapus: ")

	// membangun slice baru tanpa elemen yang akan dihapus
	found := false
	var namaHapus string
	newList := []MenuItem{}
	for _, m := range daftarMenu {
		if m.ID == id {
			found = true
			namaHapus = m.Nama
		} else {
			newList = append(newList, m)
		}
	}

	if !found {
		fmt.Println("  ID tidak ditemukan!")
		return
	}

	konfirmasi := bacaBool(fmt.Sprintf("  Yakin menghapus '%s'", namaHapus))
	if konfirmasi {
		daftarMenu = newList
		fmt.Println()
		fmt.Printf("  Menu '%s' berhasil dihapus.\n", namaHapus)
	} else {
		fmt.Println("  Penghapusan dibatalkan.")
	}
}

// tampilkanSemuaMenu menampilkan seluruh isi daftarMenu
func tampilkanSemuaMenu() {
	cetakHeaderSection("SEMUA MENU")
	cetakDaftarMenu(daftarMenu)
}

// ===================== PENCARIAN =====================

// sequentialSearch mencari menu berdasarkan kategori dari awal hingga akhir
func sequentialSearch(kategori string) []MenuItem {
	hasil := []MenuItem{}
	// melakukan perulangan satu per satu dari indeks pertama hingga terakhir
	for _, m := range daftarMenu {
		if equalFold(m.Kategori, kategori) {
			hasil = append(hasil, m)
		}
	}
	return hasil
}

// binarySearch mencari menu berdasarkan kategori menggunakan algoritma biner
// slice diurutkan sementara sebelum pencarian dilakukan
func binarySearch(kategori string) []MenuItem {
	// membuat salinan slice dan mengurutkannya berdasarkan kategori
	sorted := make([]MenuItem, len(daftarMenu))
	copy(sorted, daftarMenu)

	// mengurutkan salinan secara abjad berdasarkan kategori menggunakan insertion sort
	for i := 1; i < len(sorted); i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 && toLower(sorted[j].Kategori) > toLower(key.Kategori) {
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = key
	}

	// melakukan binary search untuk menemukan indeks yang cocok
	low, high, foundIdx := 0, len(sorted)-1, -1
	for low <= high {
		mid := (low + high) / 2
		katMid := toLower(sorted[mid].Kategori)
		katCari := toLower(kategori)
		if katMid == katCari {
			foundIdx = mid
			break
		} else if katMid < katCari {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if foundIdx == -1 {
		return []MenuItem{}
	}

	// mengumpulkan seluruh elemen dengan kategori sama di sekitar foundIdx
	hasil := []MenuItem{}
	left := foundIdx
	for left > 0 && equalFold(sorted[left-1].Kategori, kategori) {
		left--
	}
	right := foundIdx
	for right < len(sorted)-1 && equalFold(sorted[right+1].Kategori, kategori) {
		right++
	}
	for i := left; i <= right; i++ {
		hasil = append(hasil, sorted[i])
	}
	return hasil
}

// menuCari menampilkan UI sub-menu pencarian
func menuCari() {
	cetakHeaderSection("CARI MENU")
	fmt.Println()
	fmt.Println("  Pilih Algoritma Pencarian:")
	fmt.Println("  [1] Sequential Search  -- Menelusuri satu per satu dari awal")
	fmt.Println("  [2] Binary Search      -- Membagi dua data yang sudah diurutkan")
	fmt.Println()

	pilihan := bacaInt("  Pilih metode (1/2): ")
	fmt.Println()
	fmt.Println("  Contoh: coffee, non-coffee, makanan, dessert, snack")
	kategori := toLower(bacaString("  Masukkan kategori yang dicari: "))

	var hasil []MenuItem
	var metode string

	// memilih dan menjalankan algoritma pencarian sesuai input
	if pilihan == 1 {
		hasil = sequentialSearch(kategori)
		metode = "Sequential Search"
	} else if pilihan == 2 {
		hasil = binarySearch(kategori)
		metode = "Binary Search"
	} else {
		fmt.Println("  Pilihan tidak valid!")
		return
	}

	fmt.Printf("  Hasil %s -- kategori '%s' (%d item)\n", metode, kategori, len(hasil))
	cetakDaftarMenu(hasil)
}

// ===================== PENGURUTAN =====================

// selectionSort mengurutkan slice menu berdasarkan harga dari terkecil ke terbesar
func selectionSort(list []MenuItem) []MenuItem {
	n := len(list)
	// melakukan perulangan sebanyak n-1 kali untuk selection sort
	for i := 0; i < n-1; i++ {
		minIdx := i
		// mencari elemen dengan harga minimum dari posisi i+1 hingga akhir
		for j := i + 1; j < n; j++ {
			if list[j].Harga < list[minIdx].Harga {
				minIdx = j
			}
		}
		// menukar elemen minimum ke posisi i
		list[i], list[minIdx] = list[minIdx], list[i]
	}
	return list
}

// insertionSort mengurutkan slice menu berdasarkan harga dengan metode penyisipan
func insertionSort(list []MenuItem) []MenuItem {
	n := len(list)
	// melakukan perulangan mulai dari indeks ke-1 untuk insertion sort
	for i := 1; i < n; i++ {
		key := list[i]
		j := i - 1
		// menggeser elemen yang lebih besar dari key ke posisi berikutnya
		for j >= 0 && list[j].Harga > key.Harga {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = key
	}
	return list
}

// menuUrut menampilkan UI sub-menu pengurutan harga
func menuUrut() {
	cetakHeaderSection("URUTKAN MENU BERDASARKAN HARGA")
	fmt.Println()
	fmt.Println("  Pilih Algoritma Pengurutan:")
	fmt.Println("  [1] Selection Sort  -- Pilih minimum, pindahkan ke depan")
	fmt.Println("  [2] Insertion Sort  -- Sisipkan ke posisi yang tepat satu per satu")
	fmt.Println()

	pilihan := bacaInt("  Pilih metode (1/2): ")

	// membuat salinan slice agar urutan asli tidak berubah
	salinan := make([]MenuItem, len(daftarMenu))
	copy(salinan, daftarMenu)

	var metode string

	// menentukan dan menjalankan algoritma sorting sesuai pilihan
	if pilihan == 1 {
		salinan = selectionSort(salinan)
		metode = "Selection Sort"
	} else if pilihan == 2 {
		salinan = insertionSort(salinan)
		metode = "Insertion Sort"
	} else {
		fmt.Println("  Pilihan tidak valid!")
		return
	}

	fmt.Printf("  %s -- harga termurah ke termahal\n", metode)
	cetakDaftarMenu(salinan)
}

// ===================== STATISTIK =====================

// tampilkanStatistik menampilkan statistik jumlah menu per kategori dan rata-rata harga
func tampilkanStatistik() {
	fmt.Println()
	cetakGaris("=", 60)
	fmt.Println("  CAFE-MENU  --  STATISTIK LENGKAP")
	cetakGaris("=", 60)

	if len(daftarMenu) == 0 {
		fmt.Println("  Belum ada data menu.")
		return
	}

	// menggunakan map untuk menghitung jumlah dan total harga per kategori
	kategoriCount := make(map[string]int)
	kategoriTotal := make(map[string]float64)
	var totalHarga float64
	totalTersedia := 0

	// melakukan iterasi seluruh menu untuk mengumpulkan data statistik
	for _, m := range daftarMenu {
		kategoriCount[m.Kategori] = kategoriCount[m.Kategori] + 1
		kategoriTotal[m.Kategori] = kategoriTotal[m.Kategori] + m.Harga
		totalHarga = totalHarga + m.Harga
		if m.Tersedia {
			totalTersedia++
		}
	}

	// menghitung nilai rata-rata dan status ketersediaan
	rataRata := totalHarga / float64(len(daftarMenu))
	totalTidakTersedia := len(daftarMenu) - totalTersedia

	fmt.Println()
	fmt.Printf("  Total Semua Menu  : %d item\n", len(daftarMenu))
	fmt.Printf("  Rata-rata Harga   : Rp %.0f\n", rataRata)
	fmt.Printf("  Menu Tersedia     : %d item\n", totalTersedia)
	fmt.Printf("  Menu Habis        : %d item\n", totalTidakTersedia)

	// menampilkan tabel detail per kategori
	fmt.Println()
	fmt.Println("  Detail Per Kategori:")
	cetakGaris("-", 58)
	fmt.Printf("  %-20s  %5s  %s\n", "Kategori", "Total", "Rata-rata Harga")
	cetakGaris("-", 58)

	// menampilkan baris statistik untuk setiap kategori
	for kat, count := range kategoriCount {
		rataKat := kategoriTotal[kat] / float64(count)
		fmt.Printf("  %-20s  %5d  Rp %.0f\n", kat, count, rataKat)
	}
	cetakGaris("-", 58)

	// menampilkan bar chart proporsi menu per kategori
	fmt.Println()
	fmt.Println("  Proporsi Menu (Bar Chart):")
	fmt.Println()
	// menghitung dan menampilkan bar untuk setiap kategori
	for kat, count := range kategoriCount {
		persen := float64(count) / float64(len(daftarMenu)) * 100
		// menghitung panjang bar berdasarkan persentase dibagi 3
		barLen := int(persen / 3)
		if barLen < 1 {
			barLen = 1
		}
		bar := ulangiKarakter("#", barLen)
		sisa := ulangiKarakter(".", 33-barLen)
		fmt.Printf("  %-14s [%s%s] %.1f%%\n", kat, bar, sisa, persen)
	}
	fmt.Println()
	cetakGaris("=", 60)
}

// ===================== DATA CONTOH =====================

// isiDataContoh mengisi daftarMenu dengan data awal yang realistis
func isiDataContoh() {
	contoh := []MenuItem{
		{ID: nextID, Nama: "Espresso", Kategori: "coffee", Harga: 18000, Komposisi: "espresso-shot,air-panas", Tersedia: true},
		{ID: nextID + 1, Nama: "Cappuccino", Kategori: "coffee", Harga: 28000, Komposisi: "espresso,susu-steamed,foam", Tersedia: true},
		{ID: nextID + 2, Nama: "Caramel-Latte", Kategori: "coffee", Harga: 35000, Komposisi: "espresso,susu,karamel,whip-cream", Tersedia: true},
		{ID: nextID + 3, Nama: "Cold-Brew", Kategori: "coffee", Harga: 30000, Komposisi: "kopi-arabika-cold-brew-12-jam", Tersedia: false},
		{ID: nextID + 4, Nama: "Matcha-Latte", Kategori: "non-coffee", Harga: 32000, Komposisi: "matcha-powder,susu-full-cream,gula-tebu", Tersedia: true},
		{ID: nextID + 5, Nama: "Teh-Tarik", Kategori: "non-coffee", Harga: 15000, Komposisi: "teh-hitam,susu-kental-manis", Tersedia: true},
		{ID: nextID + 6, Nama: "Es-Cokelat", Kategori: "non-coffee", Harga: 22000, Komposisi: "cokelat-premium,susu,es-batu", Tersedia: true},
		{ID: nextID + 7, Nama: "Croissant", Kategori: "makanan", Harga: 22000, Komposisi: "tepung-premium,mentega,telur", Tersedia: true},
		{ID: nextID + 8, Nama: "Sandwich-Ayam", Kategori: "makanan", Harga: 38000, Komposisi: "roti-sourdough,ayam-panggang,selada,tomat,mayo", Tersedia: false},
		{ID: nextID + 9, Nama: "Cheesecake-Matcha", Kategori: "dessert", Harga: 35000, Komposisi: "cream-cheese,matcha,graham-cracker", Tersedia: true},
	}
	daftarMenu = append(daftarMenu, contoh...)
	nextID = nextID + len(contoh)
}

// ===================== MAIN =====================

func main() {
	isiDataContoh()

	fmt.Println()
	cetakGaris("=", 60)
	fmt.Println("  Selamat datang di Cafe-Menu!")
	fmt.Println("  Aplikasi Katalog Menu Digital Cafe")
	fmt.Println("  Tugas Besar Algoritma Pemrograman 2")
	cetakGaris("=", 60)

	jalan := true
	for jalan {
		cetakMainMenu()
		pilihan := bacaInt("  Pilihan Anda: ")

		// menjalankan fungsi berdasarkan pilihan dari menu utama
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
