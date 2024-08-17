package utils

func Map[T any, R any](slice []T, f func(T) (R, error)) ([]R, error) {
	var err error
	mapped := make([]R, len(slice))
	for i, v := range slice {
		mapped[i], err = f(v)
	}
	return mapped, err
}
