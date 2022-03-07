package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/novalagung/gubrak/v2"
)

type M map[string]interface{}

var cookieName = "CookieData"

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{}
	if storedCookie, _ := r.Cookie(cookieName); storedCookie != nil {
		c = storedCookie
	}

	if c.Value == "" {
		c = &http.Cookie{}
		c.Name = cookieName
		c.Value = gubrak.RandomString(32)
		c.Expires = time.Now().Add(5 * time.Minute)
		http.SetCookie(w, c)
	}

	w.Write([]byte(c.Value))
}

func ActionDelete(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{}
	c.Name = cookieName
	c.Expires = time.Unix(0, 0)
	c.MaxAge = -1
	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", ActionIndex)
	http.HandleFunc("/delete", ActionDelete)

	fmt.Println("Server started at http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}

// Property	Tipe Data	Deskripsi
// Value	string		Data yang disimpan di cookie
// Path		string		Scope path cookie
// Domain	string		Scope domain cookie
// Expires	time.Time	Durasi cookie, ditulis dalam tipe time.Time
// MaxAge	int 		Durasi cookie, ditulis dalam detik (numerik)
// Secure	bool		Scope cookie dalam konteks protocol yang digunakan ketika pengaksesan web.
// 						Property ini hanya berguna pada saat web server SSL/TLS enabled. Jika false,
// 						maka cookie yang disimpan ketika web diakses menggunakan protocol http://,
// 						tetap bisa diakses lewat https://, dan berlaku juga untuk kebalikannya.
// 						Jika true, pada saat pengaksesan lewat protocol https://, maka data cookie
// 						akan di-enkripsi. Sedangkan pada pengaksesan lewat protokol http://, cookie
// 						disimpan seperti biasa (tanpa dienkripsi). Jika dalam satu web server, dua
// 						protokol tersebut bisa diakses, http:// dan https://, maka aturan di atas
// 						tetap berlaku untuk masing-masing protokol, dengan catatan data yang disimpan
// 						lewat https:// hanya bisa diakses lewat protokol tersebut.
// HttpOnly	bool		Jika false, maka cookie bisa dibuat lewat back end (Go), maupun lewat front end
// 						(javascript) Jika true, maka cookie hanya bisa diciptakan dari back end
