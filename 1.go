// +++ Cafe-Menu +++
// Aplikasi Katalog Menu Digital Cafe — Tugas Besar Algoritma Pemrograman 2
// Dengan tampilan terminal berwarna dan menarik

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// Kode warna ANSI untuk tampilan terminal yang indah (Angelina)
const (
	Reset = "\033[0m"
	Bold  = "\033[1m"
	Dim   = "\033[2m"

	// Foreground colors
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	White  = "\033[37m"

	// Bright foreground colors
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"

	// Background
	BgBrightBlack = "\033[100m"
)

// ===================== STRUCT =====================

// MenuItem menyimpan semua data setiap item menu cafe (Putri)
type MenuItem struct {
	ID        int
	Nama      string
	Kategori  string
	Harga     float64
	Komposisi string
	Tersedia  bool
}

// ===================== VARIABEL GLOBAL =====================

// daftarMenu adalah slice utama penyimpanan semua menu (Syifa)
var daftarMenu []MenuItem
var nextID = 1
var reader = bufio.NewReader(os.Stdin)

// ===================== HELPER =====================

// colorize membungkus teks dengan warna ANSI (Angelina)
func colorize(color, text string) string {
	return color + text + Reset
}

// clearScreen membersihkan layar terminal (Putri)
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// sleep menunggu selama n milidetik (Syifa)
func sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

// printSlow mencetak teks karakter per karakter seperti efek mengetik (Angelina)
func printSlow(text string, delay int) {
	for _, ch := range text {
		fmt.Print(string(ch))
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

// bacaString membaca input string dari pengguna (Putri)
func bacaString(prompt string) string {
	fmt.Print(colorize(BrightCyan, "  ❯ ") + colorize(BrightWhite, prompt))
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// bacaInt membaca input integer dengan validasi (Syifa)
func bacaInt(prompt string) int {
	for {
		fmt.Print(colorize(BrightYellow, "  ❯ ") + colorize(BrightWhite, prompt))
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		val, err := strconv.Atoi(text)
		if err == nil {
			return val
		}
		fmt.Println(colorize(BrightRed, "  ✗ Input harus berupa angka! Coba lagi."))
	}
}

// bacaFloat membaca input float64 dengan validasi (Angelina)
func bacaFloat(prompt string) float64 {
	for {
		fmt.Print(colorize(BrightCyan, "  ❯ ") + colorize(BrightWhite, prompt))
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		val, err := strconv.ParseFloat(text, 64)
		if err == nil {
			return val
		}
		fmt.Println(colorize(BrightRed, "  ✗ Input harus berupa angka! Coba lagi."))
	}
}

// bacaBool membaca pilihan ya/tidak dari pengguna (Putri)
func bacaBool(prompt string) bool {
	for {
		input := strings.ToLower(bacaString(prompt + " (y/n): "))
		if input == "y" || input == "yes" {
			return true
		} else if input == "n" || input == "no" {
			return false
		}
		fmt.Println(colorize(BrightRed, "  ✗ Masukkan 'y' untuk Ya atau 'n' untuk Tidak."))
	}
}

// ===================== TAMPILAN =====================

// animasiSplash menampilkan splash screen animasi saat pertama buka (Syifa)
func animasiSplash() {
	clearScreen()
	sleep(100)

	// logo ASCII art untuk tampilan splash screen (Angelina)
	logo := []string{
		``,
		`   ██████╗ █████╗ ███████╗███████╗    ███╗   ███╗███████╗███╗   ██╗██╗   ██╗`,
		`  ██╔════╝██╔══██╗██╔════╝██╔════╝    ████╗ ████║██╔════╝████╗  ██║██║   ██║`,
		`  ██║     ███████║█████╗  █████╗      ██╔████╔██║█████╗  ██╔██╗ ██║██║   ██║`,
		`  ██║     ██╔══██║██╔══╝  ██╔══╝      ██║╚██╔╝██║██╔══╝  ██║╚██╗██║██║   ██║`,
		`  ╚██████╗██║  ██║██║     ███████╗    ██║ ╚═╝ ██║███████╗██║ ╚████║╚██████╔╝`,
		`   ╚═════╝╚═╝  ╚═╝╚═╝     ╚══════╝    ╚═╝     ╚═╝╚══════╝╚═╝  ╚═══╝ ╚═════╝`,
		``,
	}

	colors := []string{BrightRed, BrightYellow, BrightGreen, BrightCyan, BrightBlue, BrightMagenta, BrightWhite}
	// menampilkan logo dengan warna berbeda tiap baris untuk efek pelangi (Putri)
	for i, line := range logo {
		fmt.Println(colorize(colors[i%len(colors)], line))
		sleep(60)
	}

	fmt.Println(colorize(BrightYellow, Bold+"   ☕  Aplikasi Katalog Menu Digital Cafe  ☕"+Reset))
	fmt.Println(colorize(Dim, "   Tugas Besar Algoritma Pemrograman 2  |  eel / @jebb_24"))
	fmt.Println()

	// menampilkan animasi loading bar sebanyak 20 langkah (Syifa)
	bar := ""
	for i := 0; i < 20; i++ {
		bar += "█"
		fmt.Printf("\r   %s [%-20s] %d%%", colorize(Cyan, "Memuat sistem"), colorize(BrightGreen, bar), (i+1)*5)
		sleep(50)
	}
	fmt.Println(colorize(BrightGreen, " ✓ Siap!"))
	sleep(400)
}

// cetakDivider mencetak garis pemisah berwarna (Angelina)
func cetakDivider(char string, lebar int, warna string) {
	fmt.Println(colorize(warna, strings.Repeat(char, lebar)))
}

// cetakHeaderSection menampilkan header section dengan dekorasi (Putri)
func cetakHeaderSection(judul, icon string) {
	fmt.Println()
	cetakDivider("─", 60, Cyan)
	fmt.Printf("  %s  %s\n", icon, colorize(Bold+BrightCyan, judul))
	cetakDivider("─", 60, Cyan)
}

// cetakMainMenu menampilkan menu utama dengan tampilan menarik (Syifa)
func cetakMainMenu() {
	clearScreen()
	fmt.Println()
	cetakDivider("═", 62, BrightYellow)
	fmt.Println(colorize(BrightYellow, "  ╔══════════════════════════════════════════════════════╗"))
	fmt.Printf("  ║  %s%-52s%s  ║\n", Bold+BrightYellow, "☕  +++ CAFE-MENU +++", Reset+BrightYellow)
	fmt.Printf("  ║  %s%-56s%s  ║\n", Dim, "Aplikasi Katalog Menu Digital Cafe", Reset+BrightYellow)
	fmt.Println(colorize(BrightYellow, "  ╚══════════════════════════════════════════════════════╝"))
	cetakDivider("═", 62, BrightYellow)
	fmt.Println()

	// mendefinisikan item-item menu utama dengan ikon dan warna (Angelina)
	type itemMenu struct{ num, icon, label, warna string }
	menuItems := []itemMenu{
		{"1", "📋", "Tampilkan Semua Menu", BrightGreen},
		{"2", "➕", "Tambah Menu Baru", BrightCyan},
		{"3", "✏️ ", "Ubah Data Menu", BrightYellow},
		{"4", "🗑 ", "Hapus Menu", BrightRed},
		{"5", "🔍", "Cari Menu  (Sequential / Binary Search)", BrightMagenta},
		{"6", "📊", "Urutkan Menu (Selection / Insertion Sort)", BrightBlue},
		{"7", "📈", "Statistik Cafe", BrightYellow},
		{"0", "🚪", "Keluar", Dim + White},
	}

	// menampilkan setiap item menu utama dengan warna dan ikon (Putri)
	for _, item := range menuItems {
		fmt.Printf("  %s %s  %s\n",
			colorize(BgBrightBlack+item.warna+Bold, " "+item.num+" "),
			item.icon,
			colorize(item.warna, item.label),
		)
	}
	fmt.Println()
	cetakDivider("─", 62, Yellow)
	fmt.Printf("  %s %s\n", colorize(BrightGreen, "●"), colorize(Dim, fmt.Sprintf("Total menu: %d item", len(daftarMenu))))
}

// ===================== TAMPILAN KARTU MENU =====================

// ikonKategori mengembalikan emoji yang sesuai dengan kategori menu (Syifa)
func ikonKategori(kat string) string {
	switch strings.ToLower(kat) {
	case "coffee":
		return "☕"
	case "non-coffee":
		return "🧋"
	case "makanan":
		return "🍽"
	case "dessert":
		return "🍰"
	case "snack":
		return "🍟"
	default:
		return "🍴"
	}
}

// warnaKategori mengembalikan kode warna ANSI yang sesuai kategori (Angelina)
func warnaKategori(kat string) string {
	switch strings.ToLower(kat) {
	case "coffee":
		return BrightYellow
	case "non-coffee":
		return BrightCyan
	case "makanan":
		return BrightGreen
	case "dessert":
		return BrightMagenta
	default:
		return BrightWhite
	}
}

// cetakKartuMenu menampilkan satu item menu dalam format kartu dengan bingkai (Putri)
func cetakKartuMenu(m MenuItem, nomor int) {
	wKat := warnaKategori(m.Kategori)
	ikKat := ikonKategori(m.Kategori)

	statusStr := colorize(BrightGreen, "● TERSEDIA")
	if !m.Tersedia {
		statusStr = colorize(BrightRed, "○ HABIS   ")
	}

	// memotong komposisi jika melebihi 44 karakter agar tidak merusak tampilan (Syifa)
	komposisi := m.Komposisi
	if len(komposisi) > 44 {
		komposisi = komposisi[:41] + "..."
	}

	// mencetak bingkai kartu menu menggunakan unicode box drawing characters (Angelina)
	fmt.Printf("  %s\n", colorize(wKat, "┌─────────────────────────────────────────────────────┐"))
	fmt.Printf("  %s  %s  %s%s\n",
		colorize(wKat, "│"),
		colorize(BgBrightBlack+Bold, fmt.Sprintf(" #%02d ", m.ID)),
		colorize(Bold+BrightWhite, fmt.Sprintf("%-38s", m.Nama)),
		colorize(wKat, "│"),
	)
	fmt.Printf("  %s  %s  %-12s %s  Rp %s  %s\n",
		colorize(wKat, "│"),
		ikKat,
		colorize(wKat+Bold, fmt.Sprintf("%-12s", m.Kategori)),
		colorize(Dim, "│"),
		colorize(BrightYellow+Bold, fmt.Sprintf("%8.0f", m.Harga)),
		colorize(wKat, "│"),
	)
	fmt.Printf("  %s  %s %-46s %s\n",
		colorize(wKat, "│"),
		colorize(Dim, "⚗"),
		colorize(Dim, komposisi),
		colorize(wKat, "│"),
	)
	fmt.Printf("  %s  %s%-40s %s\n",
		colorize(wKat, "│"),
		statusStr,
		"",
		colorize(wKat, "│"),
	)
	fmt.Printf("  %s\n", colorize(wKat, "└─────────────────────────────────────────────────────┘"))

	// memberikan jarak antar setiap 2 kartu (Putri)
	if nomor%2 == 0 {
		fmt.Println()
	}
}

// cetakDaftarMenu menampilkan semua menu dalam slice sebagai kartu (Syifa)
func cetakDaftarMenu(list []MenuItem) {
	if len(list) == 0 {
		fmt.Println()
		fmt.Println(colorize(BrightRed, "  ✗ Tidak ada data menu yang ditemukan."))
		return
	}
	fmt.Println()
	// melakukan perulangan untuk menampilkan setiap kartu menu (Angelina)
	for i, m := range list {
		cetakKartuMenu(m, i+1)
	}
	fmt.Printf("\n  %s %s\n", colorize(BrightGreen, "✓"), colorize(Dim, fmt.Sprintf("Menampilkan %d menu", len(list))))
}

// ===================== CRUD MENU =====================

// tambahMenu menambahkan menu baru ke dalam daftarMenu (Putri)
func tambahMenu() {
	cetakHeaderSection("TAMBAH MENU BARU", "➕")

	nama := bacaString("Nama Menu            : ")
	if nama == "" {
		fmt.Println(colorize(BrightRed, "  ✗ Nama tidak boleh kosong!"))
		return
	}

	fmt.Println()
	fmt.Println(colorize(Dim, "  Contoh kategori: coffee | non-coffee | makanan | dessert | snack"))
	kategori := strings.ToLower(bacaString("Kategori             : "))
	harga := bacaFloat("Harga (Rp)           : ")
	komposisi := bacaString("Komposisi Bahan      : ")
	tersedia := bacaBool("Tersedia sekarang    ")

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
	cetakDivider("─", 60, BrightGreen)
	fmt.Printf("  %s Menu %s berhasil ditambahkan! (ID: #%02d)\n",
		colorize(BrightGreen, "✓"),
		colorize(BrightWhite+Bold, "'"+nama+"'"),
		item.ID,
	)
	cetakDivider("─", 60, BrightGreen)
}

// ubahMenu mengubah data menu yang sudah ada berdasarkan ID yang dipilih (Syifa)
func ubahMenu() {
	cetakHeaderSection("UBAH DATA MENU", "✏️")

	if len(daftarMenu) == 0 {
		fmt.Println(colorize(BrightRed, "  ✗ Belum ada menu!"))
		return
	}
	cetakDaftarMenu(daftarMenu)

	id := bacaInt("Masukkan ID menu yang ingin diubah: ")

	// mencari posisi menu di dalam slice berdasarkan ID (Angelina)
	idx := -1
	for i, m := range daftarMenu {
		if m.ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		fmt.Println(colorize(BrightRed, "  ✗ ID tidak ditemukan!"))
		return
	}

	m := &daftarMenu[idx]
	fmt.Println()
	fmt.Printf("  %s Mengedit: %s\n", colorize(BrightYellow, "◆"), colorize(Bold+BrightWhite, m.Nama))
	fmt.Println(colorize(Dim, "  (Tekan Enter untuk melewati / tidak mengubah field)"))
	fmt.Println()

	// memperbarui setiap field jika pengguna memberikan input baru (Putri)
	nama := bacaString(fmt.Sprintf("Nama Menu [%s]: ", m.Nama))
	if nama != "" {
		m.Nama = nama
	}
	kategori := bacaString(fmt.Sprintf("Kategori [%s]: ", m.Kategori))
	if kategori != "" {
		m.Kategori = strings.ToLower(kategori)
	}
	hargaStr := bacaString(fmt.Sprintf("Harga [%.0f]: ", m.Harga))
	if hargaStr != "" {
		if val, err := strconv.ParseFloat(hargaStr, 64); err == nil {
			m.Harga = val
		}
	}
	komposisi := bacaString(fmt.Sprintf("Komposisi [%s]: ", m.Komposisi))
	if komposisi != "" {
		m.Komposisi = komposisi
	}
	m.Tersedia = bacaBool(fmt.Sprintf("Tersedia [%v]", m.Tersedia))

	fmt.Println()
	fmt.Printf("  %s Data menu ID #%02d berhasil diperbarui!\n", colorize(BrightGreen, "✓"), id)
}

// hapusMenu menghapus menu dari daftarMenu berdasarkan ID (Syifa)
func hapusMenu() {
	cetakHeaderSection("HAPUS MENU", "🗑")

	if len(daftarMenu) == 0 {
		fmt.Println(colorize(BrightRed, "  ✗ Belum ada menu!"))
		return
	}
	cetakDaftarMenu(daftarMenu)

	id := bacaInt("Masukkan ID menu yang ingin dihapus: ")

	// membangun slice baru tanpa elemen yang akan dihapus (Angelina)
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
		fmt.Println(colorize(BrightRed, "  ✗ ID tidak ditemukan!"))
		return
	}

	konfirmasi := bacaBool(fmt.Sprintf("Yakin menghapus '%s'", namaHapus))
	if konfirmasi {
		daftarMenu = newList
		fmt.Println()
		fmt.Printf("  %s Menu %s berhasil dihapus.\n",
			colorize(BrightRed, "✓"),
			colorize(Bold, "'"+namaHapus+"'"),
		)
	} else {
		fmt.Println(colorize(Yellow, "  ○ Penghapusan dibatalkan."))
	}
}

// tampilkanSemuaMenu menampilkan seluruh isi daftarMenu (Putri)
func tampilkanSemuaMenu() {
	cetakHeaderSection("SEMUA MENU", "📋")
	cetakDaftarMenu(daftarMenu)
}

// ===================== PENCARIAN =====================

// sequentialSearch mencari menu berdasarkan kategori dari awal hingga akhir (Syifa)
func sequentialSearch(kategori string) []MenuItem {
	hasil := []MenuItem{}
	// melakukan perulangan satu per satu dari indeks pertama hingga terakhir (Angelina)
	for _, m := range daftarMenu {
		if strings.EqualFold(m.Kategori, kategori) {
			hasil = append(hasil, m)
		}
	}
	return hasil
}

// binarySearch mencari menu berdasarkan kategori menggunakan algoritma biner (Putri)
// slice diurutkan sementara sebelum pencarian dilakukan
func binarySearch(kategori string) []MenuItem {
	// membuat salinan slice dan mengurutkannya berdasarkan kategori (Syifa)
	sorted := make([]MenuItem, len(daftarMenu))
	copy(sorted, daftarMenu)

	// mengurutkan salinan secara abjad berdasarkan kategori menggunakan insertion sort (Angelina)
	for i := 1; i < len(sorted); i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 && strings.ToLower(sorted[j].Kategori) > strings.ToLower(key.Kategori) {
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = key
	}

	// melakukan binary search untuk menemukan indeks yang cocok (Putri)
	low, high, foundIdx := 0, len(sorted)-1, -1
	for low <= high {
		mid := (low + high) / 2
		cmp := strings.Compare(strings.ToLower(sorted[mid].Kategori), strings.ToLower(kategori))
		if cmp == 0 {
			foundIdx = mid
			break
		} else if cmp < 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if foundIdx == -1 {
		return []MenuItem{}
	}

	// mengumpulkan seluruh elemen dengan kategori sama di sekitar foundIdx (Syifa)
	hasil := []MenuItem{}
	left := foundIdx
	for left > 0 && strings.EqualFold(sorted[left-1].Kategori, kategori) {
		left--
	}
	right := foundIdx
	for right < len(sorted)-1 && strings.EqualFold(sorted[right+1].Kategori, kategori) {
		right++
	}
	for i := left; i <= right; i++ {
		hasil = append(hasil, sorted[i])
	}
	return hasil
}

// menuCari menampilkan UI sub-menu pencarian (Angelina)
func menuCari() {
	cetakHeaderSection("CARI MENU", "🔍")
	fmt.Println()
	fmt.Println(colorize(BrightCyan, "  Pilih Algoritma Pencarian:"))
	fmt.Println()
	fmt.Printf("  %s  Sequential Search  %s\n",
		colorize(BgBrightBlack+BrightCyan+Bold, " 1 "),
		colorize(Dim, "— Menelusuri satu per satu dari awal"),
	)
	fmt.Printf("  %s  Binary Search      %s\n",
		colorize(BgBrightBlack+BrightMagenta+Bold, " 2 "),
		colorize(Dim, "— Membagi dua data yang sudah diurutkan"),
	)
	fmt.Println()

	pilihan := bacaInt("Pilih metode (1/2): ")
	fmt.Println()
	fmt.Println(colorize(Dim, "  Contoh: coffee, non-coffee, makanan, dessert, snack"))
	kategori := strings.ToLower(bacaString("Masukkan kategori yang dicari: "))

	// menampilkan animasi titik-titik saat mencari (Putri)
	fmt.Print(colorize(Cyan, "\n  Mencari"))
	for i := 0; i < 5; i++ {
		fmt.Print(colorize(Cyan, "."))
		sleep(120)
	}

	var hasil []MenuItem
	var metode, warnaMetode string

	// memilih dan menjalankan algoritma pencarian sesuai input (Syifa)
	if pilihan == 1 {
		hasil = sequentialSearch(kategori)
		metode = "Sequential Search"
		warnaMetode = BrightCyan
	} else if pilihan == 2 {
		hasil = binarySearch(kategori)
		metode = "Binary Search"
		warnaMetode = BrightMagenta
	} else {
		fmt.Println(colorize(BrightRed, "\n  ✗ Pilihan tidak valid!"))
		return
	}

	fmt.Printf("\r  %s Hasil %s — kategori %s (%d item)\n",
		colorize(BrightGreen, "✓"),
		colorize(warnaMetode+Bold, metode),
		colorize(Bold+BrightWhite, "'"+kategori+"'"),
		len(hasil),
	)
	cetakDaftarMenu(hasil)
}

// ===================== PENGURUTAN =====================

// selectionSort mengurutkan slice menu berdasarkan harga dari terkecil ke terbesar (Angelina)
func selectionSort(list []MenuItem) []MenuItem {
	n := len(list)
	// melakukan perulangan sebanyak n-1 kali untuk selection sort (Putri)
	for i := 0; i < n-1; i++ {
		minIdx := i
		// mencari elemen dengan harga minimum dari posisi i+1 hingga akhir (Syifa)
		for j := i + 1; j < n; j++ {
			if list[j].Harga < list[minIdx].Harga {
				minIdx = j
			}
		}
		// menukar elemen minimum ke posisi i (Angelina)
		list[i], list[minIdx] = list[minIdx], list[i]
	}
	return list
}

// insertionSort mengurutkan slice menu berdasarkan harga dengan metode penyisipan (Putri)
func insertionSort(list []MenuItem) []MenuItem {
	n := len(list)
	// melakukan perulangan mulai dari indeks ke-1 untuk insertion sort (Syifa)
	for i := 1; i < n; i++ {
		key := list[i]
		j := i - 1
		// menggeser elemen yang lebih besar dari key ke posisi berikutnya (Angelina)
		for j >= 0 && list[j].Harga > key.Harga {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = key
	}
	return list
}

// menuUrut menampilkan UI sub-menu pengurutan harga (Putri)
func menuUrut() {
	cetakHeaderSection("URUTKAN MENU BERDASARKAN HARGA", "📊")
	fmt.Println()
	fmt.Println(colorize(BrightYellow, "  Pilih Algoritma Pengurutan:"))
	fmt.Println()
	fmt.Printf("  %s  Selection Sort  %s\n",
		colorize(BgBrightBlack+BrightYellow+Bold, " 1 "),
		colorize(Dim, "— Pilih minimum, pindahkan ke depan"),
	)
	fmt.Printf("  %s  Insertion Sort  %s\n",
		colorize(BgBrightBlack+BrightBlue+Bold, " 2 "),
		colorize(Dim, "— Sisipkan ke posisi yang tepat satu per satu"),
	)
	fmt.Println()

	pilihan := bacaInt("Pilih metode (1/2): ")

	// membuat salinan slice agar urutan asli tidak berubah (Syifa)
	salinan := make([]MenuItem, len(daftarMenu))
	copy(salinan, daftarMenu)

	var metode, warnaMetode string

	// menentukan dan menjalankan algoritma sorting sesuai pilihan (Angelina)
	if pilihan == 1 {
		salinan = selectionSort(salinan)
		metode = "Selection Sort"
		warnaMetode = BrightYellow
	} else if pilihan == 2 {
		salinan = insertionSort(salinan)
		metode = "Insertion Sort"
		warnaMetode = BrightBlue
	} else {
		fmt.Println(colorize(BrightRed, "  ✗ Pilihan tidak valid!"))
		return
	}

	// animasi proses pengurutan sebelum hasil ditampilkan (Putri)
	fmt.Print(colorize(Cyan, "\n  Mengurutkan"))
	for i := 0; i < 5; i++ {
		fmt.Print(colorize(Cyan, "."))
		sleep(100)
	}
	fmt.Printf("\r  %s %s — harga termurah ke termahal\n",
		colorize(BrightGreen, "✓"),
		colorize(warnaMetode+Bold, metode),
	)
	cetakDaftarMenu(salinan)
}

// ===================== STATISTIK =====================

// tampilkanStatistik menampilkan statistik jumlah menu per kategori dan rata-rata harga (Syifa)
func tampilkanStatistik() {
	clearScreen()
	fmt.Println()
	// menampilkan header statistik dengan format +++ nama aplikasi +++ (Angelina)
	cetakDivider("═", 62, BrightYellow)
	fmt.Println(colorize(BrightYellow+Bold, "  +++ Cafe-Menu +++  —  STATISTIK LENGKAP"))
	cetakDivider("═", 62, BrightYellow)

	if len(daftarMenu) == 0 {
		fmt.Println(colorize(BrightRed, "\n  ✗ Belum ada data menu."))
		return
	}

	// menggunakan map untuk menghitung jumlah dan total harga per kategori (Putri)
	kategoriCount := make(map[string]int)
	kategoriTotal := make(map[string]float64)
	var totalHarga float64
	totalTersedia := 0

	// melakukan iterasi seluruh menu untuk mengumpulkan data statistik (Syifa)
	for _, m := range daftarMenu {
		kategoriCount[m.Kategori]++
		kategoriTotal[m.Kategori] += m.Harga
		totalHarga += m.Harga
		if m.Tersedia {
			totalTersedia++
		}
	}

	// menghitung nilai rata-rata dan status ketersediaan (Angelina)
	rataRata := math.Round(totalHarga/float64(len(daftarMenu))*100) / 100
	totalTidakTersedia := len(daftarMenu) - totalTersedia

	fmt.Println()
	// menampilkan ringkasan angka utama dalam kotak info (Putri)
	fmt.Println(colorize(Cyan, "  ┌─────────────────────────────────────────────────────┐"))
	fmt.Printf("  %s  %-28s  %s\n", colorize(Cyan, "│"), "📦 Total Semua Menu",
		colorize(BrightWhite+Bold, fmt.Sprintf("%d item", len(daftarMenu)))+colorize(Cyan, "                    │"))
	fmt.Printf("  %s  %-28s  %s\n", colorize(Cyan, "│"), "💰 Rata-rata Harga",
		colorize(BrightYellow+Bold, fmt.Sprintf("Rp %.0f", rataRata))+colorize(Cyan, "                 │"))
	fmt.Printf("  %s  %-28s  %s\n", colorize(Cyan, "│"), "✅ Menu Tersedia",
		colorize(BrightGreen+Bold, fmt.Sprintf("%d item", totalTersedia))+colorize(Cyan, "                   │"))
	fmt.Printf("  %s  %-28s  %s\n", colorize(Cyan, "│"), "❌ Menu Habis",
		colorize(BrightRed+Bold, fmt.Sprintf("%d item", totalTidakTersedia))+colorize(Cyan, "                   │"))
	fmt.Println(colorize(Cyan, "  └─────────────────────────────────────────────────────┘"))

	fmt.Println()
	fmt.Println(colorize(BrightMagenta+Bold, "  📂 Detail Per Kategori:"))
	fmt.Println(colorize(Dim, "  ┌────────────────────────┬───────┬──────────────────┐"))
	fmt.Println(colorize(Dim, "  │ Kategori               │ Total │ Rata-rata Harga  │"))
	fmt.Println(colorize(Dim, "  ├────────────────────────┼───────┼──────────────────┤"))

	// menampilkan baris statistik untuk setiap kategori (Syifa)
	for kat, count := range kategoriCount {
		rataKat := math.Round(kategoriTotal[kat]/float64(count)*100) / 100
		wKat := warnaKategori(kat)
		ik := ikonKategori(kat)
		fmt.Printf("  %s %s %-20s %s  %s  %s\n",
			colorize(Dim, "│"),
			ik,
			colorize(wKat+Bold, fmt.Sprintf("%-18s", kat)),
			colorize(Dim, fmt.Sprintf("│  %3d  │", count)),
			colorize(BrightYellow, fmt.Sprintf("Rp %8.0f", rataKat)),
			colorize(Dim, "│"),
		)
	}
	fmt.Println(colorize(Dim, "  └────────────────────────┴───────┴──────────────────┘"))

	// menampilkan bar chart proporsi menu per kategori (Angelina)
	fmt.Println()
	fmt.Println(colorize(BrightBlue+Bold, "  📊 Proporsi Menu (Bar Chart):"))
	fmt.Println()
	// menghitung dan menampilkan bar untuk setiap kategori (Putri)
	for kat, count := range kategoriCount {
		persen := float64(count) / float64(len(daftarMenu)) * 100
		// menghitung panjang bar berdasarkan persentase dibagi 3 (Syifa)
		barLen := int(persen / 3)
		if barLen < 1 {
			barLen = 1
		}
		bar := strings.Repeat("█", barLen)
		sisa := strings.Repeat("░", 33-barLen)
		wKat := warnaKategori(kat)
		ik := ikonKategori(kat)
		fmt.Printf("  %s %-14s %s%s %s%.1f%%%s\n",
			ik,
			colorize(wKat, kat),
			colorize(wKat, bar),
			colorize(Dim, sisa),
			BrightWhite,
			persen,
			Reset,
		)
	}
	fmt.Println()
	cetakDivider("═", 62, BrightYellow)
}

// ===================== DATA CONTOH =====================

// isiDataContoh mengisi daftarMenu dengan data awal yang realistis (Angelina)
func isiDataContoh() {
	contoh := []MenuItem{
		{ID: nextID, Nama: "Espresso", Kategori: "coffee", Harga: 18000, Komposisi: "espresso shot, air panas", Tersedia: true},
		{ID: nextID + 1, Nama: "Cappuccino", Kategori: "coffee", Harga: 28000, Komposisi: "espresso, susu steamed, foam", Tersedia: true},
		{ID: nextID + 2, Nama: "Caramel Latte", Kategori: "coffee", Harga: 35000, Komposisi: "espresso, susu, karamel, whip cream", Tersedia: true},
		{ID: nextID + 3, Nama: "Cold Brew", Kategori: "coffee", Harga: 30000, Komposisi: "kopi arabika cold brew 12 jam", Tersedia: false},
		{ID: nextID + 4, Nama: "Matcha Latte", Kategori: "non-coffee", Harga: 32000, Komposisi: "matcha powder, susu full cream, gula tebu", Tersedia: true},
		{ID: nextID + 5, Nama: "Teh Tarik", Kategori: "non-coffee", Harga: 15000, Komposisi: "teh hitam, susu kental manis", Tersedia: true},
		{ID: nextID + 6, Nama: "Es Cokelat", Kategori: "non-coffee", Harga: 22000, Komposisi: "cokelat premium, susu, es batu", Tersedia: true},
		{ID: nextID + 7, Nama: "Croissant Butter", Kategori: "makanan", Harga: 22000, Komposisi: "tepung premium, mentega prancis, telur", Tersedia: true},
		{ID: nextID + 8, Nama: "Sandwich Ayam Panggang", Kategori: "makanan", Harga: 38000, Komposisi: "roti sourdough, ayam panggang, selada, tomat, mayo", Tersedia: false},
		{ID: nextID + 9, Nama: "Cheesecake Matcha", Kategori: "dessert", Harga: 35000, Komposisi: "cream cheese, matcha, graham cracker base", Tersedia: true},
	}
	daftarMenu = append(daftarMenu, contoh...)
	nextID += len(contoh)
}

// ===================== MAIN =====================

func main() {
	// mengisi data contoh dan menjalankan animasi splash saat aplikasi dimulai (Putri)
	isiDataContoh()
	animasiSplash()

	for {
		cetakMainMenu()
		pilihan := bacaInt("Pilihan Anda: ")

		// menjalankan fungsi berdasarkan pilihan dari menu utama (Syifa)
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
			clearScreen()
			fmt.Println()
			cetakDivider("═", 62, BrightYellow)
			printSlow(colorize(BrightYellow, "  ☕  Terima kasih telah menggunakan +++ Cafe-Menu +++"), 18)
			fmt.Println()
			printSlow(colorize(Dim, "  Sampai jumpa!  —  eel / @jebb_24"), 15)
			fmt.Println()
			cetakDivider("═", 62, BrightYellow)
			fmt.Println()
			os.Exit(0)
		default:
			fmt.Println()
			fmt.Println(colorize(BrightRed, "  ✗ Pilihan tidak valid! Silakan pilih angka 0-7."))
		}

		fmt.Println()
		bacaString("Tekan Enter untuk kembali ke menu utama...")
	}
}
