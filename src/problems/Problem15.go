package problems 

import "fmt"

var coords map[int]map[int]int64

func Problem15() {
	size := 20;
	coords = make(map[int]map[int]int64)

	for i := 0; i <= size; i++ {
		coords[i] = make(map[int]int64);
    	coords[i][0] = 1;
    	coords[0][i] = 1;
	}

	calculatePaths(1, size);

	fmt.Println(fmt.Sprintf("Longest path for size %d is %d.", size, coords[len(coords) - 1][len(coords) - 1]))
}

func calculatePaths( startAt int, max int ) {
	for i := startAt; i <= max; i++ {
		coords[i][startAt] = coords[i][startAt - 1] + coords[i - 1][startAt];
        coords[startAt][i] = coords[startAt - 1][i] + coords[startAt][i - 1];	
	}
	if startAt < max {
        calculatePaths( startAt + 1, max );
    }
}