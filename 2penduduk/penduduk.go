package main
import "fmt"
const nmax = 2000
type penduduk struct {
	nama, kota string
	tahun int
}
type tabpenduduk[nmax] penduduk
func caridatapenduduk (data tabpenduduk, n int, caritahun int) int {
	var l, m, h, idx int
	l = 0
	h = n - 1
	idx = -1 
	for l <= h && idx == -1 {
		m = l + (h - l) / 2
		if data[m].tahun == caritahun {
			idx = m 
		} else if data[m].tahun < caritahun {
			l = m + 1 
		} else {
			h = m - 1 
		}
	}
	return idx 
}
func main() {
	var n, i, caritahun, hasilidx int
	var data tabpenduduk
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&data[i].nama, &data[i].tahun, &data[i].kota)
	}
	fmt.Scan(&caritahun)
	hasilidx = caridatapenduduk(data, n, caritahun)
	if hasilidx != -1 {
		fmt.Printf(data[hasilidx].nama, data[hasilidx].tahun, data[hasilidx].kota)
		fmt.Printf("ditemukan di index ke-%d\n", hasilidx)
	} else {
		fmt.Println("Data Tidak Ditemukan")
	}
}