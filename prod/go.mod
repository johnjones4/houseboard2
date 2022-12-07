module main

go 1.18

replace github.com/johnjones4/houseboard2/core => ../core

require (
	github.com/fogleman/gg v1.3.0
	github.com/johnjones4/houseboard2/core v0.0.0-00010101000000-000000000000
)

require (
	github.com/arran4/golang-ical v0.0.0-20221122102835-109346913e54 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	golang.org/x/image v0.2.0 // indirect
)
