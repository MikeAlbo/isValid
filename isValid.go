package isValid

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

var (
	intRegex       = regexp.MustCompile(`^[0-9]*$`)
	charRegex      = regexp.MustCompile(`^[a-zA-Z]*$`)
	apartmentRegex = regexp.MustCompile(`([A-Za-z0-9])\w+`)
	addressRegex   = regexp.MustCompile(`^\d+\s[A-z]+\s[A-z]+`)
	// bool
	// state

)

//type Valid interface {
//	Run(...func()error)
//}

type Value struct {
	text *string
}

func Run(functions ...func() error) error {
	for _, f := range functions {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

func (v *Value) minLength(min int) error {
	if len(*v.text) < min {
		return errors.New(fmt.Sprintf("\"%s\" is less than the min value of %d", *v.text, min))
	}
	return nil
}

func (v *Value) maxLength(max int) error {
	if len(*v.text) > max {
		return errors.New(fmt.Sprintf("\"%s\" is greater than the max value of %d", *v.text, max))
	}
	return nil
}

func (v *Value) isInt() error {
	if !intRegex.MatchString(*v.text) {
		return errors.New(fmt.Sprintf("\"%s\" is is not an intiger", *v.text))
	}
	return nil
}

func (v *Value) isChar() error {
	if !charRegex.MatchString(*v.text) {
		return errors.New(fmt.Sprintf("\"%s\" should only contain values between a-z", *v.text))
	}
	return nil
}

func (v *Value) isAddress() error {
	if !addressRegex.MatchString(*v.text) {
		return errors.New(fmt.Sprintf("\"%s\" is not a correctly formated address", *v.text))
	}
	return nil
}

func (v *Value) isApartment() error {
	if !apartmentRegex.MatchString(*v.text) {
		return errors.New(fmt.Sprintf("\"%s\" is a correctly formatted apartment  number", *v.text))
	}
	return nil
}

func Length(text *string, min, max int) error {
	start := time.Now()
	v := Value{text: text}

	test := []func() error{
		func() error { return v.minLength(min) },
		func() error { return v.maxLength(max) },
	}
	if err := Run(test...); err != nil {
		fmt.Println(time.Since(start))
		return err
	}
	fmt.Println(time.Since(start))
	return nil
}

func CharString(text *string, min, max int) error {
	start := time.Now()
	v := Value{text: text}

	test := []func() error{
		func() error { return v.minLength(min) },
		func() error { return v.maxLength(max) },
		func() error { return v.isChar() },
	}

	if err := Run(test...); err != nil {
		fmt.Println(time.Since(start))
		return err
	}

	fmt.Println(time.Since(start))
	return nil
}

func IntString(text *string, min, max int) error {
	start := time.Now()
	v := Value{text: text}

	test := []func() error{
		func() error { return v.minLength(min) },
		func() error { return v.maxLength(max) },
		func() error { return v.isInt() },
	}

	if err := Run(test...); err != nil {
		fmt.Println(time.Since(start))
		return err
	}

	fmt.Println(time.Since(start))
	return nil
}

func Address(text *string, min, max int) error {
	start := time.Now()
	v := Value{text: text}

	test := []func() error{
		func() error { return v.minLength(min) },
		func() error { return v.maxLength(max) },
		func() error { return v.isAddress() },
	}

	if err := Run(test...); err != nil {
		fmt.Println(time.Since(start))
		return err
	}

	fmt.Println(time.Since(start))
	return nil
}

func ApartmentValidator(text *string, min,max int) error {
	start := time.Now()
	v := Value{text: text}

	test := []func() error{
		func() error { return v.minLength(min) },
		func() error { return v.maxLength(max) },
		func() error { return v.isApartment() },
	}

	if err := Run(test...); err != nil {
		fmt.Println(time.Since(start))
		return err
	}

	fmt.Println(time.Since(start))
	return nil
}
