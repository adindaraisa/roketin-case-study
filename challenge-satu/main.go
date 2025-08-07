package main

import "fmt"

func main() {
	var earthHour, earthMinute, earthSecond int

	fmt.Print("Masukkan jam di bumi (jam menit detik): ")
	fmt.Scan(&earthHour, &earthMinute, &earthSecond)

	totalEarthSeconds := earthHour*3600 + earthMinute*60 + earthSecond

	// Konversi total detik bumi menjadi total detik di planet Roketin
	// 1 hari bumi = 24*3600 detik
	// 1 hari Roketin = 10*100*100 detik = 100000 detik
	totalRoketinSeconds := totalEarthSeconds * 100000 / 86400

	// Hitung jam, menit, detik di Roketin
	roketinHour := totalRoketinSeconds / (100 * 100)
	roketinMinute := (totalRoketinSeconds % (100 * 100)) / 100
	roketinSecond := totalRoketinSeconds % 100

	fmt.Printf("On earth %02d:%02d:%02d, on Roketin planet %02d:%02d:%02d\n",
		earthHour, earthMinute, earthSecond,
		roketinHour, roketinMinute, roketinSecond)
}
