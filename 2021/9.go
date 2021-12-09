/* --- Day 9: Smoke Basin ---
These caves seem to be lava tubes. Parts are even still volcanically active; small hydrothermal vents release smoke into the caves that slowly settles like rain.

If you can model how the smoke flows through the caves, you might be able to avoid it and be that much safer. The submarine generates a heightmap of the floor of the nearby caves for you (your puzzle input).

Smoke flows to the lowest point of the area it's in. For example, consider the following heightmap:

2199943210
3987894921
9856789892
8767896789
9899965678
Each number corresponds to the height of a particular location, where 9 is the highest and 0 is the lowest a location can be.

Your first goal is to find the low points - the locations that are lower than any of its adjacent locations. Most locations have four adjacent locations (up, down, left, and right); locations on the edge or corner of the map have three or two adjacent locations, respectively. (Diagonal locations do not count as adjacent.)

In the above example, there are four low points, all highlighted: two are in the first row (a 1 and a 0), one is in the third row (a 5), and one is in the bottom row (also a 5). All other locations on the heightmap have some lower adjacent location, and so are not low points.

The risk level of a low point is 1 plus its height. In the above example, the risk levels of the low points are 2, 1, 6, and 6. The sum of the risk levels of all low points in the heightmap is therefore 15.

Find all of the low points on your heightmap. What is the sum of the risk levels of all low points on your heightmap?

--- Part Two ---
Next, you need to find the largest basins so you know what areas are most important to avoid.

A basin is all locations that eventually flow downward to a single low point. Therefore, every low point has a basin, although some basins are very small. Locations of height 9 do not count as being in any basin, and all other locations will always be part of exactly one basin.

The size of a basin is the number of locations within the basin, including the low point. The example above has four basins.

The top-left basin, size 3:

2199943210
3987894921
9856789892
8767896789
9899965678
The top-right basin, size 9:

2199943210
3987894921
9856789892
8767896789
9899965678
The middle basin, size 14:

2199943210
3987894921
9856789892
8767896789
9899965678
The bottom-right basin, size 9:

2199943210
3987894921
9856789892
8767896789
9899965678
Find the three largest basins and multiply their sizes together. In the above example, this is 9 * 14 * 9 = 1134.

What do you get if you multiply together the sizes of the three largest basins?
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func bfsSize(data [][]int, coord [2]int) int {
	queue := make([][2]int, 0)
	count := 0
	visited := make(map[[2]int]bool)
	queue = append(queue, [2]int{coord[0], coord[1]})

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if visited[curr] {
			continue
		}
		visited[curr] = true
		count += 1
		row, col := curr[0], curr[1]
		if row != 0 && data[row-1][col] != 9 && !visited[[2]int{row - 1, col}] {
			queue = append(queue, [2]int{row - 1, col})
		}
		if col != 0 && data[row][col-1] != 9 && !visited[[2]int{row, col - 1}] {
			queue = append(queue, [2]int{row, col - 1})
		}
		if col != len(data[0])-1 && data[row][col+1] != 9 && !visited[[2]int{row, col + 1}] {
			queue = append(queue, [2]int{row, col + 1})
		}
		if row != len(data)-1 && data[row+1][col] != 9 && !visited[[2]int{row + 1, col}] {
			queue = append(queue, [2]int{row + 1, col})
		}
	}
	return count
}

func main() {
	file, err := os.Open("9_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([][]int, 0)
	lows := make([][2]int, 0)
	basins := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := make([]int, 0)
		for _, char := range scanner.Text() {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
		}
		data = append(data, row)
	}
	for i, row := range data {
		for j, num := range row {
			if i == 0 {
				if j == 0 {
					if data[i][j+1] > num && data[i+1][j] > num {
						lows = append(lows, [2]int{i, j})
					}
				} else if j == len(row)-1 {
					if data[i][j-1] > num && data[i+1][j] > num {
						lows = append(lows, [2]int{i, j})
					}
				} else {
					if data[i][j-1] > num && data[i][j+1] > num && data[i+1][j] > num {
						lows = append(lows, [2]int{i, j})
					}
				}
			} else if i == len(data)-1 {
				if j == 0 {
					if data[i-1][j] > num && data[i][j+1] > num {
						lows = append(lows, [2]int{i, j})
					}
				} else if j == len(row)-1 {
					if data[i][j-1] > num && data[i-1][j] > num {
						lows = append(lows, [2]int{i, j})
					}
				} else {
					if data[i][j-1] > num && data[i][j+1] > num && data[i-1][j] > num {
						lows = append(lows, [2]int{i, j})
					}
				}
			} else if j == 0 {
				if data[i-1][j] > num && data[i+1][j] > num && data[i][j+1] > num {
					lows = append(lows, [2]int{i, j})
				}
			} else if j == len(row)-1 {
				if data[i-1][j] > num && data[i+1][j] > num && data[i][j-1] > num {
					lows = append(lows, [2]int{i, j})
				}
			} else {
				if data[i-1][j] > num && data[i+1][j] > num && data[i][j-1] > num && data[i][j+1] > num {
					lows = append(lows, [2]int{i, j})
				}
			}
		}
	}

	risk := 0
	for _, v := range lows {
		risk += data[v[0]][v[1]] + 1
	}

	for _, v := range lows {
		basins = append(basins, bfsSize(data, [2]int{v[0], v[1]}))
	}

	threeLargest := [3]int{0, 0, 0}
	for _, size := range basins {
		if size > threeLargest[0] {
			threeLargest[2] = threeLargest[1]
			threeLargest[1] = threeLargest[0]
			threeLargest[0] = size
		} else if size > threeLargest[1] {
			threeLargest[2] = threeLargest[1]
			threeLargest[1] = size
		} else if size > threeLargest[2] {
			threeLargest[2] = size
		}
	}

	fmt.Println("Sum of risk levels:", risk)
	fmt.Println("Product of three largest basins:", threeLargest[0]*threeLargest[1]*threeLargest[2])
}
