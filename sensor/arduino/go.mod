module github.com/hybridgroup/gophercon2019

go 1.16

require (
	github.com/hybridgroup/gophercon2019/freemono v0.0.0-00010101000000-000000000000
	tinygo.org/x/drivers v0.16.0
	tinygo.org/x/tinydraw v0.0.0-20200416172542-c30d6d84353c
	tinygo.org/x/tinyfont v0.2.1
)

replace github.com/hybridgroup/gophercon2019/freemono => ./fonts
