func wrap(h middleware.Handler) httprouter.Handle {
	h = auth.Authenticate(h)
	h = templates.WithTemplates(h, tmpl)
	h = middleware.WithContextValue(h, configContextKey(0), &cfg)
	return base(h)
}

r.GET("/hello", wrap(helloPage))
