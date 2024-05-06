package ports

// GinContext implement HTTPContext
type HTTPContext interface {
	JSON(int, interface{})
	// define other methods for HTTP framework or net/http to implement implicity
}
