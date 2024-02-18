package form

import "github.com/go-playground/form"

// use a single instance of Decoder, it caches struct info
var Decoder *form.Decoder

func InitDecoder() {
	Decoder = form.NewDecoder()
}
