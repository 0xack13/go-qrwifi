package main

import (
	"fmt"
	"image/png"
	"net/http"
	"text/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

// apple wifi code : WIFI:T:WPA;S:mynetwork;P:mypass;;
type Page struct {
	Title string
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/generator/", viewCodeHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "QR Code Generator"}

	t, _ := template.ParseFiles("generator.html")
	t.Execute(w, p)
}

func viewCodeHandler(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("s")
	p := r.FormValue("p")
	sp := fmt.Sprintf("WIFI:T:WPA;S:%s;P:%s;;", s, p)
	fmt.Print(sp)
	qrCode, _ := qr.Encode(sp, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)

	png.Encode(w, qrCode)
}
