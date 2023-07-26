package structValidator

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func Validate[T interface{}](obj *T) (*T, []error) {
	validate := validator.New()

	// Custom validate
	validate.RegisterValidation("cpf", validateCPF)

	err := validate.Struct(obj)

	if err == nil {
		return obj, nil
	}

	validationErrors := err.(validator.ValidationErrors)
	errs := make([]error, len(validationErrors))
	fmt.Println(len(validationErrors))
	for index, err := range validationErrors {
		fmt.Println(err.Tag())
		switch err.Tag() {
		case "required":
			errs[index] = errors.New("Campo: " + err.Field() + " é obrigatório!")
		case "min":
			errs[index] = errors.New("Campo: " + err.Field() + " é menor que o tamanho mínimo: " + err.Param())
		case "max":
			errs[index] = errors.New("Campo: " + err.Field() + " é maior que o tamanho máximo: " + err.Param())
		case "email", "cpf":
			errs[index] = errors.New("Campo: " + err.Field() + " é inválido!")
		}
	}

	fmt.Println(errs)
	return new(T), errs
}

func validateCPF(fl validator.FieldLevel) bool {
	// Expressão regular para validar o formato do CPF (apenas números e 11 dígitos)
	cpfRegex := regexp.MustCompile(`^\d{11}$`)
	cpf := fl.Field().String()

	// Verificar o formato do CPF
	if !cpfRegex.MatchString(cpf) {
		return false
	}

	// Outras verificações podem ser feitas para validar o CPF real
	// Por exemplo, aplicar a regra do dígito verificador

	return true
}
