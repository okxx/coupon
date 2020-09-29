package app

type C struct {
}

type HandlerFunc func(*C)

type Routes interface {
	Use(args ...interface{}) Routes
	Get(relative string, handles ...HandlerFunc) Routes
	Post(relative string, handles ...HandlerFunc) Routes
	Delete(relative string, handles ...HandlerFunc) Routes
	Put(relative string, handles ...HandlerFunc) Routes
	Patch(relative string, handles ...HandlerFunc) Routes
	Options(relative string, handles ...HandlerFunc) Routes
	Head(relative string, handles ...HandlerFunc) Routes
	Static(relative string, handles ...HandlerFunc) Routes

	Group(prefix string, handles ...HandlerFunc) Routes
}

type Route struct {

	use			bool			`json:"use"`


	Method 		string 			`json:"method"`
	Path 		string 			`json:"path"`
	Params 		[]string 		`json:"params"`
	Handles 	[]HandlerFunc 	`json:"handles"`
}

// match route
func (r *Route) match() bool {

	return true
}
