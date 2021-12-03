package erratum

func Use(opener ResourceOpener, input string) (resError error) {
	res, err := opener()
	_, ok := err.(TransientError)
	keepTry := ok
	for keepTry {
		res, err = opener()
		_, ok := err.(TransientError)
		if err == nil {
			break
		} else if ok {
			continue
		} else if err != nil {
			keepTry = false
			return err
		}
	}
	defer res.Close()
	defer func() {
		recoverError := recover()
		frobError, ok := recoverError.(FrobError)
		if ok {
			res.Defrob(frobError.defrobTag)
			resError = frobError.inner
		} else {
			resError, _ = recoverError.(error)
		}
	}()
	res.Frob(input)

	return nil
}
