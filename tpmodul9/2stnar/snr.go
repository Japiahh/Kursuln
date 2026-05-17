package main
import "fmt"
type item struct {
	nama string
	harga int
} 
type arr[999] item
func main () {
	var n, i, t, ax, in int
	var pmahal, pmurah string
	var it arr
	fmt.Scan(&n)
	t = 0
	for i = 0; i < n; i++{
		fmt.Print("Nama item: ")
		fmt.Scan(&it[i].nama)
		fmt.Print("Harga item: ")
		fmt.Scan(&it[i].harga)
		t = it[i].harga
		if i == 0 {
			ax = t
			in = t
			pmahal = it[i].nama
			pmurah = it[i].nama
		}
		if t > ax {
			ax = t
			pmahal = it[i].nama
		}
		if t < in {
			in = t
			pmurah = it[i].nama
		}
	}
	fmt.Printf("%s %s \n", "Nama Item", "Harga")
	for i = 0; i < n; i++ {
		fmt.Printf("%s %d \n", it[i].nama, it[i].harga)
	}
	fmt.Printf("Item termahal %s \n", pmahal)
	fmt.Printf("Item termahal %s ", pmurah)
}