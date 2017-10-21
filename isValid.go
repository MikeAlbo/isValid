package isValid

import (
	"errors"
	"fmt"
	//"regexp"
	"time"
)

//var (
//	intRegex = regexp.MustCompile(`^[0-9]*$`)
//	charRegex = regexp.MustCompile(`^[a-zA-Z]*$`)
//	apartmentRegex = regexp.MustCompile(`([A-Za-z0-9])\w+`)
//	addressRegex = regexp.MustCompile(`^\d+\s[A-z]+\s[A-z]+`)
//	// bool
//	// state
//
//)
//
//type Valid interface {
//	Run(...func()error)
//}


type Value struct {
	text *string
}

func Run(functions ...func()error) error {
	for _,f := range functions {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}


func (v *Value)minLength(min int)error{
	if len(*v.text) < min {
		return errors.New(fmt.Sprintf("\"%s\" is less than the min value of %d", *v.text, min))
	}
	return nil
}

func (v *Value)maxLength(max int) error{
	if len(*v.text) > max {
		return  errors.New(fmt.Sprintf("\"%s\" is greater than the max value of %d", *v.text, max))
	}
	return nil
}

func Length(text *string, min, max int) error {
	start := time.Now()
	var v =  Value{text:text}

	test :=[]func() error {
		func() error { return v.minLength(min) },
		func() error { return v.maxLength(max) },
	}
	if err := Run(test...); err != nil {
		fmt.Println(time.Since(start))
		return  err
	}
	fmt.Println(time.Since(start))
	return nil
}

//func (v *Value)charString()(bool, error){
//
//}
//
//func (v *Value)intString()(bool, error){
//
//}
//
//func(v *Value)address()(bool, error){
//
//}
//
//func(v *Value)apartmentValidator()(bool, error){
//
//}

