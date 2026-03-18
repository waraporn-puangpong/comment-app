package server 

import "github.com/gin-gonic/gin"

type Group struct { 
	server      *Server          
	routerGroup *gin.RouterGroup 
}

func (g Group) GET(path string, handlerFn HandlerFunc, middlewares ...gin.HandlerFunc) { 
	handlers := append(
		middlewares,
		func(ctx *gin.Context) {
			g.server.handle(ctx, handlerFn) // ส่งให้ตัวกลางของ Server เรียก handler และจัดการ error
		},
	)

	g.routerGroup.GET(path, handlers...) // ผูก path กับ handlers 
}

func (g Group) POST(path string, handlerFn HandlerFunc, middlewares ...gin.HandlerFunc) {
	handlers := append(
		middlewares,
		func(ctx *gin.Context) {
			g.server.handle(ctx, handlerFn) // ส่งให้ตัวกลางของ Server เรียก handler และจัดการ error
		},
	)

	g.routerGroup.POST(path, handlers...) // ผูก path กับ handlers 
}
