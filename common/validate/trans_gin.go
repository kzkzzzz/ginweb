package validate

import (
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	ginVt    *validator.Validate
	ginTrans ut.Translator
)

func init() {
	registerGin()
}

func registerGin() {
	if vt2, ok := binding.Validator.Engine().(*validator.Validate); ok {
		ginVt = vt2
		ginTrans = registerTranslate(ginVt)
	}

}
