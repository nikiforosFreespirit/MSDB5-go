package app

// Action interface
type Action interface {
	Action(request, origin string) (string, string, error)
}