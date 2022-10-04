package driver

type Router interface {
	RegisterRouter()
}

type RegistryContract interface {
	RunApplication()
}

func Run(rv RegistryContract) {
	if rv != nil {
		rv.RunApplication()
	}
}
