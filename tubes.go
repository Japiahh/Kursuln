package main
import "fmt"
const nmax = 100
type Peserta struct {
	ID int 
	nama string
	katalog string
	bidangminat string
	tanggalpendaftaran string
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
		if katal == 1 {
			d[i].tanggalpendaftaran = "7 - 11 September 2026"
		} else if katalog == 2 {
			d[i].tanggalpendaftaran = "14 - 18 September 2026"
		}  else if katalog == 3 {
			d[i].tanggalpendaftaran = "21 - 25 September 2026"
		}  else if katalog == 4 {
			d[i].tanggalpendaftaran = "28 September - 2 Oktober2026"
		}  else if katalog == 5 {
			d[i].tanggalpendaftaran = "5 - 9 Oktober 2026"
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
	acending(&d, 0, 1, &sel)
	fmt.Println("Data Diperbaharui!")
}

func lihatpeserta (banyakdatasaatini int, d tabregistrasi, sel *tabregistrasi) {
	var n string
	fmt.Printf("#Lihat Peserta \n \n")
	fmt.Println(" .../lihat_Peserta \n \n")
	lihatdatapeserta(banyakdatasaatini, d)
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
				fmt.Println("Data yang diurutkan (descending) : ")
				fmt.Printf("\n%-20s %-20s %-20s %-20s %-20s", "ID", "Nama", "katalog", "idang Minat", "Tanggal Pendaftaran")
				for i = 0; i < banyakdatasaatini; i++ {
					fmt.Printf("\n%-20d %-20s %-20s %-20s %-20s", d[i].ID, d[i].nama, d[i].katalog, d[i].bidangminat, d[i].tanggalpendaftaran)
				}
			} else if n == "up"{
				acending(d, 1, 0, &sel)
				fmt.Println("Data yang diurutkan (ascending) : ")
				fmt.Printf("\n%-20s %-20s %-20s %-20s %-20s", "ID", "Nama", "katalog", "idang Minat", "Tanggal Pendaftaran")
				for i = 0; i < banyakdatasaatini; i++ {
					fmt.Printf("\n%-20d %-20s %-20s %-20s %-20s", d[i].ID, d[i].nama, d[i].katalog, d[i].bidangminat, d[i].tanggalpendaftaran)
				}
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
				fmt.Println("Data yang diurutkan (descending) : ")
				fmt.Printf("\n%-20s %-20s %-20s %-20s %-20s", "ID", "Nama", "katalog", "idang Minat", "Tanggal Pendaftaran")
				for i = 0; i < banyakdatasaatini; i++ {
					fmt.Printf("\n%-20d %-20s %-20s %-20s %-20s", d[i].ID, d[i].nama, d[i].katalog, d[i].bidangminat, d[i].tanggalpendaftaran)
				}
			} else if n == "up"{
				acending(d, 0, 1, &sel)
				fmt.Println("Data yang diurutkan (ascending) : ")
				fmt.Printf("\n%-20s %-20s %-20s %-20s %-20s", "ID", "Nama", "katalog", "idang Minat", "Tanggal Pendaftaran")
				for i = 0; i < banyakdatasaatini; i++ {
					fmt.Printf("\n%-20d %-20s %-20s %-20s %-20s", d[i].ID, d[i].nama, d[i].katalog, d[i].bidangminat, d[i].tanggalpendaftaran)
				}
			}
		}
		if !(n == "bck" || n == "id" || n == "up") {
			fmt.Println("Pilihan tidak valid! Coba lagi...")
		}
	}
}
func lihatdatapeserta (banyakdatasaatini int, d tabregistrasi) {
	var i int
	fmt.Printf("\n%-20s %-20s %-20s %-20s %-20s", "ID", "Nama", "katalog", "idang Minat", "Tanggal Pendaftaran")
	for i = 0; i < banyakdatasaatini; i++ {
		fmt.Printf(" %-20d %-20s %-20s %-20s", d[i].ID, d[i].nama, d[i].katalog, d[i].bidangminat, d[i].tanggalpendaftaran)
	}
}
func acending(d tabregistrasi, id, nm, banyakdatasaatini int, sel *tabregistrasi) {
	var i, idx, pass int
	var temp Peserta
	var b, x bool
	if id == 1 {
		pass = 1
		for pass < banyakdatasaatini {
			temp = d[pass]
			i = pass
			for i > 0 && d[i - 1].ID > temp.ID {
				d[i] = d[i - 1]
				i = i - 1
			}
			d[i] = temp
			pass = pass + 1
		}
	}
	if nm == 1 {
		*sel = d 
		pass = 1
		for pass < banyakdatasaatini {
			idx = pass - 1
			i = pass
			for i < banyakdatasaatini {
				if sel[idx].nama > sel[i].nama { 
					idx = i
				}
				i = i + 1
			}
			temp = sel[pass-1]
			sel[pass-1] = sel[idx]
			sel[idx] = temp
			pass = pass + 1
		}
		d = *sel
	}
}
	
func decending(d tabregistrasi, id, nm int) { 
	var i, idx, pass, temp int
	var temp Peserta
	var b, x bool
	b = false
	x = false
	if id == 1 {
		pass = 1
			for pass < banyakdatasaatini {
				idx = pass - 1
				i = pass
				for i < banyakdatasaatini {
					if d[idx].ID < d[i].ID { 
						idx = i
					}
					i = i + 1
				}
				temp = d[pass - 1] 
				d[pass - 1] = d[idx]
				d[idx] = temp1
				
				pass = pass + 1
			}
			b = true
	}
	if nm == 1 {
		if id == 1 {
			pass = 1
			for pass < banyakdatasaatini {
				temp = d[pass]
				i = pass
				for i > 0 && d[i - 1].nama < temp.nama {
					d[i] = d[i - 1]
					i = i - 1
				}
				d[i] = temp
				pass = pass + 1
			}
		}
	}
}
func editpeserta(banyakdatasaatini int, d *tabregistrasi) {
	var k int
	var n string
	fmt.Println("Cari data nama lengkap : ")
	fmt.Println()
	lihatdatapeserta(banyakdatasaatini, d)
	fmt.Print("\n \nInput[Nama Lengkap]> ")
	fmt.Scan(&n)
	namal = carinamalengkap(sel, banyakdatasaatini, n)
	if data == -1 {
		fmt.Printf("Data Tidak Ditemukan! Coba lagi atau keluar...[y/any]\n \n")
		fmt.Print("Input[y/any]> ")
		fmt.Scan(&n)
		if n == "y" {
			editpeserta(banyakdatasaatini, d)
		}
	} else {
		fmt.Printf("\n \n Data ditemukan : \n \n")
		fmt.Printf("\n%-20s %-20s %-20s %-20s %-20s", "ID", "Nama", "katalog", "Bidang Minat", "Tanggal Pendaftaran")
		fmt.Printf("\n%-20d %-20s %-20s %-20s %-20s \n \n", d[namal].ID, d[namal].nama, d[namal].katalog, d[namal].bidangminat, d[namal].tanggalpendaftaran)
		fmt.Print("confirm[n/any]> ")
		fmt.Scan(&n)
		if n == "n" {
			editpeserta(banyakdatasaatini, d)
		} else {
			tabelpeserta(3, 0)
			fmt.Println()
			fmt.Print("input[nm/kt/bm]> ")
			fmt.Scan(&n)
			for n != "ok" {
				if n = "nm" {
					fmt.Printf("input[Data Nama lengkap baru] \n> ")
					fmt.Scan(&n)
					d[namal].nama = n
					fmt.Print("confirm[ok/kt/bm]> ")
					fmt.Scan(&n)
				} else if n = "kt" {
					tabelbidangminatdankatalog(0)
					fmt.Printf("input[Data Katalog baru] \n> ")
					fmt.Scan(&k)
					compi = comp(k, 0)
					d[namal].katalog = compi
					fmt.Print("confirm[ok/kt/bm]> ")
					fmt.Scan(&n)
				} else if n = "bm" {
					tabelbidangminatdankatalog(1)
					fmt.Printf("input[Data Bidang Minat baru] \n> ")
					fmt.Scan(&k)
					compi = comp(0, k)
					d[namal].bidangminat = compi
					fmt.Print("confirm[ok/kt/bm]> ")
					fmt.Scan(&n)
				} else if n == "ok" {

				} else {
					fmt.Println("Pilihan tidak valid! Coba lagi...")
					fmt.Print("confirm[nm/kt/bm/ok]> ")
					fmt.Scan(&n)
				}
			}
		}
	}
	acending(d, 0, 1, banyakdatasaatini, &sel)
}
func hapuspeserta (banyakdatasaatini *int, d, sel *tabregistrasi, totalhapus *int) {
	var n string
	var i, asel, ad int
	fmt.Println("Cari data nama lengkap : ")
	fmt.Println()
	lihatdatapeserta(banyakdatasaatini, d)
	fmt.Print("\n \nInput[Nama Lengkap]> ")
	fmt.Scan(&n)
	asel = carinamalengkap(*sel, *banyakdatasaatini, n)
	if asel == -1 {
		fmt.Printf("Data Tidak Ditemukan! Coba lagi atau keluar...[y/any]\n \n")
		fmt.Print("Input[y/any]> ")
		fmt.Scan(&n)
		if n == "y" {
			hapuspeserta(banyakdatasaatini, d, sel)
		}
	} else {
		fmt.Printf("\n\n Data ditemukan : \n\n")
		fmt.Printf("\n%-20s %-20s %-20s %-20s %-20s", "ID", "Nama", "katalog", "Bidang Minat", "Tanggal Pendaftaran")
		fmt.Printf("\n%-20d %-20s %-20s %-20s %-20s \n \n", sel[asel].ID, sel[asel].nama, sel[asel].katalog, sel[asel].bidangminat, sel[asel].tanggalpendaftaran)
		fmt.Printf("confirm[y/any]> ")
		fmt.Scan(&n)
		if n == "y" {
			ad = -1
			for i = 0; i < *banyakdatasaatini; i++ {
				if d[i].ID == sel[asell].ID {
					ad = i
					i = *banyakdatasaatini 
				}
			}
			if ad != -1 {
				for i = ad; i < *banyakdatasaatini - 1; i++ {
					d[i] = d[i + 1]
				}
				for i = asel; i < *banyakdatasaatini - 1; i++ {
					sel[i] = sel[i + 1]
				}
				*banyakdatasaatini = *banyakdatasaatini - 1
				fmt.Println("Data berhasil dihapus!")
			}
		} else {
			fmt.Println("Penghapusan dibatalkan.")
		}
	}
	totalhapus = totalhapus + 1
}

func pengaturanpeserta(d *tabregistrasi, banyakdatasaatini *int,) {
	var pilih int
	for pilih != 5 {
		tabelpeserta(0,0)
		fmt.Print("input[1/2/3/4/5]> ")
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
func caripeserta(d, sel tabregistrasi, banyakdatasaatini int) {
	var data int
	var pilih, pc string
	var v bool
	data = -1
	tabelmain(2, 0)
	fmt.Print("input[nm/bm]> ")
	fmt.Scan(&pilih)
	if pilih == "nm" {
		fmt.Println("Ketik Nama Lengkap[sesuai data]> ")
		fmt.Scan(&pilih)
		data = carinamalengkap(sel, banyakdatasaatini, pilih)
		if data == -1 {
		fmt.Printf("Data Tidak Ditemukan! \n \n")
		} else {
			fmt.Printf("\n \n Data ditemukan : \n \n")
			fmt.Printf("\n%-20s %-20s %-20s %-20s %-20s", "ID", "Nama", "katalog", "idang Minat", "Tanggal Pendaftaran")
			fmt.Printf("\n%-20d %-20s %-20s %-20s %-20s", d[data].ID, d[data].nama, d[data].katalog, d[data].bidangminat, d[data].tanggalpendaftaran)
		}
	} else if pilih == "bm" {
		tabelmain(2, 1)
		fmt.Println("input>")
		fmt.Scan(&pilih)
		for pilih != 1 <= 5 {
			fmt.Println("Pilihan tidak valid! Coba lagi...")
			fmt.Println()
			tabelmain(2, 1)
			fmt.Println("input>")
			fmt.Scan(&pilih)
		}  
		for i = 0; i < banyakdatasaatini; i++ {
			if d[i].bidangminat == pilih {
				if ketemu == false {
					fmt.Printf("\n \n Data ditemukan : \n \n")
					fmt.Printf("\n%-20s %-20s %-20s %-20s %-20s", "ID", "Nama", "katalog", "Bidang Minat", "Tanggal Pendaftaran")
					ketemu = true
				}
				fmt.Printf("\n%-20d %-20s %-20s %-20s %-20s", d[i].ID, d[i].nama, d[i].katalog, d[i].bidangminat, d[i].tanggalpendaftaran)
			}
		}
		if ketemu == false {
			fmt.Printf("Data Tidak Ditemukan! \n \n")
		}
	}
	fmt.Print("confirm[any]> ")
	fmt.Scan(&pilih)
}
func carinamalengkap(sel tabregistrasi, banyakdatasaatini int, namalengkap string) int { 
	var l, m, h, idx int
	l = 0
	h = banyakdatasaatini - 1
	idx = -1 
	for l <= h && idx == -1 {
		m = l + (h - l) / 2
		if sel[m].nama == namalengkap {
			idx = m 
			return idx
		} else if sel[m].tahun < sel {
			l = m + 1 
		} else {
			h = m - 1 
		}
	}
}

func ringkasanstatistik(d tabregistrasi, banyakdatasaatini, totalhapus int) {
	lihatdatapeserta(banyakdatasaatini, d)
	fmt.Println()
	fmt.Println("Statistik Data : ")
	for i = 0; i < banyakdatasaatini; i++ {
		if d[i].bidangminat == "Pemograman Web" {
			pw = pw + 1
		} else if d[i].bidangminat == "Data Sains" {
			ds = ds + 1
		} else if d[i].bidangminat == "Desain UI/UX" {
			du = du + 1
		} else if d[i].bidangminat == "Keamanan Cyber" {
			kc = kc + 1
		} else if d[i].bidangminat == "Mobile Development" {
			md = md + 1
		}
		if d[i].bidangminat == "Teknologi" {
			t = t + 1
		} else if d[i].bidangminat == "Bisnis Digital" {
			bd = bd + 1
		} else if d[i].bidangminat == "Desain Kreatif" {
			dk = dk + 1
		} else if d[i].bidangminat == "Kecerdasan Buatan" {
			kb = kb + 1
		} else if d[i].bidangminat == "Keamanan Sistem" {
			ks = ks + 1
		}
	}
	fmt.Printf("Total data : %d", banyakdatasaatini)
	fmt.Println()
	fmt.Println("Data Katalog : ")
	fmt.Printf("| %-20s | %-20s | %-20s |", "Nama Data", "Jumlah", "Persentase")
	fmt.Printf("| %-20s | %-20d | %-20f% |", "Pemograman Web", pw, float64((pw/banyakdatasaatini)*100))
	fmt.Printf("| %-20s | %-20d | %-20f% |", "Data Sains", ds, float64((ds/banyakdatasaatini)*100))
	fmt.Printf("| %-20s | %-20d | %-20f% |", "Desain UI/UX", du, float64((du/banyakdatasaatini)*100))
	fmt.Printf("| %-20s | %-20d | %-20f% |", "Keamanan Cyber", kc, float64((kc/banyakdatasaatini)*100))
	fmt.Printf("| %-20s | %-20d | %-20f% |", "Mobile Develoment", md, float64((md/banyakdatasaatini)*100))
	fmt.Println()
	fmt.Println("Data Bidang Minat : ")
	fmt.Printf("| %-20s | %-20s | %-20s |", "Nama Data", "Jumlah", "Persentase")
	fmt.Printf("| %-20s | %-20d | %-20f% |", "Teknologi", t, float64((t/banyakdatasaatini)*100))
	fmt.Printf("| %-20s | %-20d | %-20f% |", "Bisnis Digital", bd, float64((bd/banyakdatasaatini)*100))
	fmt.Printf("| %-20s | %-20d | %-20f% |", "Desain Kreatif", dk, float64((dk/banyakdatasaatini)*100))
	fmt.Printf("| %-20s | %-20d | %-20f% |", "Kecerdasan Buatan", kb, float64((kb/banyakdatasaatini)*100))
	fmt.Printf("| %-20s | %-20d | %-20f% |", "Keamanan Sistem", ks, float64((ks/banyakdatasaatini)*100))
	fmt.Println()
	fmt.Println("Informasi Tambahan : ")
	fmt.Printf("Total hapus : %d", totalhapus)
}

func tabelmain(tm, tms int) {
	if tm == 0 {
		fmt.Printf("KURSUSLN (Sistem Pendaftaran Kursus Online Terpadu) \n \n")
		fmt.Printf(" /main \n \n")
		fmt.Println("Pilihan :")
		fmt.Println(" 1 = Pengaturan Perserta")
		fmt.Println(" 2 = Cari Data Perserta")
		fmt.Println(" 3 = Lihat Peserta")
		fmt.Println(" 4 = Ringkasan Statistik")
		fmt.Println(" 5 = Selesai")
	} else if tm == 2 {
		fmt.Printf("#Cari Peserta \n \n")
		fmt.Printf(" main/Cari_Peserta \n \n")
		fmt.Println("Pilihan :")
		fmt.Println(" nm = Nama")
		fmt.Println(" bm = Bidang Minat")
	} else if tm == 2 && tms == 1 {
		fmt.Println("Cari Bidang Minat :")
		tabelbidangminatdankatalog(1)
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
		tabelbidangminatdankatalog(0)
	} else if tpe == 1 && tpi == 2 {
		fmt.Printf("#Tambah Peserta \n \n")
		fmt.Printf(" .../Peserta/Tambah_Peserta-Pilihan_Bidang_Minat \n \n")
		fmt.Println("Pilihan Bidang Minat :")
		tabelbidangminatdankatalog(1)
	} else if tpe == 2 {
		fmt.Printf("#Lihat Peserta \n \n")
		fmt.Printf(" .../Peserta/Lihat_Peserta \n \n")
		fmt.Println("Cara Lihat :")
		fmt.Println(" id[up/dwn] = Berdasarkan ID (Up/Down)")
		fmt.Println(" nm[up/dwn] = Berdasarkan Nama (Up/Down)")
		fmt.Println(" bck = kembali")
	} else if tpe == 3 {
		fmt.Printf("#Daftar Edit Peserta \n \n")
		fmt.Printf(" .../Peserta/Edit_Peserta/Daftar_Edit_Peserta \n \n")
		fmt.Println("nm = Nama")
		fmt.Println("kt = Katalog")
		fmt.Println("bm = Bidang Minat")
		fmt.Println("ok = konfirmasi") 
	} 
}
func tabelbidangminatdankatalog (tbk) {
	if tbk == 0 {
		fmt.Println(" 1 = Pemograman Web")
		fmt.Println(" 2 = Data Sains")
		fmt.Println(" 3 = Desain UI/UX")
		fmt.Println(" 4 = Keamanan Cyber")
		fmt.Println(" 5 = Mobile Development")
	} else if tbk == 1 {
		fmt.Println(" 1 = Teknologi")
		fmt.Println(" 2 = Bisnis Digital")
		fmt.Println(" 3 = Kecerdasan Buatan")
		fmt.Println(" 4 = Desain Kreatif")
		fmt.Println(" 5 = Keamanan Sistem")
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
	var pilih2 string
	var target string
	var d, sel tabregistrasi
	var banyakdatasaatini int
	for pilih != 5 {
		tabelmain(0)
		fmt.Print("input[1/2/3/4]> ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			pengaturanpeserta(&d, &banyakdatasaatini)
		} else if pilih == 2 {
			caripeserta(d, banyakdatasaatini)
		} else if pilih == 3 {
			lihatpeserta(d, banyakdatasaatini)
		} else if pilih == 4 {
			ringkasanstatistik(d, banyakdatasaatini, totalhapus)
		} else if pilih == 5 {
			fmt.Println("Program selesai.")
		} else {
			fmt.Println("Pilihan tidak valid! Coba lagi...")
		}
	}
} 
