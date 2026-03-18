package server 

import (
	"net/http" 

	"github.com/gin-gonic/gin" 
)

type Context struct { 
	*gin.Context 
}

type HandlerFunc func(*Context) error 

func (s *Server) handle(ctx *gin.Context, handlerFn HandlerFunc) { 
	c := &Context{Context: ctx} // แปลง gin.Context เป็น Context

	if err := handlerFn(c); err != nil { // ถ้า handler คืน error ให้ตอบกลับเป็น 500
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(), // ส่งข้อความ errorแบบ JSON
		})
		return 
	}
}
