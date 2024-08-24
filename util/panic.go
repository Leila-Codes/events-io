package util

func PanicHandler(errs chan error) {
	for err := range errs {
		panic(err)
	}
}
