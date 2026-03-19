package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func New() *Server {
	r := gin.Default()      // สร้าง Gin engine
	r.Use(corsMiddleware()) // ใช้ middleware สําหรับ CORS

	return &Server{ 
		engine: r, // ใส่ Gin engine ที่สร้างไว้ลงในฟิลด์ engine
	}
}

func (s *Server) Run(addr string) error {
	return s.engine.Run(addr) // ส่งต่อให้ Gin เป็นคนเปิดเซิร์ฟเวอร์
}

func (s *Server) Group(path string, middlewares ...gin.HandlerFunc) Group {
	return Group{
		server:      s,                                    // เก็บอ้างอิง Server
		routerGroup: s.engine.Group(path, middlewares...), // สร้าง gin.RouterGroup ภายใต้ path ที่กำหนด
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// golang gin cors middleware
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.Next()
	}
}
