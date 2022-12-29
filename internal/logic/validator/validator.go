package validatorI

import (
	"fmt"
	"strings"
	"wiki/internal/model"
	"wiki/internal/service"
)

type sValidatorI struct {
	RuleFunc map[string]func(string, string) *model.ValidatorResult
}

func init() {
	service.RegisterValidatorI(New())
}
func New() *sValidatorI {
	return &sValidatorI{
		RuleFunc: map[string]func(string, string) *model.ValidatorResult{
			"contains": Contains,
			"equal":    Equal,
			"compare":  Compare,
		},
	}
}

func (v *sValidatorI) GetValidator(k string) func(string, string) *model.ValidatorResult {
	if f, ok := v.RuleFunc[strings.ToLower(k)]; ok {
		return f
	}
	fmt.Println("v *sValidator.GetValidator 未找到已实现的断言方法!")
	return nil
}

func Contains(a string, b string) *model.ValidatorResult {
	r := strings.Contains(a, b)
	return &model.ValidatorResult{
		IsPass:  r,
		Message: fmt.Sprintf("contains( %v ) & ( %v ) is %v", a, b, r),
	}
}

func Compare(a string, b string) *model.ValidatorResult {
	r := strings.Compare(a, b)
	if r == 0 {
		return &model.ValidatorResult{
			IsPass:  true,
			Message: fmt.Sprintf("compare( %v ) & ( %v ) is %v", a, b, r),
		}
	}
	return &model.ValidatorResult{
		IsPass:  false,
		Message: fmt.Sprintf("compare( %v ) & ( %v ) is %v", a, b, r),
	}
}

func Equal(a string, b string) *model.ValidatorResult {
	r := a == b
	return &model.ValidatorResult{
		IsPass:  r,
		Message: fmt.Sprintf("compare( %v ) & ( %v ) is %v", a, b, r),
	}
}
