package helper

func IfErrNotNul(err error) {
	if err != nil {
		panic(err)
		return
	}
}
