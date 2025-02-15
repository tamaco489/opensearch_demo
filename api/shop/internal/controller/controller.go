package controller

type Controllers struct {
	env string
}

func NewControllers(env string) *Controllers {
	return &Controllers{
		env: env,
	}
}
