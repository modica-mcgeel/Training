package controller

func AddUser() {
	user := []string{"Lia", "McGee", "Service Desk Service"}
	fields := map[string]string{"Company Name:": "ACC", "First Name:": "Lia"}
	for _, v := range user {
		println(v)
	}

	for f, v := range fields {
		println(f, v)
	}
}
