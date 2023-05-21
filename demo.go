package main

import "fmt"

// Deklarasi Konstanta
const nmax int = 5
const usermax int = 100
const lamplife int = 30

// Deklarasi Tipe Bentukan island
type island struct {
	nama    [nmax]int
	nmember int
}

// Deklarasi Tipe Bentukan bridge
type bridge struct {
	m1, m2 int
}

// Deklarasi Tipe Bentukan user
type user struct {
	nama string
	time int
}

// Deklarasi Tipe Bentukan users
type users struct {
	info  [usermax]user
	nuser int
}

func main() {
	// deklarasi variabel
	var u users
	// panggil procedure menu
	menu(&u)
}

func header() {
	/* I.S. -
	   F.S. menampilkan header dari aplikasi */
	fmt.Println("\n ===== Selamat Datang ===== ")
	fmt.Println("    Bridge Crossing Game    ")
	fmt.Println("   Algoritma Pemrograman    ")
	fmt.Println(" -------------------------- ")
}

func menu(A *users) {
	/* I.S. terdefinisi tipe bentukan A yang berisi sejumlah A.nUser pemain pada array A.info
	   F.S. menampilkan menu permainan, dan array A.info terupdate apabila terdapat pemain baru */
	// a. Buatlah variabel string opsi yang digunakan untuk menyimpan pilihan menu yang dipilih, inisialisasi dengan string kosong
	var menudipilih string
	// b. Buatlah while dengan kondisi variabel opsi tidak sama dengan string "3" karena pilihan "3" digunakan untuk keluar dari aplikasi
	for menudipilih != "3" {
		// c. panggil procedure header() yang sudah dibuat. Kemudian tampilkan pilihan opsi 1 hingga 3 sesuai gambar tampilan menu di atas
		header()
		fmt.Println(" 1. Mulai Permainan")
		fmt.Println(" 2. Lihat Waktu")
		fmt.Println(" 3. Keluar")
		fmt.Println(" --------------------------- ")
		// d. Tampilkan string "Pilihan anda: ", kemudian minta masukan untuk variabel opsi
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&menudipilih)
		// e. Lakukan perulangan dengan while untuk masukan variabel opsi selama opsi yang diberikan tidak valid (selain 1, 2 dan 3).
		for menudipilih != "1" && menudipilih != "2" && menudipilih != "3" {
			fmt.Print("Pilihan Anda: ")
			fmt.Scan(&menudipilih)
		}
		// f. Tambahkan percabangan untuk setiap nilai variabel opsi "1" dan "2" sebelum endwhile. Panggil procedure mulai_permainan(A) untuk opsi "1" dan lihat_skor(*A) untuk opsi "2"
		if menudipilih == "1" {
			mulai_permainan(&*A)
		} else if menudipilih == "2" {
			lihat_skor(*A)
		}
	}
	// g. Setelah endwhile tampilkan string "Terima kasih :)".
	fmt.Println("Terima kasih :)")
}

func lihat_skor(A users) {
	/* I.S. terdefinisi tipe bentukan A yang berisi sejumlah A.nUser pemain pada array A.info
	   F.S. menampilkan semua data pemain pada array A.info */
	for i := 0; i < A.nuser; i++ {
		fmt.Println(i+1, ".", "Nama: ", A.info[i].nama, " dengan Waktu: ", A.info[i].time)
	}
}

func setup_permainan(left, right *island) {
	/* I.S. -
	   F.S. left.nMember bernilai 0 dan right.nMember bernilai 5, sedangkan right.member berisi 1, 3, 6, 8, dan 12 */
	left.nmember = 0
	right.nama = [nmax]int{1, 3, 6, 8, 12}
	right.nmember = 5
}

func sequential_search(A island, x int) int {
	/* mengembalikan indeks dari x di dalam array A.info, atau -1 apabila tidak ditemukan */
	for i := 0; i < A.nmember; i++ {
		if x == A.nama[i] {
			return i
		}
	}
	return -1
}

func add_member(A *island, x int) {
	/* I.S. terdefinisi tipe bentukan A yang berisi array A.info sebanyak A.nMember elemen. dan sebuah bilangan bulat x
	   F.S. menambahkan x kedalam array A.info dan A.nMember bertambah 1, apabila array penuh tampilkan "Pulau telah penuh"*/
	i := 0
	if nmax >= A.nmember+i {
		A.nama[A.nmember+i] = x
		A.nmember++
	} else {
		fmt.Println("Pulau telah penuh")
	}
}

func del_member(A *island, x int) {
	/* I.S. terdefinisi tipe bentukan A yang berisi array A.info sebanyak A.nMember elemen. dan sebuah bilangan bulat x
	   F.S. menghapus elemen dengan nilai x  dari array A.info, dan A.nMember berkurang 1, apabila x tidak ditemukan maka tampilkan "Orang yang dicari tidak ditemukan"*/
	if A.nmember != 0 {
		// lakukan pencarian x di dalam A dengan memanggil fungsi sequential_search
		var idx = sequential_search(*A, x)
		// apabila data ditemukan maka lakukan proses penghapusan
		if idx != -1 {
			for idx < A.nmember-1 {
				A.nama[idx] = A.nama[idx+1]
				idx++
			}
		} else {
			fmt.Println("Orang yang dicari tidak ditemukan")
		}
	}
}

func print_step(left, right island, time int) {
	/* I.S. terdefinisi tipe bentukan left dan right yang berisi left.nMember dan right.nMember orang, serta sebuah bilangan bulat time
	   F.S. menampilkan array left.member dan right.member serta sisa waktu secara horizontal */
	fmt.Print(" Kiri: ")
	var i int
	// tampilkan isi array left.member secara horizontal
	for i = 0; i < left.nmember; i++ {
		fmt.Print(left.nama[i], " ")
	}
	fmt.Print("\t\t Kanan: ")
	// tampilkan isi array right.member secara horizontal
	for i = 0; i < right.nmember; i++ {
		fmt.Print(right.nama[i], " ")
	}
	// tampilkan selisih waktu dengan LAMPLIFE
	fmt.Print("\t\t Sisa Waktu: ")
	fmt.Println(lamplife - time)
	fmt.Println(" --------------------------- \n")
}

func max(a, b int) int {
	/* mengembalikan nilai terbesar antara bilangan bulat a dan b*/
	if a > b {
		return a
	}
	return b
}

func move(left, right *island, b bridge, time *int) {
	/* I.S. terdefinisi tipe bentukan left dan right yang berisi left.nMember dan right.nMember orang, b berisi sepasang orang yang menyeberang jembatan, serta sebuah bilangan bulat time
	   F.S. left dan right terupdate sesuai dengan nilai pada b.m1 dan b.m2. Time terakumulasi dengan nilai terbesar antara b.m1 dan b.m2*/
	left.nama[left.nmember] = b.m1
	left.nama[left.nmember+1] = b.m2
	left.nmember = left.nmember + 2
	del_member(&*right, b.m1)
	del_member(&*right, b.m2)
	right.nmember = right.nmember - 2
	*time = *time + max(b.m1, b.m2)
}

func back(left, right *island, b bridge, time *int) {
	/* I.S. terdefinisi tipe bentukan left dan right yang berisi left.nMember dan right.nMember orang, b.m1 berisi orang yang kembali, serta sebuah bilangan bulat time
	   F.S. left dan right terupdate sesuai dengan nilai pada b.m1. Time terakumulasi dengan nilai b.m1*/
	right.nama[right.nmember] = b.m1
	right.nmember = right.nmember + 1
	del_member(&*left, b.m1)
	left.nmember = left.nmember - 1
	*time = *time + b.m1
}

func mulai_permainan(A *users) {
	/* I.S. terdefinisi tipe bentukan A yang berisi sejumlah A.nUser pemain pada array A.info
	   proses:
	   1. setup permainan
	   2. menerima masukan berupa username pemain baru.
	   3. berisi algoritma permaian bridge crossing hingga selesai
	   F.S. array A.info bertambah 1 data pemain, nilai A.nUser bertambah 1*/
	// a. deklrasi variabel string nama dan integer time (inisilisasi dengan 0) untuk menyimpan data pemain.
	var namapemain string
	var time int
	var b bridge
	time = 0
	// b. deklrasikan variabel untuk pulau kiri dan kanan (bertipe island) dan variabel b beripe bridge.
	var l, r island
	// c. lakukan setup permainan dengan memanggil procedure setup_permainan().
	setup_permainan(&l, &r)
	// d. panggil procedure header dan minta masukan nama pemain.
	header()
	fmt.Print("Nama pemain:")
	fmt.Scan(&namapemain)
	fmt.Println(" Mulai permainan... ")
	// e. buatlah perulangan while dengan kondisi selama di pulau kanan masih ada orang.
	for r.nmember != 0 {
		// f. tampilkan step saat ini dengan memanggil print_step.
		print_step(l, r, time)
		// g. minta masukan untuk pasangan orang yang akan menyeberang.
		fmt.Println(" Menyeberang (lihat di seberang kanan)...")
		fmt.Print("Orang 1: ")
		fmt.Scan(&b.m1)
		fmt.Print("Orang 2: ")
		fmt.Scan(&b.m2)
		// i. move() untuk pasangan orang yang akan menyeberang.
		move(&l, &r, b, &time)
		// j. tampilkan step saat ini dengan memanggil print_step.
		print_step(l, r, time)
		// k. jika masih terdapat orang di pulau kanan maka
		if r.nmember != 0 {
			fmt.Println(" Kembali...(lihat di seberang kiri)")
			// l. minta masukan untuk orang yang akan kembali mengantar lampu
			fmt.Print("Orang:")
			fmt.Scan(&b.m1)
			// m. back() untuk orang yang kembali mengantar lampu.
			back(&l, &r, b, &time)
		}
	}
	// n. tambahkan data pemain ke dalam A
	A.info[A.nuser].nama = namapemain
	A.info[A.nuser].time = time
	A.nuser++
	// o. Tampilkan waktu yang dihabiskan pemain "Selesai, dengan waktu: x"
	fmt.Println("Selesai dengan waktu: ", A.info[A.nuser-1].time)

}
