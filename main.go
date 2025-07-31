package main

import (
	"log"
	"net/http"

	"main.go/logging"
)

func cameraHandler(w http.ResponseWriter, r *http.Request) {
	// Örnek bir kamera endpoint'i
	logging.AppLogger.Println("Kamera endpoint çağrıldı")

	// Audit loglama
	logging.LogAudit("user123", "/camera", r.Method, http.StatusOK, r.RemoteAddr, nil, "Kamera erişildi")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Kamera endpointi çalışıyor"))
}

func main() {

	http.HandleFunc("/camera", cameraHandler)

	// Uygulama logu örneği
	logging.AppLogger.Println("Uygulama başlatıldı")

	log.Println("Sunucu 8080 portunda çalışıyor...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logging.AppLogger.Fatalf("Sunucu başlatılamadı: %v", err)
	}

}
