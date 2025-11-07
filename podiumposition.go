package piscine

func PodiumPosition(podium [][]string) [][]string {
	var result [][]string

	for i := len(podium) - 1; i >= 0; i-- {
		result = append(result, podium[i])
	}

	return result
}
