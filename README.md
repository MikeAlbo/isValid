## isValid
a validation package for simple text inputs

* character strings
* int strings
* street address 
* apartment numbers

## goals
1. create underlying methods that return specif errors related to their regex test
2. pass bools and errors through channels
3. have the pubic methods check channels for responses 
4. Public methods will return an error or nil to the end user

### things to think about
* possibly add go routines to each function call in the variadic Run function
    * create a new error channel `errors = make(chan error, len(funcs))`
    * return an error or nil to the channel `errChan <- go f() return ` 
    * watch channel for errors 
    *maybe overkill :-) 