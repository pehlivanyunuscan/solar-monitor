package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"main.go/logging"
	"main.go/prometheus"
)

func cameraHandler(w http.ResponseWriter, r *http.Request) {
	// Örnek bir kamera endpoint'i
	logging.AppLogger.Println("Kamera endpoint çağrıldı")

	// Prometheus'tan kamera metriklerini çek
	// Örn: camera_status isminde bir metrik
	promURL := os.Getenv("PROM_URL") // Prometheus URL'sini ortam değişkeninden al
	if promURL == "" {
		promURL = "http://localhost:9090" // Varsayılan adres
	}
	metric := "camera_status" // Çekmek istediğin metrik

	result, err := prometheus.QueryPrometheus(promURL, metric)
	if err != nil || len(result.Data.Result) == 0 {
		logging.AppLogger.Printf("Prometheus verisi alınamadı: %v", err)
		http.Error(w, "Prometheus verisi alınamadı", http.StatusInternalServerError)
		return
	}

	// Tüm kameraların durumunu bir map olarak hazırla
	cameraStates := make(map[string]string)
	for _, res := range result.Data.Result {
		cameraName := res.Metric["camera"]
		statusValue := res.Value[1].(string)
		if statusValue == "1" {
			cameraStates[cameraName] = "açık"
		} else {
			cameraStates[cameraName] = "kapalı"
		}
	}

	logging.LogAudit("user123", "/camera", r.Method, http.StatusOK, r.RemoteAddr, nil, "Kamera durumları listelendi")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cameraStates)
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
