package main

import "fmt"

const NMAX int = 10

var idProyek int = 101

type event struct {
	id             int
	kategori       string
	nama           string
	jumlahDonasi   int
	targetDonasi   int
	deskripsiEvent [100]string
	jumlahOrang    int
	status         bool
}

type tabEvent [NMAX]event

func main() {
	var x, nData, cek int
	var event tabEvent
	for {
		menu()
		fmt.Scan(&x)
		switch x {
		case 1:
			buatEvent(&event, &nData)
			cek = 1
		case 2:
			listEvent(event, nData, cek)
		case 3:
			cekProyek(&event, &nData)
		case 4:
			donasi(&event, nData)
		}
		if x == 5 {
			fmt.Print("Sampai Jumpa Kembali ^_^")
			break

		}
	}
}

func menu() {
	fmt.Println("=============================================")
	fmt.Println()
	fmt.Println("           CROWDFUNDING PLUS PLUS            ")
	fmt.Println()
	fmt.Println("=============================================")
	fmt.Println("                     MENU                    ")
	fmt.Println("=============================================")
	fmt.Println("1. Buat Event")
	fmt.Println("2. List event")
	fmt.Println("3. Cek Event ")
	fmt.Println("4. Donasi	  ")
	fmt.Println("5. Exit")
	fmt.Print("Pilih (1/2/3/4/5) ? ")
	fmt.Println()
}

func buatEvent(A *tabEvent, n *int) {
	var z, i int

	if *n >= NMAX {
		*n = NMAX
		fmt.Println("Kuota Proyek Sudah Mencapai Batas Maksimum")
	}
	fmt.Println("====================================================================================================")
	fmt.Println(" Helo Selamat datang, silahkan isi data dibawah ini untuk membuat proyek baru yang akan di publish ")
	fmt.Println("====================================================================================================")

	A[*n].id = idProyek
	idProyek++
	fmt.Println("Silahkan pilih kategori projek anda :")
	fmt.Println("1. Teknologi")
	fmt.Println("2. Kesehatan")
	fmt.Println("3. Startup  ")
	fmt.Println("4. Kreatif  ")
	fmt.Print("(1/2/3/4) : ")
	fmt.Scan(&z)
	fmt.Println()
	if z == 1 {
		A[*n].kategori = "Teknologi"
	} else if z == 2 {
		A[*n].kategori = "Kesehatan"
	} else if z == 3 {
		A[*n].kategori = "Startup"
	} else if z == 4 {
		A[*n].kategori = "Kreatif"
	}

	fmt.Print("Silahkan Masukan Judul Event : ")
	fmt.Scan(&A[*n].nama)
	fmt.Print("Masukan Target Dana : ")
	fmt.Scan(&A[*n].targetDonasi)
	fmt.Println("Deskripsikan Proyek anda (masukan '.' untuk mengakhiri):")
	i = 0
	for {
		var word string
		fmt.Scan(&word)
		if word == "." {
			break
		}
		if i < 100 {
			A[*n].deskripsiEvent[i] = word
			i++
		} else {
			fmt.Println("Deskripsi terlalu panjang, dihentikan.")
			break
		}
	}
	fmt.Println("=============================")
	fmt.Println("Projek berhasil dipublish !!")
	fmt.Println("=============================")
	fmt.Println()
	*n = *n + 1
}

func cekProyek(event *tabEvent, n *int) {
	var cari int
	var namaDicari string
	fmt.Println("=======================================================================================")
	fmt.Println("Halo, kami menyediakan beberapa menu jika ingin mengubah data yang sudah di publish :")
	fmt.Println("=======================================================================================")
	fmt.Println("1. Edit Proyek")
	fmt.Println("2. Hapus Proyek")
	fmt.Println("3. Mencari Proyek")
	fmt.Print("opsi yang dipilih (1/2/3) ? ")
	fmt.Scan(&cari)
	fmt.Println()

	if cari == 1 {
		fmt.Print("Masukkan nama proyek yang ingin diedit: ")
		fmt.Scan(&namaDicari)
		editProyek(event, *n, namaDicari)
	} else if cari == 2 {
		fmt.Print("Masukkan nama proyek yang ingin dihapus: ")
		fmt.Scan(&namaDicari)
		hapusProyek(event, n, namaDicari)
	} else if cari == 3 {
		fmt.Print("Masukkan nama/kategori proyek yang ingin dicari: ")
		fmt.Scan(&namaDicari)
		mencariSemuaProyek(*event, *n, namaDicari)
	}
}

func sequentialSearchIndeks(data tabEvent, n int, x string) int {
	var idx, i int

	idx = -1
	for i < n && idx == -1 {
		if data[i].nama == x || data[i].kategori == x {
			idx = i
		}
		i++
	}
	return idx
}

func sequentialSearchStatus(A tabEvent, n int) {
	var idx, i, j int
	for i < n {
		if A[i].status == true {
			idx = i
			fmt.Println("=========================================================================================================")
			fmt.Printf("ID: %d || Judul: %s || Kategori: %s || Target Donasi: Rp %d || Total Donasi Terkumpul: Rp %d || Jumlah Donatur: %d", A[idx].id, A[idx].nama, A[idx].kategori, A[idx].targetDonasi, A[idx].jumlahDonasi, A[idx].jumlahOrang)
			fmt.Println("|| Deskripsi Proyek:")
			j = 0
			for j < 100 && A[i].deskripsiEvent[j] != "" && A[i].deskripsiEvent[j] != "." {
				fmt.Print(A[i].deskripsiEvent[j], " ")
				j++
			}
			fmt.Println()
			fmt.Println("=========================================================================================================")
		}
		i++
	}
}

func mencariSemuaProyek(data tabEvent, n int, x string) {
	var ketemu bool
	var i, j int

	ketemu = false
	i = 0

	for i < n {
		if data[i].nama == x || data[i].kategori == x {
			fmt.Println("============================")
			fmt.Printf("Proyek Ditemukan:\n")
			fmt.Printf("Id Proyek : %d\n", data[i].id)
			fmt.Printf("Judul: %s\n", data[i].nama)
			fmt.Printf("Kategori: %s\n", data[i].kategori)
			fmt.Printf("Target Donasi: Rp %d\n", data[i].targetDonasi)
			fmt.Printf("Total Donasi Terkumpul: Rp %d\n", data[i].jumlahDonasi)
			fmt.Printf("Jumlah Donatur: %d\n", data[i].jumlahOrang)
			fmt.Println("Deskripsi Proyek:")
			j = 0

			for j < 100 && data[i].deskripsiEvent[j] != "" && data[i].deskripsiEvent[j] != "." {
				fmt.Print(data[i].deskripsiEvent[j], " ")
				j++
			}
			fmt.Println()
			ketemu = true
		}
		i++
	}
	if !ketemu {
		fmt.Println("Proyek Tidak Ditemukan")
		fmt.Println()
	}
}

func hapusProyek(data *tabEvent, jumlah *int, x string) {
	var idx, i int

	idx = sequentialSearchIndeks(*data, *jumlah, x)
	if idx == -1 {
		fmt.Println("Proyek Tidak ditemukan")
	} else {
		i = idx
		for i <= *jumlah-2 {
			data[i] = data[i+1]
			i++
		}
		*jumlah = *jumlah - 1
	}
	fmt.Println("Proyek berhasil Dihapus")
}

func editProyek(data *tabEvent, jumlah int, x string) {
	var idx, i int

	idx = sequentialSearchIndeks(*data, jumlah, x)
	if idx == -1 {
		fmt.Println("Proyek Tidak Ditemukan")
	} else {
		fmt.Print("Masukkan Nama Proyek Baru: ")
		fmt.Scan(&data[idx].nama)
		fmt.Print("Masukkan Target Dana Baru: ")
		fmt.Scan(&data[idx].targetDonasi)
		fmt.Println("Masukkan Deskripsi Baru: ")
		fmt.Println("Deskripsikan Proyek anda (masukan '.' untuk mengakhiri):")
		i = 0
		for {
			var word string
			fmt.Scan(&word)
			if word == "." {
				break
			}
			if i < 100 {
				data[*&jumlah].deskripsiEvent[i] = word
				i++
			} else {
				fmt.Println("Deskripsi terlalu panjang, dihentikan.")
				break
			}
		}
		fmt.Println("Proyek berhasil Diganti")
	}

}

func donasi(A *tabEvent, n int) {
	var x string
	var idx, donasi int

	fmt.Println("=================================================================")
	fmt.Println("Haloo, silahkan masukan nama projek yang ingin di donasikan : ")
	fmt.Scan(&x)
	idx = sequentialSearchIndeks(*A, n, x)

	if idx == -1 {
		fmt.Print("Proyek Tidak ditemukan")
	} else {
		if A[idx].jumlahDonasi >= A[idx].targetDonasi {
			fmt.Print("Proyek ini sudah memenuhi target")
		} else {
			fmt.Printf("Silahkan Masukan Nominal yang akan di donasikan ke proyek %s  : Rp ", A[idx].nama)
			fmt.Scan(&donasi)
			A[idx].jumlahDonasi += donasi
			if A[idx].jumlahDonasi >= A[idx].targetDonasi {
				fmt.Printf("Terimakasih sudah melakukan donasi kepada Proyek %s sebesar %d, dengan ini kami dapat menyentuh target donasi yang dibutuhkan ^_^", A[idx].nama, donasi)
				A[idx].jumlahOrang += 1
				A[idx].status = true
				return
			}
			fmt.Printf("Terimakasih sudah melakukan donasi kepada Proyek %s sebesar %d", A[idx].nama, donasi)
			A[idx].jumlahOrang += 1

		}
	}
}

func listEvent(A tabEvent, n int, cek int) {
	var answer int
	if cek < 1 {
		fmt.Print("Proyek Tidak ditemukan")
	} else {
		fmt.Printf("List proyek:\n")
		cetakData(A, n)
	}
	fmt.Print()
	fmt.Println("1. Lihat proyek yang sudah mencapai target")
	fmt.Println("2. Urutkan list proyek berdasarkan jumlah donatur")
	fmt.Println("3. Kembali")
	fmt.Println("Pilih (1/2/3)")
	fmt.Scan(&answer)
	if answer == 1 {
		sequentialSearchStatus(A, n)
	} else if answer == 2 {
		insertionSort(&A, n)
		fmt.Print("List proyek setelah diurutkan:")
		cetakData(A, n)
	} else {
		return
	}
}

func cetakData(A tabEvent, n int) {
	var i, j int
	for i = 0; i < n; i++ {
		fmt.Println("=========================================================================================================")
		fmt.Printf("ID: %d || Judul: %s || Kategori: %s || Target Donasi: Rp %d || Total Donasi Terkumpul: Rp %d || Jumlah Donatur: %d", A[i].id, A[i].nama, A[i].kategori, A[i].targetDonasi, A[i].jumlahDonasi, A[i].jumlahOrang)
		fmt.Println("|| Deskripsi Proyek:")
		fmt.Printf("Deskripsi Proyek: ")
		j = 0
		for j < 100 && A[i].deskripsiEvent[j] != "" && A[i].deskripsiEvent[j] != "." {
			fmt.Print(A[i].deskripsiEvent[j], " ")
			j++
		}
		fmt.Println()
		fmt.Println("=========================================================================================================")
	}
}

func insertionSort(A *tabEvent, n int) {
	var i, pass int
	var temp event

	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.jumlahOrang < A[i-1].jumlahOrang {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
		pass = pass + 1
	}
}
