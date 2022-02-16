As I'm pretty new to Golan and WebAssembly, I'll be keeping my learning notes here:

- Unfortunatly, Golang does not support compiling to WASM with linked C libraries, as discussed
  here: https://github.com/veandco/go-sdl2/issues/438. Project can go sideways now:
  - [ ] write WASM Backend in JavaScript
  - [ ] use Golang library which can be compiled to WASM ([ebiten](https://ebiten.org/))

