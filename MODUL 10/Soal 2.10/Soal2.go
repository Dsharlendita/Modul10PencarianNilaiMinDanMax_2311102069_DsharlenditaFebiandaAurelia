package main

import "fmt"

func main() {
	var totalIkan, ikanPerWadah int
	fmt.Print("Masukkan total jumlah ikan: ")
	fmt.Scan(&totalIkan)
	fmt.Print("Masukkan jumlah ikan per wadah: ")
	fmt.Scan(&ikanPerWadah)

	if totalIkan <= 0 || ikanPerWadah <= 0 || totalIkan > 1000 {
		fmt.Println("Jumlah ikan dan jumlah ikan per wadah harus positif, dan total ikan tidak boleh lebih dari 1000.")
		return
	}

	beratIkan := make([]float64, totalIkan)

	fmt.Println("Masukkan berat masing-masing ikan:")
	for i := 0; i < totalIkan; i++ {
		fmt.Scan(&beratIkan[i])
	}

	jumlahWadah := (totalIkan + ikanPerWadah - 1) / ikanPerWadah 
	beratDiWadah := make([]float64, jumlahWadah)

	for i := 0; i < totalIkan; i++ {
		indeksWadah := i / ikanPerWadah 
		beratDiWadah[indeksWadah] += beratIkan[i]
	}

	fmt.Println("Total berat ikan di setiap wadah:")
	for _, berat := range beratDiWadah {
		fmt.Printf("%.2f kg ", berat)
	}
	fmt.Println()

	var totalBerat float64
	for _, berat := range beratDiWadah {
		totalBerat += berat
	}
	beratRataRata := totalBerat / float64(jumlahWadah)

	fmt.Printf("Berat rata-rata ikan per wadah: %.2f kg\n", beratRataRata)
}
