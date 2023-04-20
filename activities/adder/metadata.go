package sample

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	Num1 int `md:"num1,required"`
	Num2 int `md:"num2,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	num1, _ := coerce.ToInt(values["num1"])
	num2, _ := coerce.ToInt(values["num2"])
	r.Num1 = num1
	r.Num2 = num2
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"num1": r.Num1,
		"num2": r.Num2,
	}
}

type Output struct {
	Sum int `md:"sum"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	sum, _ := coerce.ToInt(values["sum"])
	o.Sum = sum
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"sum": o.Sum,
	}
}
