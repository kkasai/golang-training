 go build gopl.io/ch3/mandelbrot
./mandelbrot | go run imageconv.go -f jpeg > mandelbrot2.jpeg

./mandelbrot | go run imageconv.go -f png > mandelbrot2.png

./mandelbrot | go run imageconv.go -f gif > mandelbrot2.gif
