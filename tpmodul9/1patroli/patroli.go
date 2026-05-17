package main
import "fmt"
type game struct {
	nama string
	populasi float64
	nilai int
}
type tabgame[20] game
func main () {
	var n, k, i, d int
	var tk float64
	var tab tabgame
	fmt.Scan(&n, &k)
	for i = 0; i < n; i++ {
		fmt.Scan(&tab[i].nama)
	}
	for i = 0; i < n; i++ {
		d = k / 10
		k = k - d
		tab[i].populasi = float64(d)
		tab[i].nilai = d % 1000
	}
	for i = 0; i < n; i++ {
		tk = float64(tab[i].nilai) * 0.025
		if tk > 5 {
			fmt.Printf("%s dengan tingkat kejahatan: %.2f \n", tab[i].nama, tk)
		}
	}
}