/*
input data default dulu
sparepart :
busi 50000 0
rantai 80000 0
oli 60000 0

pelanggan :
agus
asep
muldi

transaksi :
kosong
*/
package main

import "fmt"

const NMAX = 1000

type sparepart struct {
	nama           string
	harga, diganti int
}

type time struct {
	tanggal, bulan, tahun int
}

type transaksi struct {
	// n menyatakan banyaknya transaksi //
	id, tarif, n int
	waktu        time
	plgn         pelanggan
	spr          sparepart
}

type pelanggan struct {
	nama string
}

type customer [NMAX]string

type histori [NMAX]transaksi

type sukucadang [NMAX]sparepart

func main() {
	var arrPelanggan customer
	var arrSparepart sukucadang
	var s, p int

	mulai()
	nambah_dataPelanggan(&arrPelanggan, &p)
	nambah_dataSparepart(&arrSparepart, &s)
}

func mulai() {
	var blank string

	fmt.Println("========================================")
	fmt.Println("=       APLIKASI SERVICE MOTOR         =")
	fmt.Println("=           Selamat Datang             =")
	fmt.Println("========================================")
	fmt.Print("Tekan tombol apapun untuk lanjut...")
	fmt.Scanln(&blank)
	fmt.Println("")
	menuawal()
}

func menuawal() {
	var n int

	for n != 6 {
		fmt.Println("========================================")
		fmt.Println("=       APLIKASI SERVICE MOTOR         =")
		fmt.Println("=             Menu Utama               =")
		fmt.Println("========================================")
		fmt.Println("=  1. Penambahan Data                  =")
		fmt.Println("=  2. Pengubahan Data                  =")
		fmt.Println("=  3. Penghapusan Data                 =")
		fmt.Println("=  4. Pencarian Data                   =")
		fmt.Println("=  5. Menampilkan Data                 =")
		fmt.Println("=  6. Keluar                           =")
		fmt.Println("= ------------------------------------ =")
		fmt.Print("Masukkan pilihan anda : ")
		fmt.Scanln(&n)
		fmt.Println("========================================")
		if n == 1 {
			nambah_data()
		} else if n == 2 {
			//ubah_data()
		} else if n == 3 {
			//hapus_data()
		} else if n == 4 {
			//cari_data()
		} else if n == 5 {
			tampil_data()
		} else {
			fmt.Print("Pilihan Anda tidak valid. Silakan pilih opsi yang tersedia!")
		}
	}
	fmt.Println("")
	fmt.Println("========================================")
	fmt.Println("=    Terima kasih telah menggunakan    =")
	fmt.Println("=        APLIKASI SERVICE MOTOR        =")
	fmt.Println("========================================")
	fmt.Println("")
}

func nambah_data() {
	var arrPelanggan customer
	var arrTransaksi histori
	var arrSparepart sukucadang
	var n int
	var s, p, t int

	fmt.Println("========================================")
	fmt.Println("=       APLIKASI SERVICE MOTOR         =")
	fmt.Println("=        Page Penambahan Data          =")
	fmt.Println("========================================")
	fmt.Println("=  Pilih data yang ingin ditambah      =")
	fmt.Println("=  1. Pelanggan                        =")
	fmt.Println("=  2. Sparepart                        =")
	fmt.Println("=  3. Transaksi                        =")
	fmt.Println("=  4. Keluar                           =")
	fmt.Println("= ------------------------------------ =")
	fmt.Print("=  Masukkan pilihan anda : ")
	fmt.Scan(&n)
	fmt.Println("========================================")
	if n == 1 {
		nambah_dataPelanggan(&arrPelanggan, &p)
	} else if n == 2 {
		nambah_dataSparepart(&arrSparepart, &s)
	} else if n == 3 {
		nambah_dataTransaksi(&arrTransaksi, &t)
	}
}

/*func nambah_data(arrTransaksi *histori, idx transaksi) {
	fmt.Println("========================================")
	fmt.Println("=       APLIKASI SERVICE MOTOR         =")
	fmt.Println("=        Page Penambahan Data          =")
	fmt.Println("========================================")
	fmt.Print("  Masukkan nama customer: ")
	fmt.Scan(&arrTransaksi[idx.n].pelanggan.nama)
	fmt.Print("  Masukkan tanggal transaksi (DD/MM/YY): ")
	fmt.Scan(&arrTransaksi[idx.n].waktu)
	fmt.Print("  Masukkan tarif: ")
	fmt.Scan(&arrTransaksi[idx.n].tarif)
	fmt.Print("  Masukkan nama sparepart yang diperbaiki: ")
	fmt.Scan(&arrTransaksi[idx.n].sparepart.nama)
	fmt.Println("========================================")
	arr
	idx.n++

}*/

func nambah_dataPelanggan(T *customer, p *int) {
	fmt.Println("== Aplikasi Service Motor ==")
	fmt.Println("  Penambahan Data Pelanggan ")
	fmt.Println("----------------------------")
	fmt.Print("Nama : ")
	fmt.Scanln(&T[*p])
	fmt.Println("----------------------------")
}

func nambah_dataSparepart(T *sukucadang, s *int) {
	fmt.Println("== Aplikasi Service Motor ==")
	fmt.Println("  Penambahan Data Sparepart ")
	fmt.Println("----------------------------")
	fmt.Print("Nama Sparepart : ")
	fmt.Scanln(&T[*s].nama)
	fmt.Print("Harga Sparepart : ")
	fmt.Scanln(&T[*s].harga)
	fmt.Println("----------------------------")
}

func nambah_dataTransaksi(T *histori, t *int) {
	fmt.Println("== Aplikasi Service Motor ==")
	fmt.Println("  Penambahan Data Sparepart ")
	fmt.Println("----------------------------")
	fmt.Print("Nama Pelanggan : ")
	fmt.Scanln(&T[*t].plgn)
	fmt.Print("Waktu : ")
	fmt.Scanln(&T[*t].tarif)
	fmt.Print("Tarif : ")
	fmt.Scanln(&T[*t].waktu)
	fmt.Print("Sparepart yang diganti : ")
	fmt.Scanln(&T[*t].spr)
	fmt.Println("----------------------------")
}

/*
func ubah_data() {
	var arrPelanggan customer
	var arrTransaksi histori
	var arrSparepart sukucadang
	var n int

	fmt.Println("========================================")
	fmt.Println("=       APLIKASI SERVICE MOTOR         =")
	fmt.Println("=        Page Pengubahan Data          =")
	fmt.Println("========================================")
	fmt.Println("=  Pilih data yang ingin diubah        =")
	fmt.Println("=  1. Pelanggan                        =")
	fmt.Println("=  2. Sparepart                        =")
	fmt.Println("=  3. Transaksi")
	fmt.Println("=  4. Keluar")
	fmt.Print("Masukkan pilihan anda : ")
	fmt.Println("----------------------------")
	fmt.Scan(&n)
	if n == 1 {
		ubah_dataPelanggan(&arrPelanggan)
	} else if n == 2 {
		ubah_dataSparepart(&arrTransaksi)
	} else if n == 3 {
		ubah_dataTransaksi(&arrSparepart)
	}
}

func hapus_data() {
	var arrPelanggan pelanggan
	var arrTransaksi histori
	var arrSparepart sukucadang
	var n int

	fmt.Println("== Aplikasi Service Motor ==")
	fmt.Println("----------------------------")
	fmt.Println("Pilih data yang ingin dihapus")
	fmt.Println("1. Pelanggan")
	fmt.Println("2. Sparepart")
	fmt.Println("3. Transaksi")
	fmt.Println("4. Keluar")
	fmt.Print("Masukkan pilihan anda : ")
	fmt.Println("----------------------------")
	fmt.Scan(&n)
	if n == 1 {
		hapus_dataPelanggan(&arrPelanggan)
	} else if n == 2 {
		hapus_dataSparepart(&arrTransaksi)
	} else if n == 3 {
		hapus_dataTransaksi(&arrSparepart)
	}
}

func cari_data() {
	var arrPelanggan customer
	var arrTransaksi histori
	var arrSparepart sukucadang
	var n int

	fmt.Println("========================================")
	fmt.Println("=       APLIKASI SERVICE MOTOR         =")
	fmt.Println("=         Page Pencarian Data          =")
	fmt.Println("========================================")
	fmt.Println("=  Pilih data yang ingin dicari        =")
	fmt.Println("=  1. Pelanggan                        =")
	fmt.Println("=  2. Sparepart                        =")
	fmt.Println("=  3. Transaksi                        =")
	fmt.Println("=  4. Keluar                           =")
	fmt.Println("----------------------------------------")
	fmt.Print("   Masukkan pilihan anda : ")
	fmt.Scan(&n)
	fmt.Println("========================================")
	if n == 1 {
		cari_dataPelanggan(&arrPelanggan)
	} else if n == 2 {
		cari_dataSparepart(&arrTransaksi)
	} else if n == 3 {
		cari_dataTransaksi(&arrSparepart)
	}
}
*/

func tampil_data() {
	var arrPelanggan customer
	var arrTransaksi histori
	var arrSparepart sukucadang
	var n, p, s, t int

	fmt.Println("========================================")
	fmt.Println("=       APLIKASI SERVICE MOTOR         =")
	fmt.Println("=       Page Menampilkan Data          =")
	fmt.Println("========================================")
	fmt.Println("=  Pilih data yang ingin ditambah      =")
	fmt.Println("=  1. Pelanggan                        =")
	fmt.Println("=  2. Sparepart                        =")
	fmt.Println("=  3. Transaksi                        =")
	fmt.Println("=  4. Keluar                           =")
	fmt.Println("----------------------------------------")
	fmt.Print("   Masukkan pilihan anda : ")
	fmt.Scan(&n)
	fmt.Println("========================================")
	if n == 1 {
		tampil_dataPelanggan(&arrPelanggan, &p)
	} else if n == 2 {
		tampil_dataSparepart(&arrSparepart, &s)
	} else if n == 3 {
		tampil_dataTransaksi(&arrTransaksi, &t)
	}
}

func tampil_dataPelanggan(T *customer, p *int) {
	for i := 0; i < *p; i++ {
		fmt.Println(i, "\t")
		fmt.Println(T[i])
	}
}

func tampil_dataSparepart(T *sukucadang, s *int) {
	for i := 0; i < *s; i++ {
		fmt.Println(i, "\t")
		fmt.Println(T[i])
	}
}

func tampil_dataTransaksi(T *histori, t *int) {
	for i := 0; i < *t; i++ {
		fmt.Println(i, "\t")
		fmt.Println(T[i])
	}
}
