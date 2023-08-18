package gx

func Component(selector string) *component {
	return &component{selector: selector}
}

type ConstructorFunc func(props any) any

type component struct {
	selector    string
	imports     []*component
	template    string
	constructor ConstructorFunc
}

func (c *component) Imports(components ...*component) *component {
	c.imports = components
	return c
}

func (c *component) Template(template string) *component {
	c.template = template
	return c
}

func (c *component) Constructor(constructor ConstructorFunc) *component {
	c.constructor = constructor
	return c
}
