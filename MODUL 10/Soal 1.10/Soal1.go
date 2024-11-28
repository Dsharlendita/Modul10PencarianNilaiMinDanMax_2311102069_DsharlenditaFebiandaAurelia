package main

import "fmt"

func main() {
	var jumlahKelinci int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&jumlahKelinci)

	if jumlahKelinci <= 0 || jumlahKelinci > 1000 {
		fmt.Println("Jumlah anak kelinci harus berada antara 1 dan 1000.")
		return
	}

	beratKelinci := make([]float64, jumlahKelinci)

	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < jumlahKelinci; i++ {
		fmt.Scan(&beratKelinci[i])
	}

	beratTerendah := beratKelinci[0]
	beratTertinggi := beratKelinci[0]

	for i := 1; i < jumlahKelinci; i++ {
		if beratKelinci[i] < beratTerendah {
			beratTerendah = beratKelinci[i]
		}
		if beratKelinci[i] > beratTertinggi {
			beratTertinggi = beratKelinci[i]
		}
	}

	fmt.Printf("Berat terendah: %.2f kg\n", beratTerendah)
	fmt.Printf("Berat tertinggi: %.2f kg\n", beratTertinggi)
}
