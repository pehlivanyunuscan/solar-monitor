package main

import (
	"log"
	"os"
)

// Genel uygulama logları için logger
var appLogger *log.Logger

// Audit (kim, ne yaptı) logları için logger
var auditLogger *log.Logger

func init() {
	// Uygulama log dosyası
	appLogFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Uygulama log dosyası açılamadı: %v", err)
	}
	appLogger = log.New(appLogFile, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Audit log dosyası
	auditLogFile, err := os.OpenFile("audit.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Audit log dosyası açılamadı: %v", err)
	}
	auditLogger = log.New(auditLogFile, "AUDIT: ", log.Ldate|log.Ltime|log.Lshortfile)
}
