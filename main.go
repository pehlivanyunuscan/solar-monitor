package main

import (
	"encoding/json"
	"log"
	"net/http"

	"main.go/logging"
	"main.go/prometheus"
)

func cameraHandler(w http.ResponseWriter, r *http.Request) {
	// Örnek bir kamera endpoint'i
	logging.AppLogger.Println("Kamera endpoint çağrıldı")

	// Prometheus'tan kamera metriklerini çek
	// Örn: kamera_aktif_sayisi isminde bir metrik
	promURL := "http://localhost:9090" // Prometheus adresin
	metric := "kamera_aktif_sayisi"    // Çekmek istediğin metrik

	result, err := prometheus.QueryPrometheus(promURL, metric)
	if err != nil || len(result.Data.Result) == 0 {
		logging.AppLogger.Printf("Prometheus verisi alınamadı: %v", err)
		http.Error(w, "Prometheus verisi alınamadı", http.StatusInternalServerError)
		return
	}

	// Metrik değerini çek
	value := result.Data.Result[0].Value[1]

	// Audit loglama
	logging.LogAudit("user123", "/camera", r.Method, http.StatusOK, r.RemoteAddr, nil, "Kamera erişildi")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"aktif_kamera_sayisi": value,
	})
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
