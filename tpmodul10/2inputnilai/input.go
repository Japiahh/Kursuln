package main
import "fmt"
const nmax = 100
type student struct {
	nim, name string
	grade float64
}
type students[nmax-1] student
func main () {
	var n, i, j, k  int
	var d bool
	var a students
	i = 0
	fmt.Scan(&a[i].nim)
	for a[i].nim != "STOP" {
		fmt.Scan(&a[i].name, &a[i].grade)
		i = i + 1
		fmt.Scan(&a[i].nim)
	}
	n = i
	k = 0
	i = 0
	for i < n {
		d = false
		j = 0
		for j < k {
			if a[i].nim == a[j].nim {
				d = true
			}
			j = j + 1
		} 
		if !d {
			a[k] = a[i]
			k = k + 1
		}
		i = i + 1
	}
	n = k
	i = 0
	for i < n {
		fmt.Println(a[i].nim, a[i].name, a[i].grade)
		i = i + 1
	}
}