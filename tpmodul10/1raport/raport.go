package main
import "fmt"
const nmax = 999
type rapot struct {
	matkul string
	nilai int
}
type tabraport[nmax-1] rapot
func main () {
	var n, i, s int
	var a tabraport
	var k bool
	fmt.Scan(&n)
	i = 0
	k = false
	for i < n {
		fmt.Scan(&a[i].matkul, &a[i].nilai)
		i = i + 1 
	}
	fmt.Scan(&s)
	i = 0
	for i < n {
		if a[i].nilai == s {
			fmt.Println(a[i].matkul, a[i].nilai)
			k = true
		}
		i = i + 1
	}
	if k == true {
		fmt.Println("Data ditemukan!")
	} else {
		fmt.Println("Data tidak ditemukan!")
	}
}