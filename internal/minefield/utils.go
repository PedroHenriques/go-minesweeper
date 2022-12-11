package minefield

import "fmt"

/*
findTilePatch finds all the tile indexes that belong to the patch of the provded
tile.
Returns a slice of tile indexes that belong in the patch.
*/
func findTilePatch(minefield *minefield, initialTileIndex int) []int {
	tilesToReveal := map[string]int{}
	tilesToCheck := []int{initialTileIndex}

	checkIndex := 0
	for checkIndex < len(tilesToCheck) {
		indexToCheck := tilesToCheck[checkIndex]
		checkIndex++
		checkRowIndex := indexToCheck / minefield.cols
		checkColIndex := indexToCheck % minefield.cols

		tilesToReveal[fmt.Sprintf("%v:%v", checkRowIndex, checkColIndex)] = indexToCheck

		if minefield.tiles[indexToCheck].adjacentMines != 0 {
			continue
		}
		if minefield.tiles[indexToCheck].hasMine {
			continue
		}

		for rOffset := -1; rOffset <= 1; rOffset++ {
			rIndex := checkRowIndex + rOffset
			if rIndex < 0 || rIndex > minefield.rows-1 {
				continue
			}

			for cOffset := -1; cOffset <= 1; cOffset++ {
				cIndex := checkColIndex + cOffset
				if cIndex < 0 || cIndex > minefield.cols-1 {
					continue
				}
				if _, ok := tilesToReveal[fmt.Sprintf("%v:%v", rIndex, cIndex)]; ok {
					continue
				}

				tilesToCheck = append(tilesToCheck, rIndex*minefield.cols+cIndex)
			}
		}
	}

	tileIndexes := make([]int, len(tilesToReveal))
	mapIndex := 0
	for _, tileIndex := range tilesToReveal {
		tileIndexes[mapIndex] = tileIndex
		mapIndex++
	}

	return tileIndexes
}

/*
uniqueTileIndexes removes duplicate tile indexes from the provided slice.
*/
func uniqueTileIndexes(source []int) []int {
	keys := make(map[int]bool, len(source))
	output := []int{}

	for _, tIndex := range source {
		if _, found := keys[tIndex]; !found {
			keys[tIndex] = true
			output = append(output, tIndex)
		}
	}

	return output
}

/*
calcTileIndex calulates the tile index based on the tile's row and col
indexes.
*/
func calcTileIndex(rowIndex int, colIndex int, numCols int) int {
	return rowIndex*numCols + colIndex
}
