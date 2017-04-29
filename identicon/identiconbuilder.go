package identicon

import (
	"image"
)

type IdenticonBuilder interface {
	Build(in []byte) (image.RGBA, error)
}