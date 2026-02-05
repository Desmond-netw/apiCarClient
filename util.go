package main

// Building lookup maps

func manufacturerMap(manufacturer []Manufacturer) map[int]string {
	manuMap := make(map[int]string)
	for _, manuf := range manufacturer {
		manuMap[manuf.ID] = manuf.Name
	}
	return manuMap
}

// category map
func categoryMap(category []Category) map[int]string {
	catMap := make(map[int]string)
	for _, cats := range category {
		catMap[cats.ID] = cats.Name
	}
	return catMap
}
