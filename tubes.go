package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
)

type Komponen struct {
	SerialNumber int
	Name         string
	Status       string
	Temperature  float64
}

var dataPC []Komponen

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func konfirmasiMenu() bool {
	var pilihan string
	for {
		fmt.Print("\nKembali ke menu utama? (y/n): ")
		fmt.Scan(&pilihan)
		if pilihan == "y" || pilihan == "Y" {
			return true
			} else if pilihan == "n" || pilihan == "N" {
				return false
			}
		fmt.Println("Input tidak valid. Silakan masukkan 'y' atau 'n'.")
	}
}

func tentukanStatus(suhu float64) string {
	if suhu < 70 {
		return "Normal"
		} else if suhu >= 70 && suhu < 90 {
			return "Lag"
		}
		return "Overheat"
}
	
func tambah() {
		var k Komponen
		fmt.Print("Masukkan Serial Number: ")
		fmt.Scan(&k.SerialNumber)
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&k.Name)
		fmt.Print("Masukkan Suhu: ")
		fmt.Scan(&k.Temperature)
		
		k.Status = tentukanStatus(k.Temperature)
		
		dataPC = append(dataPC, k)
		fmt.Printf("\nData berhasil ditambahkan dengan status: %s!\n", k.Status)
	}
	
func ubah() {
		var sn int
		fmt.Print("Masukkan Serial Number yang ingin diubah: ")
		fmt.Scan(&sn)
		for i, k := range dataPC {
			if k.SerialNumber == sn {
				fmt.Print("Masukkan Nama Baru: ")
			fmt.Scan(&dataPC[i].Name)
			fmt.Print("Masukkan Suhu Baru: ")
			fmt.Scan(&dataPC[i].Temperature)
			
			dataPC[i].Status = tentukanStatus(dataPC[i].Temperature)
			
			fmt.Printf("\nData berhasil diubah! Status sekarang: %s\n", dataPC[i].Status)
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func hapus() {
	var sn int
	fmt.Print("Masukkan Serial Number yang ingin dihapus: ")
	fmt.Scan(&sn)
	for i, k := range dataPC {
		if k.SerialNumber == sn {
			dataPC = append(dataPC[:i], dataPC[i+1:]...)
			fmt.Println("\nData berhasil dihapus.")
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func cariSequential() {
	var name string
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scan(&name)
	for _, k := range dataPC {
		if k.Name == name {
			fmt.Printf("Ditemukan: SN=%d, Status=%s, Suhu=%.1f\n", k.SerialNumber, k.Status, k.Temperature)
			return
		}
	}
	fmt.Println("Tidak ditemukan.")
}

func cariBinary() {
	var name string
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scan(&name)
	low, high := 0, len(dataPC)-1
	for low <= high {
		mid := (low + high) / 2
		if dataPC[mid].Name == name {
			fmt.Printf("Ditemukan: SN=%d, Status=%s, Suhu=%.1f\n", dataPC[mid].SerialNumber, dataPC[mid].Status, dataPC[mid].Temperature)
			return
			} else if dataPC[mid].Name < name {
				low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Tidak ditemukan.")
}

func selectionSort() {
	n := len(dataPC)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if dataPC[j].SerialNumber < dataPC[min].SerialNumber {
				min = j
			}
		}
		dataPC[i], dataPC[min] = dataPC[min], dataPC[i]
	}
	fmt.Println("Data berhasil diurutkan berdasarkan Serial Number (Selection Sort).")
}

func insertionSort() {
	n := len(dataPC)
	for i := 1; i < n; i++ {
		key := dataPC[i]
		j := i - 1
		for j >= 0 && dataPC[j].SerialNumber > key.SerialNumber {
			dataPC[j+1] = dataPC[j]
			j--
		}
		dataPC[j+1] = key
	}
	fmt.Println("Data berhasil diurutkan berdasarkan Serial Number (Insertion Sort).")
}

func statistik() {
	if len(dataPC) == 0 {
		fmt.Println("Data kosong.")
		return
	}
	masalah := 0
	totalSuhu := 0.0
	for _, k := range dataPC {
		if k.Status != "Normal" {
			masalah++
		}
		totalSuhu += k.Temperature
	}
	fmt.Printf("Jumlah komponen bermasalah (Lag/Overheat): %d\n", masalah)
	fmt.Printf("Rata-rata suhu seluruh komponen: %.2f\n", totalSuhu/float64(len(dataPC)))
}

func tampilkanData() {
	if len(dataPC) == 0 {
		fmt.Println("Belum ada data komponen.")
		return
	}
	for _, k := range dataPC {
		fmt.Printf("SN: %d | Nama: %s | Status: %s | Suhu: %.1f\n", k.SerialNumber, k.Name, k.Status, k.Temperature)
	}
}


func main() {
	var pilihan int
	for {
		clearScreen()

		fmt.Println()
		fmt.Println("Sistem Monitoring SehatinPC      ")
		fmt.Println()
		fmt.Println("1. Tambah Komponen")
		fmt.Println("2. Ubah Komponen")
		fmt.Println("3. Hapus Komponen")
		fmt.Println("4. Tampilkan Semua Data")
		fmt.Println("5. Cari Komponen (Sequential Search - By Name)")
		fmt.Println("6. Cari Komponen (Binary Search - By Name)")
		fmt.Println("7. Urutkan Data (Selection Sort - By Serial)")
		fmt.Println("8. Urutkan Data (Insertion Sort - By Serial)")
		fmt.Println("9. Statistik Kesehatan")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilihan)

		if pilihan == 0 {
			fmt.Println("\nTerima kasih telah menggunakan SehatinPC!")
			break
		}

		switch pilihan {
		case 1:
			tambah()
		case 2:
			ubah()
		case 3:
			hapus()
		case 4:
			tampilkanData()
		case 5:
			cariSequential()
		case 6:
			sort.Slice(dataPC, func(i, j int) bool { return dataPC[i].Name < dataPC[j].Name })
			cariBinary()
		case 7:
			selectionSort()
		case 8:
			insertionSort()
		case 9:
			statistik()
		default:
			fmt.Println("Pilihan tidak valid.")
		}

		fmt.Println()
		
		if !konfirmasiMenu() {
			fmt.Println("\nTerima kasih telah menggunakan SehatinPC!")
			break
		}
	}
}