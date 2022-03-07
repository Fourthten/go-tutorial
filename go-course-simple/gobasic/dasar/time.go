package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var time1 = time.Now()
	fmt.Printf("time1 %v\n", time1)
	// time.Date(year, month, date, hour, minute, second, nanosecond, timezone)
	var time2 = time.Date(2022, 12, 20, 10, 10, 0, 0, time.UTC)
	fmt.Printf("time2 %v\n", time2)

	var now = time.Now()
	fmt.Println("year:", now.Year(), "month:", now.Month())

	var layoutFormat, value string
	var date time.Time
	layoutFormat = "2020-01-01 20:05:00"
	value = "2021-01-01 21:10:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())
	layoutFormat = "02/01/2020 MST"
	value = "02/05/2021 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())

	var dates, _ = time.Parse(time.RFC822, "02 Sep 21 07:00 WIB")
	fmt.Println(dates.String())

	var _date, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
	var dateS1 = _date.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dateS1", dateS1)
	var dateS2 = _date.Format(time.RFC3339)
	fmt.Println("dateS2", dateS2)

	// var _dates, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")
	// if err != nil {
	// 	fmt.Println("error", err.Error())
	// 	return
	// }
	// fmt.Println(_dates)

	fmt.Println("start")
	time.Sleep(time.Second * 2)
	fmt.Println("after 2s")

	// for true {
	// 	fmt.Println("Hello !")
	// 	time.Sleep(1 * time.Second)
	// }

	var timer = time.NewTimer(2 * time.Second)
	fmt.Println("start1")
	<-timer.C
	fmt.Println("finish1")

	var ch = make(chan bool)
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("expired2")
		ch <- true
	})
	fmt.Println("start2")
	<-ch
	fmt.Println("finish2")

	<-time.After(2 * time.Second)
	fmt.Println("expired3")

	// done := make(chan bool)
	// ticker := time.NewTicker(time.Second)
	// go func() {
	// 	time.Sleep(4 * time.Second)
	// 	done <- true
	// }()
	// for {
	// 	select {
	// 	case <-done:
	// 		ticker.Stop()
	// 		return
	// 	case t := <-ticker.C:
	// 		fmt.Println("Hello !", t)
	// 	}
	// }

	var timeout = 3
	var _ch = make(chan bool)
	go timers(timeout, _ch)
	// go watcher(timeout, _ch)
	var input string
	fmt.Print("is 125/25 ? ")
	fmt.Scan(&input)
	if input == "5" {
		fmt.Println("answer right")
	} else {
		fmt.Println("answer wrong")
	}

	start := time.Now()
	time.Sleep(3 * time.Second)
	duration := time.Since(start)
	fmt.Println("time sec:", duration.Seconds())
	fmt.Println("time min:", duration.Minutes())
	fmt.Println("time hour:", duration.Hours())

	t1 := time.Now()
	time.Sleep(3 * time.Second)
	t2 := time.Now()
	durations := t2.Sub(t1)
	fmt.Println("time:", durations.Seconds())
	fmt.Println("time:", durations.Minutes())
	fmt.Println("time:", durations.Hours())

	var n time.Duration = 3
	_duration := n * time.Second
	fmt.Println(_duration)

	_n := 3
	_durations := time.Duration(_n) * time.Second
	fmt.Println(_durations)
}

func timers(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out! ", timeout, "s")
	os.Exit(0)
}

// Method			Return Type		Penjelasan
// now.Year()		int				Tahun
// now.YearDay()	int				Hari ke-? di mulai awal tahun
// now.Month()		int				Bulan
// now.Weekday()	string			Nama hari. Bisa menggunakan now.Weekday().String() untuk mengambil bentuk string-nya
// now.ISOWeek()	( int, int)		Tahun dan minggu ke-? mulai awal tahun
// now.Day()		int				Tanggal
// now.Hour()		int				Jam
// now.Minute()		int				Menit
// now.Second()		int				Detik
// now.Nanosecond()	int				Nano Detik
// now.Local()		time.Time		Date-time dalam timezone lokal
// now.location()	*time.Location	Mengambil informasi lokasi, apakah local atau utc. Bisa menggunakan now.Location().String() untuk mengambil bentuk string-nya
// now.Zone()		( string, int )	Mengembalikan informasi timezone offset dalam string dan numerik. Sebagai contoh WIB, 25200
// now.IsZero()		bool			Deteksi apakah nilai object now adalah 01 Januari tahun 1, 00:00:00 UTC. Jika iya maka bernilai true
// now.UTC()		time.Time		Date-time dalam timezone UTC
// now.Unix()		int64			Date-time dalam format unix time
// now.UnixNano()	int64			Date-time dalam format unix time. Infomasi nano detik juga dimasukan
// now.String()		string			Date-time dalam string

// Layout Format	Penjelasan									Contoh Data
// 2006				Tahun 4 digit								2015
// 006				Tahun 3 digit								015
// 06				Tahun 2 digit								15
// 01				Bulan 2 digit								05
// 1				Bulan 1 digit jika dibawah bulan 10, 		5, 12
// 					selainnya 2 digit
// January			Nama bulan dalam bahasa inggris				Sepetember, August
// Jan				Nama bulan dalam bahasa inggris, 3 huruf	Sep, Aug
// 02				Tanggal 2 digit								02
// 2				Tanggal 1 digit jika dibawah bulan 10, 		8, 31
// 					selainnya 2 digit
// Monday			Nama hari dalam bahasa inggris				Saturday, Friday
// Mon				Nama hari dalam bahasa inggris, 3 huruf		Sat, Fri
// 15				Jam dengan format 24 jam					18
// 03				Jam dengan format 12 jam 2 digit			05, 11
// 3				Jam dengan format 12 jam 1 digit jika 		5, 11
// 					dibawah jam 11, selainnya 2 digit
// PM				AM/PM, biasa digunakan dengan format 		PM, AM
// 					jam 12 jam
// 04				Menit 2 digit								08
// 4				Menit 1 digit jika dibawah menit 10, 		8, 24
// 					selainnya 2 digit
// 05				Detik 2 digit								06
// 5				Detik 1 digit jika dibawah detik 10, 		6, 36
// 					selainnya 2 digit
// 999999			Nano detik									124006
// MST				Lokasi timezone								UTC, WIB, EST
// Z0700			Offset timezone								Z, +0700, -0200

// Predefined Layout Format		Layout Format
// time.ANSIC					Mon Jan _2 15:04:05 2006
// time.UnixDate				Mon Jan _2 15:04:05 MST 2006
// time.RubyDate				Mon Jan 02 15:04:05 -0700 2006
// time.RFC822					02 Jan 06 15:04 MST
// time.RFC822Z					02 Jan 06 15:04 -0700
// time.RFC850					Monday, 02-Jan-06 15:04:05 MST
// time.RFC1123					Mon, 02 Jan 2006 15:04:05 MST
// time.RFC1123Z				Mon, 02 Jan 2006 15:04:05 -0700
// time.RFC3339					2006-01-02T15:04:05Z07:00
// time.RFC3339Nano				2006-01-02T15:04:05.999999999Z07:00
// time.Kitchen					3:04PM
// time.Stamp					Jan _2 15:04:05
// time.StampMilli				Jan _2 15:04:05.000
// time.StampMicro				Jan _2 15:04:05.000000
// time.StampNano				Jan _2 15:04:05.000000000
// time.Nanosecond				1
// time.Microsecond				1000
// time.Millisecond				1000000
// time.Second					1000000000
// time.Minute					1000000000000
// time.Hour					1000000000000000
