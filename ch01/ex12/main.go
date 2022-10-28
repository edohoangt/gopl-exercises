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
	"strconv"
)

func main() {
	var cycles = 5 // default # cycles

	handler := func(w http.ResponseWriter, r *http.Request) {
			cyclesStr := r.FormValue("cycles")
			if cyclesStr != "" {
				var err error
				cycles, err = strconv.Atoi(cyclesStr)
				if err != nil {
					fmt.Fprintf(w, "bad cycles param: %s", cyclesStr)
					return
				}
			}
			lissajous(cycles, w)
		}

	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8002", nil))
}
	
func lissajous(cycles int, out io.Writer) {
	var palette = []color.Color{color.White, 
		color.RGBA{0x00, 0xFF, 0x00, 0xFF},
		color.RGBA{0xFF, 0x00, 0x00, 0xFF}, 
		color.RGBA{0x00, 0x00, 0xFF, 0xFF},
	}

	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < 2*math.Pi*float64(cycles); t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			indexColor := uint8(rand.Float64() * 3) + 1 // random int from 1 to 3
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				indexColor)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}