package scraper

type Routes map[string]SubRoutesStruct

type Broken map[string]int

func NewRoute() Routes {
	return make(Routes)
}

type SubRoutesStruct struct {
	SubRoutes []string
}

func (r Routes) Add(val string, subRoutes []string) {
	r[val] = SubRoutesStruct{
		SubRoutes: subRoutes,
	}
}
func (r Routes) Count() int {
	return len(r)
}

// Has checks if a string exists in the set.
func (r Routes) Has(str string) bool {
	_, ok := r[str]
	return ok
}

func (r Routes) Append(target string, source string) {
	targetRoute, targetExists := r[target]
	sourceRoute, sourceExists := r[source]

	if targetExists && sourceExists {
		targetRoute.SubRoutes = append(targetRoute.SubRoutes, sourceRoute.SubRoutes...)
		r[target] = targetRoute // Update the target route in the map
	}
}
