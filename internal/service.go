package internal

import (
	"github.com/rs/zerolog"
)

type Service struct {
 	repo repository
	log  zerolog.Logger 
}
