package scraper

import (
	"reflect"
	"testing"
)

func TestNewRoute(t *testing.T) {
	routes := NewRoute()
	if routes == nil {
		t.Fatal("NewRoute() returned nil, expected an initialized map")
	}
	if routes.Count() != 0 {
		t.Errorf("Expected new route count to be 0, got %d", routes.Count())
	}
}

func TestAddAndHas(t *testing.T) {
	routes := NewRoute()
	key := "key"
	subRoutes := []string{"link1", "link2"}

	routes.Add(key, subRoutes)

	if !routes.Has(key) {
		t.Errorf("Expected Has(%q) to be true, got false", key)
	}

	// Verify that the subroutes were added correctly.
	routeStruct, ok := routes[key]
	if !ok {
		t.Fatalf("Expected key %q to be present in routes", key)
	}
	if !reflect.DeepEqual(routeStruct.SubRoutes, subRoutes) {
		t.Errorf("Expected SubRoutes %v, got %v", subRoutes, routeStruct.SubRoutes)
	}
}

func TestCount(t *testing.T) {
	routes := NewRoute()

	// Initially count should be 0.
	if count := routes.Count(); count != 0 {
		t.Errorf("Expected count 0, got %d", count)
	}

	// After adding entries, count should reflect the number of keys.
	routes.Add("a", []string{"linkA"})
	routes.Add("b", []string{"linkB"})
	if count := routes.Count(); count != 2 {
		t.Errorf("Expected count 2, got %d", count)
	}
}

func TestAppend(t *testing.T) {
	routes := NewRoute()
	targetKey := "target"
	sourceKey := "source"

	targetSubRoutes := []string{"t1", "t2"}
	sourceSubRoutes := []string{"s1", "s2"}

	routes.Add(targetKey, targetSubRoutes)
	routes.Add(sourceKey, sourceSubRoutes)

	// Append the subroutes from source to target.
	routes.Append(targetKey, sourceKey)

	// Expect target's subroutes to be its original ones appended with source's subroutes.
	expectedSubRoutes := append([]string(nil), targetSubRoutes...)
	expectedSubRoutes = append(expectedSubRoutes, sourceSubRoutes...)

	actualSubRoutes := routes[targetKey].SubRoutes
	if !reflect.DeepEqual(actualSubRoutes, expectedSubRoutes) {
		t.Errorf("Expected target subroutes %v, got %v", expectedSubRoutes, actualSubRoutes)
	}
}

func TestAppendNonExisting(t *testing.T) {
	routes := NewRoute()
	targetKey := "target"
	nonExistingKey := "nonexistent"

	// Only add target route.
	targetSubRoutes := []string{"t1", "t2"}
	routes.Add(targetKey, targetSubRoutes)

	// Attempt to append from a non-existing source. The target should remain unchanged.
	routes.Append(targetKey, nonExistingKey)
	if actualSubRoutes := routes[targetKey].SubRoutes; !reflect.DeepEqual(actualSubRoutes, targetSubRoutes) {
		t.Errorf("Expected target subroutes %v, got %v", targetSubRoutes, actualSubRoutes)
	}

	// Attempt to append to a non-existing target. Nothing should happen.
	routes.Append(nonExistingKey, targetKey)
	if routes.Has(nonExistingKey) {
		t.Errorf("Expected Has(%q) to be false, got true", nonExistingKey)
	}
}
