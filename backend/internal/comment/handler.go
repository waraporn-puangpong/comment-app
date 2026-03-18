package comment

import (
	"backend/internal/db"
	"backend/internal/server"
	"backend/models"
	"net/http"
	"time"
)

type IHandler interface {
	SaveComment(ctx *server.Context) error
	GetComments(ctx *server.Context) error
}

type handler struct {
	service IService
}

func NewHandler(apiServer *server.Server, dbconn db.IMongo) IHandler {
	return &handler{
		service: NewService(dbconn),
	}
}

// SaveComment godoc
// @Summary Create a comment
// @Description Create a new comment
// @Tags comments
// @Accept json
// @Produce json
// @Param comment body models.Comment true "Comment payload"
// @Success 201 {object} models.Comment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /comments/update-comment [post]
func (h *handler) SaveComment(ctx *server.Context) error {
	var input models.Comment

	//1. แปลง JSON body เป็น struct models.Comment
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request body",
		})
		return nil
	}

	//2. ถ้าไม่ส่ง CreatedAt มา หรือส่งเป็นค่า zero time, ให้ตั้งเป็นเวลาปัจจุบัน
	if input.CreatedAt.IsZero() {
		input.CreatedAt = time.Now().UTC()
	}

	if err := h.service.SaveComment(ctx, input); err != nil {
		return err
	}

	ctx.JSON(http.StatusCreated, input)
	return nil
}

// GetComments godoc
// @Summary List comments
// @Description Get all comments sorted by newest date
// @Tags comments
// @Produce json
// @Success 200 {array} models.Comment
// @Failure 500 {object} map[string]string
// @Router /comments [get]
func (h *handler) GetComments(ctx *server.Context) error {
	comments, err := h.service.GetComments()
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, comments)
	return nil
}
