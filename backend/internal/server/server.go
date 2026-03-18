package server

import "github.com/gin-gonic/gin"

type Server struct {
	engine *gin.Engine
}

func New() *Server {
	r := gin.Default() // สร้าง Gin engine

	return &Server{ // คืนค่า pointer ของ Server
		engine: r, // ใส่ Gin engine ที่สร้างไว้ลงในฟิลด์ engine
	}
}

func (s *Server) Run(addr string) error {
	return s.engine.Run(addr) // ส่งต่อให้ Gin เป็นคนเปิดเซิร์ฟเวอร์
}

func (s *Server) Group(path string, middlewares ...gin.HandlerFunc) Group {
	return Group{
		server:      s,                                    // เก็บอ้างอิง Server ปัจจุบัน
		routerGroup: s.engine.Group(path, middlewares...), // สร้าง gin.RouterGroup ภายใต้ path ที่กำหนด
	}
}
