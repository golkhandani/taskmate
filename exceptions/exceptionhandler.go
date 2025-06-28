package exceptions

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}
