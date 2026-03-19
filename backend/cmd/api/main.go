package main

import (
	"backend/docs"
	"backend/internal/db"
	"backend/internal/router"
	"backend/internal/server"
	"context"
	"log"
	"time"
)

const (
	defaultMongoURI = "mongodb://localhost:27017/?readPreference=primary&directConnection=true"
	defaultDBName   = "comment-app"
	defaultHTTPAddr = ":8080"
)

func main() {
	// 1. สร้าง context พร้อม timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 2. เชื่อมต่อ MongoDB
	mongoConn, err := db.NewMongo(ctx, defaultMongoURI, defaultDBName)
	if err != nil {
		log.Fatalf("failed to connect mongo: %v", err)
	}

	// ปิด MongoDB ตอนโปรแกรมกำลังจะจบ โดยกำหนด timeout กันการค้าง
	defer func() {
		// ใช้ context แยกสำหรับขั้นตอนปิด connection
		closeCtx, closeCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer closeCancel()

		// พยายามปิด connection และบันทึก error ถ้ามี
		if err := mongoConn.Close(closeCtx); err != nil {
			log.Printf("failed to close mongo connection: %v", err)
		}
	}()

	// 3. สร้าง HTTP server
	docs.SwaggerInfo.BasePath = "/api"
	apiServer := server.New()

	// 4. ลงทะเบียน route
	router.Setup(apiServer, mongoConn)

	// 5. เริ่มรัน API server ที่พอร์ต :8080
	if err := apiServer.Run(defaultHTTPAddr); err != nil {
		log.Fatalf("failed to run api server: %v", err)
	}
}
