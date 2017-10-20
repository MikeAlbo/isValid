package isValid

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	intRegex = regexp.MustCompile(`^[0-9]*$`)
	charRegex = regexp.MustCompile(`^[a-zA-Z]*$`)
	apartmentRegex = regexp.MustCompile(`([A-Za-z0-9])\w+`)
	addressRegex = regexp.MustCompile(`^\d+\s[A-z]+\s[A-z]+`)
	// bool
	// state

)

type IsValid interface {
	ErrorCheck() chan error
}

type Value struct {
	text *string
}


func (v *Value)minLength(min int) (bool, error){
	if len(*v.text) < min {
		return false, errors.New(fmt.Sprintf("\"%s\" is less than the min value of %d", v.text, min))
	}
	return true, nil
}

func (v *Value)maxLength(max int) (bool, error){
	if len(*v.text) > max {
		return false, errors.New(fmt.Sprintf("\"%s\" is greater than the max value of %d", v.text, max))
	}
	return true, nil
}

func (v *Value)charString()(bool, error){
	if
}

func (v *Value)intString()(bool, error){

}

func(v *Value)address()(bool, error){

}

func(v *Value)apartmentValidator()(bool, error){

}