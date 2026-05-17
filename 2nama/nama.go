package main
import "fmt"
const nmax = 9999
type arrstring [nmax] string
func bacadata (a *arrstring, n int) {
	var i int 
	for i = 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
}
func cetakdata (a *arrstring, n int) {
	var i int
	for i = 0; i < n; i++ {
		if i == n - 1 {
			fmt.Printf("%s", a[i])
		} else {
			fmt.Printf("%s, ", a[i])
		}
	}
	fmt.Printf(".\n\n")
}
func selectionsortascend (a *arrstring, n int) {
	var i, idx, pass int
	var temp string
	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if a[idx] > a[i] {
				idx = i
			}
			i = i + 1
		}
		temp = a[pass - 1]
		a[pass - 1] = a[idx]
		a[idx] = temp
		pass = pass + 1
	} 
	fmt.Printf("Data setelah diurutkan secara Ascending: \n")
	cetakdata(a, n)
}
func selectionsortdescend (a *arrstring, n int) {
	var i, idx, pass int
	var temp string
	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if a[idx] < a[i] {
				idx = i
			}
			i = i + 1
		}
		temp = a[pass - 1]
		a[pass - 1] = a[idx]
		a[idx] = temp
		pass = pass + 1
	}
	fmt.Printf("Data setelah diurutkan secara Descending: \n")
	cetakdata(a, n)
} 
func main () {
	var n int
	var a arrstring
	fmt.Scan(&n)
	bacadata(&a, n)
	selectionsortascend(&a, n)
	selectionsortdescend(&a, n)
}