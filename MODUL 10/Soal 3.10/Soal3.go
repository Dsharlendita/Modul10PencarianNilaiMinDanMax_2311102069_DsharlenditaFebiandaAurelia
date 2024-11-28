package main

import "fmt"

type daftarBerat [100]float64

func cariMinMax(berat daftarBerat, min, max *float64) {
	if len(berat) == 0 {
		return
	}

	*min = berat[0]
	*max = berat[0]

	for _, b := range berat {
		if b < *min {
			*min = b
		}
		if b > *max {
			*max = b
		}
	}
}

func hitungRataRata(berat daftarBerat, jumlahData int) float64 {
	var total float64
	for i := 0; i < jumlahData; i++ {
		total += berat[i]
	}
	return total / float64(jumlahData)
}

func main() {
	var dataBerat daftarBerat
	var jumlahData int

	fmt.Print("Masukkan jumlah data berat balita: ")
	fmt.Scan(&jumlahData)

	if jumlahData <= 0 || jumlahData > 100 {
		fmt.Println("Jumlah data berat balita harus antara 1 dan 100.")
		return
	}

	for i := 0; i < jumlahData; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBerat[i])
	}

	var beratMin, beratMax float64
	cariMinMax(dataBerat, &beratMin, &beratMax)

	fmt.Printf("Berat balita terkecil: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita terbesar: %.2f kg\n", beratMax)

	rataRata := hitungRataRata(dataBerat, jumlahData)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
}
