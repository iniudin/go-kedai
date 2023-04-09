package validation

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
	"strings"
)

type CustomValidator struct {
	validator  *validator.Validate
	translator ut.Translator
}
type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (cv *CustomValidator) Validate(i interface{}) (errors []ErrorResponse) {
	err := cv.validator.Struct(i)
	if err == nil {
		return nil
	}

	for _, e := range err.(validator.ValidationErrors) {
		translatedErr := fmt.Errorf(e.Translate(cv.translator))

		el := ErrorResponse{}
		el.Field = strings.ToLower(e.StructField())
		el.Message = translatedErr.Error()
		errors = append(errors, el)
	}
	return errors
}

func NewValidator(validate *validator.Validate, translator ut.Translator) *CustomValidator {
	return &CustomValidator{validator: validate, translator: translator}
}

func InitializeValidator() (*validator.Validate, ut.Translator) {
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslation.RegisterDefaultTranslations(validate, trans)
	return validate, trans
}
