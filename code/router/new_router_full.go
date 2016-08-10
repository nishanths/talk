mw := router.NewMiddlewareChain(
	base,
	middleware.WithContextValue(configContextKey(0), &cfg),
	templates.WithTemplates(tmpl),
	auth.Authenticate
)

r.GET("/hello", mw, helloPage)
