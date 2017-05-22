package services

func Discovery() Service {
	return Service{key: "discovery"}
}

func Imperium() Service {
	return Service{key: "imperium"}
}

func Socket() Service {
	return Service{key: "socket"}
}

func Project() Service {
	return Service{key: "project"}
}

func Registry() Service {
	return Service{key: "registry"}
}

func Apps() Service {
	return Service{key: "apps"}
}

func Vendors() Service {
	return Service{key: "vendors"}
}

func Schema() Service {
	return Service{key: "schema"}
}

func Preference() Service {
	return Service{key: "preference"}
}

func Notify() Service {
	return Service{key: "notify"}
}

func DataPipe() Service {
	return Service{key: "datapipe"}
}

func Config() Service {
	return Service{key: "config"}
}
