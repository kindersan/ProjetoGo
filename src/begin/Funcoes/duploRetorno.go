package main
import ("fmt")


func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

func main() {
	region, continent := location("NYC")
	fmt.Printf("Carson Matt mora em %s, %s\n", region, continent)
}

/*func location(name, city string) (region, continent string) {
	switch city {
	case "New York", "LA", "Chicago":
		continent = "North America"
	default:
		continent = "Unknown"
	}
	return
}

func main() {
	region, continent := location("Matt", "LA")
	fmt.Printf("%s lives in %s", region, continent)
}*/
