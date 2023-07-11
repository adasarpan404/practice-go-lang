package utils

func CheckNilError(err error, response string) {

	if response == "panic" {
		panic(err)
	}

}
