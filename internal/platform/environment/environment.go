package environment

const (
	Development int = iota
	Production
)

func GetFromString(env string) int {
	switch env {
	case "production":
		return Production
	default:
		return Development
	}
}
