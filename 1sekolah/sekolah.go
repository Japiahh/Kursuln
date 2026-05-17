package main
import "fmt"
const nmax = 1001
type tabmurid[nmax] string
func main () {
	var n, i, l, m, h, idx  int
	var murid tabmurid
	var carimurid string
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&murid[i])
	}
	fmt.Scan(&carimurid)
	l = 0
	h = n - 1
	idx = -1
	for l <= h && idx == -1 {
		m = l + (h - l) / 2
		if murid[m] == carimurid {
			idx = m
		} else if murid[m] < carimurid {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	if idx != -1 {
		fmt.Printf("Murid terdaftar dan berada di urutan absen ke-%d\n", idx + 1)
	} else {
		fmt.Println("Murid tidak terdaftar.")
	}
}