package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	mu    sync.Mutex
	count int
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/gif", handlerGIF)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func handlerGIF(w http.ResponseWriter, r *http.Request) {

	palette[1] = color.RGBA{uint8(rand.Int()), uint8(rand.Int()), uint8(rand.Int()), 255}
	fmt.Printf("palette: %v\n", palette[1])
	lissajous(w)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", count)
	mu.Unlock()
}

var palette = []color.Color{color.White, color.Black}

func lissajous(out io.Writer) {
	const (
		whitelndex = 0 // Первый цвет палитры
		blacklndex = 1 // Следующий цвет палитры
	)
	const (
		cycles  = 8      // Количество полных колебаний x
		res     = 0.0001 // Угловое разрешение
		size    = 250    // Канва изображения охватывает [size..+size]
		nframes = 64     // Количество кадров анимации
		delay   = 8      // Задержка между кадрами (единица - 10мс)
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // Относительная частота колебаний у
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blacklndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Примечание: игнорируем ошибки
}
