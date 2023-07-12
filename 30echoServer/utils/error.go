package utils

func CheckNilError(err error, response string) {

	if response == "panic" {
		if err != nil {
			panic(err)
		}

	}

}
