package logging

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// Genel uygulama logları için logger
var AppLogger *log.Logger

// Audit (kim, ne yaptı) logları için logger
var AuditLogger *log.Logger

func init() {
	// Uygulama log dosyası
	appLogFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Uygulama log dosyası açılamadı: %v", err)
	}
	AppLogger = log.New(appLogFile, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Audit log dosyası
	auditLogFile, err := os.OpenFile("audit.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Audit log dosyası açılamadı: %v", err)
	}
	AuditLogger = log.New(auditLogFile, "AUDIT: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Audit log için struct
type AuditLog struct {
	Timestamp  string      `json:"timestamp"`
	User       string      `json:"user"`
	Endpoint   string      `json:"endpoint"`
	Method     string      `json:"method"`
	StatusCode int         `json:"status_code"`
	ClientIP   string      `json:"client_ip"`
	Params     interface{} `json:"params,omitempty"`
	Message    string      `json:"message,omitempty"`
}

func LogAudit(user, endpoint, method string, statusCode int, clientIP string, params interface{}, message string) {
	auditLog := AuditLog{
		Timestamp:  time.Now().Format(time.RFC3339),
		User:       user,
		Endpoint:   endpoint,
		Method:     method,
		StatusCode: statusCode,
		ClientIP:   clientIP,
		Params:     params,
		Message:    message,
	}

	jsonEntry, err := json.Marshal(auditLog)
	if err != nil {
		AppLogger.Printf("Audit log oluşturulamadı: %v", err)
		return
	}

	AuditLogger.Println(string(jsonEntry))
}
