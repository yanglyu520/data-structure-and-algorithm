func subdomainVisits(cpdomains []string) []string {
	dm := make(map[string]int)

	//transform the input into a hashmap
	for _, item := range cpdomains {
		//split the visits from domain name by space separator
		raw := strings.Split(item, " ")

		//convert the number of visits from string to int
		numOfVisits, _ := strconv.Atoi(raw[0])

		//split the domains like this x.y.z --> [x y z]
		domains := strings.Split(raw[1], ".")

		//create the new hashmap like this {x.y.z:10, y.z:10, z:10}
		for i, _ := range domains {
			dm[strings.Join(domains[i:], ".")] += numOfVisits

		}
	}

	//translate the hashmap into a result array
	result := []string{}
	for k, v := range dm {
		result = append(result, strconv.Itoa(v)+" "+k)
	}
	return result
}

