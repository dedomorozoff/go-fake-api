package api

import (
	_ "embed"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"net/http"

	"github.com/alexl/go-fake-api/internal/utils"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

//go:embed assets/moon.jpg
var moonImage []byte

// CreateLunarWatermark создает изображение с водяным знаком на Луне
func CreateLunarWatermark() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Парсинг multipart form
		if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request", nil)
			return
		}

		// Получение сообщения
		message := r.FormValue("message")
		if validationErrors := utils.ValidateWatermarkMessage(message); len(validationErrors) > 0 {
			utils.RespondWithValidationError(w, validationErrors)
			return
		}

		// Получение файла
		file, _, err := r.FormFile("fileimage")
		if err != nil {
			validationErrors := map[string][]string{
				"fileimage": {"field fileimage can not be blank"},
			}
			utils.RespondWithValidationError(w, validationErrors)
			return
		}
		defer file.Close()

		// Декодирование изображения
		img, _, err := image.Decode(file)
		if err != nil {
			validationErrors := map[string][]string{
				"fileimage": {"invalid image file"},
			}
			utils.RespondWithValidationError(w, validationErrors)
			return
		}

		// Создание нового изображения для рисования
		bounds := img.Bounds()
		rgba := image.NewRGBA(bounds)
		draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

		// Добавление водяного знака
		addWatermark(rgba, message)

		// Отправка изображения
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Disposition", "attachment; filename=lunar-watermark.jpg")
		
		if err := jpeg.Encode(w, rgba, &jpeg.Options{Quality: 90}); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to encode image", nil)
			return
		}
	}
}

// addWatermark добавляет текстовый водяной знак на изображение
func addWatermark(img *image.RGBA, text string) {
	col := color.RGBA{255, 255, 255, 255} // Белый цвет
	point := fixed.Point26_6{
		X: fixed.Int26_6(img.Bounds().Dx()/2 - len(text)*3) * 64,
		Y: fixed.Int26_6(img.Bounds().Dy() - 50) * 64,
	}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(text)
}
