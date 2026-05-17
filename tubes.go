package main
import "fmt"
const nmax = 100
type Peserta struct {
	nama string
	ID int 
	katalog string
	bidangminat string
}
type tabregistrasi [nmax] Peserta
func tambahpeserta(banyakdatasaatini *int, d *tabregistrasi){
	var i, h, katal, bidangm int
	fmt.Printf("#Tambah Peserta \n \n")
	fmt.Println(" .../Peserta/Tambah_Peserta \n \n")
	fmt.Println("Banyak Peserta :")
	fmt.Print("input> ")
	fmt.Scan(&n)
	h = 0
	for i = *banyakdatasaatini; i < n; i++ {
		h = h + 1 
		fmt.Printf("Perserta ke-%d", h)
		fmt.Print("Nama> ")
		fmt.Scan(&d[i].nama)
		
		tabelpeserta(1, 1)
		fmt.Print("Katalog> ")
		fmt.Scan(&katal)
		if katal < 10 && katal > 0 {
			d[i].katalog = comp(katal, 0)
		} else {
			for !(katal < 10 && katal > 0) {
				fmt.Println("Pilihan tidak valid! Coba lagi..")
				tabelpilihankatalog()
				fmt.Scan(&katal)
			}
			d[i].katalog = comp(katal, 0)
		}
		
		tabelpeserta(1, 2)
		fmt.Scan(&bidangm)
		if bidangm < 10 && bidangm > 0 {
			d[i].bidangm = comp(0, bidangm)
		} else {
			for !(bidangm < 10 && bidangm > 0) {
				fmt.Println("Pilihan tidak valid! Coba lagi...")
				tabelpilihanbidangminat()
				fmt.Scan(&bidangm)
			}
			d[i].bidangm = comp(0, bidangm)
		}
		
		d[i].ID = (10000 * katal) + (1000 * bidangm) + i	
	}
	*banyakdatasaatini = *banyakdatasaatini + n
	acending(&d, 0, 0, &sel)
	fmt.Println("Data Diperbaharui!")
}

func lihatpeserta (banyakdatasaatini int, d tabregistrasi, sel *tabregistrasi) {
	var i int
	var n string
	fmt.Printf(" %-20s %-20s %-20s %-20s", "ID", "Nama", "Katalog", "Bidang Minat")
	for i = 0; i < banyakdatasaatini; i++ {
		fmt.Printf(" %-20d %-20s %-20s %-20s", d[i].ID, d[i].nama, d[i].katalog, d[i].bidangminat)
	}
	fmt.Println()
	for n != "bck" {
		tabelpeserta(2, 0)
		fmt.Print("input[id/nm/back]> ")
		fmt.Scan(&n)
		if n == "id" {
			fmt.Print("input[up/dn]> ")
			fmt.Scan(&n)
			for n != "up" && n != "dn" {
				fmt.Println("Pilihan tidak valid! Coba lagi...")
				fmt.Print("input[id/nm/back]> ")
				fmt.Scan(&n)
			}
			if n == "dn" {
				decending(d, 1, 0)
			} else if n == "up"{
				acending(d, 1, 0, &sel)
			}
		} else if n == "nm" {
			fmt.Print("input[up/dn]> ")
			fmt.Scan(&n)
			for n != "up" && n != "dn" {
				fmt.Println("Pilihan tidak valid! Coba lagi...")
				fmt.Print("input[up/dn]> ")
				fmt.Scan(&n)
			}
			if n == "dn" {
				decending(d, 0, 1)
			} else if n == "up"{
				acending(d, 0, 1, &sel)
			}
		}
		if !(n == "bck" || n == "id" || n == "up") {
			fmt.Println("Pilihan tidak valid! Coba lagi...")
		}
	}
}
func acending(d tabregistrasi, id, nm int, sel *tabregistrasi) {
	var 
	if id == 1 {
		
	}
	if nm == 1 {
		
	}
}
func decending(d tabregistrasi, id, nm int) { 
	if id == 1 {
		
	}
	if nm == 1 {
		
	}
}
func editpeserta(banyakdatasaatini int, d *tabregistrasi) {
	var i int
}
func hapuspeserta () {
	
}

func pengaturanpeserta(d *tabregistrasi, banyakdatasaatini *int,) {
	var pilih int
	for pilih != 5 {
		tabelpeserta(0,0)
		fmt.Print("input> ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			tambahpeserta(&banyakdatasaatini, &d)
		} else if pilih == 2 {
			lihatpeserta(banyakdatasaatini, d)
		} else if pilih == 3 {
			editpeserta(banyakdatasaatini, &d)
		} else if pilih == 4 {
			hapuspeserta(&banyakdatasaatini, &d)
		} else if pilih == 5 {
			fmt.Printf("Kembali ke Main... \n \n")
		} else {
			fmt.Println("Pilihan tidak valid! Coba lagi...")
		}
	}
}
func caripeserta(d tabregistrasi, banyakdatasaatini int) int {
	var n string
	
	
}
func caribidangminat(d tabregistrasi, banyakdatasaatini int){
	var i int
}
func carinamalengkap() { /*sel*/
	
}

func ringkasanstatistik() {
	
}

func tabelmain(tm int) {
	if tm == 0 {
		fmt.Printf("KURSUSLN (Sistem Pendaftaran Kursus Online Terpadu) \n \n")
		fmt.Printf(" /main \n \n")
		fmt.Println("Pilihan :")
		fmt.Println(" 1 = Pengaturan Perserta")
		fmt.Println(" 2 = Cari Data Perserta")
		fmt.Println(" 3 = Lihat Peserta")
		fmt.Println(" 4 = Ringkasan Statistik")
		fmt.Println(" 5 = ")
	}
}
func tabelpeserta(tpe, tpi int) {
	if tpe == 0 {
		fmt.Printf("#Pengaturan Peserta \n \n")
		fmt.Printf(" .../Peserta \n \n")
		fmt.Println("Pilihan :")
		fmt.Println(" 1 = Tambah Peserta")
		fmt.Println(" 2 = Lihat Peserta")
		fmt.Println(" 3 = Edit Peserta")
		fmt.Println(" 4 = Hapus Peserta")
		fmt.Println(" 5 = Kembali")
	} else if tpe == 1 && tpi == 1 {
		fmt.Printf("#Tambah Peserta \n \n")
		fmt.Printf(" .../Peserta/Tambah_Peserta-Pilihan_Katalog \n \n")
		fmt.Println("Pilihan Katalog Kursus :")
		fmt.Println(" 1 = ")
		fmt.Println(" 2 = ")
		fmt.Println(" 3 = ")
		fmt.Println(" 4 = ")
		fmt.Println(" 5 = ")
	} else if tpe == 1 && tpi == 2 {
		fmt.Printf("#Tambah Peserta \n \n")
		fmt.Printf(" .../Peserta/Tambah_Peserta-Pilihan_Bidang_Minat \n \n")
		fmt.Println("Pilihan Katalog Kursus :")
		fmt.Println(" 1 = ")
		fmt.Println(" 2 = ")
		fmt.Println(" 3 = ")
		fmt.Println(" 4 = ")
		fmt.Println(" 5 = ")
	} else if tpe == 2 {
		fmt.Printf("#Lihat Peserta \n \n")
		fmt.Printf(" .../Peserta/Lihat_Peserta \n \n")
		fmt.Println("Cara Lihat :")
		fmt.Println(" id[up/dwn] = Berdasarkan ID (Up/Down)")
		fmt.Println(" nm[up/dwn] = Berdasarkan Nama (Up/Down)")
		fmt.Println(" bck = kembali")
	} else if tpe == 3 {
		fmt.Printf("#Edit Peserta \n \n")
		fmt.Printf(" .../Peserta/Edit_Peserta \n \n")
		fmt.Println()
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
	} else if tpe == 4 {
		fmt.Printf("#Hapus Peserta \n \n")
		fmt.Printf(" .../Peserta/Hapus_Peserta \n \n")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
	}
}

func comp(katalog, bidangminat int) string {
	if katalog == 1 {
		return "Pemograman Web"
	} else if katalog == 2 {
		return "Data Sains"
	} else if katalog == 3 {
		return "Desain UI/UX"
	} else if katalog == 4 {
		return "Keamanan Cyber"
	} else if katalog == 5 {
		return "Mobile Development"
	}
	
	if bidangminat == 1 {
		return "Teknologi"
	} else if bidangminat == 2 {
		return "Bisnis Digital"
	} else if bidangminat == 3 {
		return "Desain Kreatif"
	} else if bidangminat == 4 {
		return "Kecerdasan Buatan"
	} else if bidangminat == 5 {
		return "Keamanan Sistem"
	}
	return
}

func main() {
	var pilih int
	var target string
	var d, sel tabregistrasi
	var banyakdatasaatini int
	for pilih != 10 {
		tabelmain(0)
		fmt.Print("input> ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			pengaturanpeserta(&d, &banyakdatasaatini)
		} else if pilih == 2 {
			caripeserta(&d, &banyakdatasaatini)
		} else if pilih == 3 {
			lihatpeserta(&d, &banyakdatasaatini)
		} else if pilih == 4 {
			
		}
	}
} 