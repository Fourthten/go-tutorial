package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	kubus      Kubus   = Kubus{4}
	realVolume float64 = 64
	realArea   float64 = 96
	realAround float64 = 48
)

func TestCountVolume(t *testing.T) {
	t.Logf("Volume: %.2f", kubus.Volume())
	if kubus.Volume() != realVolume {
		t.Errorf("Wrong! %.2f", realVolume)
	}
}

func TestCountArea(t *testing.T) {
	t.Logf("Area: %.2f", kubus.Area())
	if kubus.Area() != realArea {
		t.Errorf("Wrong! %.2f", realArea)
	}
}

func TestCountAround(t *testing.T) {
	t.Logf("Around: %.2f", kubus.Around())
	if kubus.Around() != realAround {
		t.Errorf("Wrong! %.2f", realAround)
	}
}

// go test testing.go testing_test.go -v

func BenchmarkCountArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.Area()
	}
}

// go test testing.go testing_test.go -bench=.

func TestCountVolume2(t *testing.T) {
	assert.Equal(t, kubus.Volume(), realVolume, "wrong volume")
}

func TestCountArea2(t *testing.T) {
	assert.Equal(t, kubus.Area(), realArea, "wrong area")
}

func TestCountAround2(t *testing.T) {
	assert.Equal(t, kubus.Around(), realAround, "wrong around")
}

// Method		Utility
// Log()		Menampilkan log
// Logf()		Menampilkan log menggunakan format
// Fail()		Menandakan terjadi Fail() dan proses testing fungsi tetap diteruskan
// FailNow()	Menandakan terjadi Fail() dan proses testing fungsi dihentikan
// Failed()		Menampilkan laporan fail
// Error()		Log() diikuti dengan Fail()
// Errorf()		Logf() diikuti dengan Fail()
// Fatal()		Log() diikuti dengan failNow()
// Fatalf()		Logf() diikuti dengan failNow()
// Skip()		Log() diikuti dengan SkipNow()
// Skipf()		Logf() diikuti dengan SkipNow()
// SkipNow()	Menghentikan proses testing fungsi dilanjutkan ke testing fungsi setelahnya
// Skiped()		Menampilkan laporan skip
// Parallel()	Menge-set bahwa eksekusi testing adalah parallel

// Package		Utility
// assert		Berisikan tools standar untuk testing
// http			Berisikan tools untuk keperluan testing http
// mock			Berisikan tools untuk mocking object
// require		Sama seperti assert, hanya saja jika terjadi fail pada saat test akan menghentikan eksekusi program
// suite		Berisikan tools testing yang berhubungan dengan struct dan method
