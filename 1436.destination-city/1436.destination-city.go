func destCity(paths [][]string) string {
	cityMap := make(map[string]bool, len(paths))

	for _, path := range paths {
		if _, exist := cityMap[path[0]]; !exist {
			cityMap[path[0]] = false
		} else {
			delete(cityMap, path[0])
		}

		if _, exist := cityMap[path[1]]; !exist {
			cityMap[path[1]] = true
		} else {
			delete(cityMap, path[1])
		}

	}

	for city, is := range cityMap {
		if is {
			return city
		}
	}

	return ""
}
