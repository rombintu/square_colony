package game

// Main object
type Battlefield struct {
	playerList []Player
	// BaseList   []*Base
	resources  []Resource
	cellsCount int
	size       [2]int
}

// Create new Battlefield of game
func NewBattlefield(
	countResorces int,
	playersInfoList [][]string,
	sizeField [2]int,
) Battlefield {
	resourcesInfoList := RandomSelectFromArr(
		resourceTypeListNames,
		countResorces,
	)

	var playerList []Player
	// var baseList []*Base
	var resources []Resource
	var cells int

	for i, info := range playersInfoList {
		player := newPlayer(i+1, info)
		playerList = append(playerList, player)
		// baseList = append(baseList, &player.Base)
		cells++
	}

	for i, info := range resourcesInfoList {
		resource := newResource(i+1, info)
		resources = append(resources, resource)
		cells++
	}

	// logger.Info("Creating new battlefield")
	return Battlefield{
		playerList: playerList,
		// BaseList:   baseList,
		resources:  resources,
		cellsCount: cells,
		size:       sizeField,
	}
}

// Generate new point of cells
func (bf *Battlefield) Init(size [2]int) {
	var cells []*Cell
	// var points [][2]int

	for i := 0; i < len(bf.playerList); i++ {
		cells = append(cells, &bf.playerList[i].Base.Cell)
	}

	for i := 0; i < len(bf.resources); i++ {
		cells = append(cells, &bf.resources[i].Cell)
	}

	for i, cell := range cells {
		cell.SetID(i + 1)
		cell.SetNewPoint(size, cells)
		// points = append(points, cell.GetPoint())
	}
}

// Show main battlefield in console
// func (bf *Battlefield) ShowBattlefield(points [][2]int) {

// 	print("0  | ")

// 	for y0 := 1; y0 < bf.size[1]+1; y0++ {
// 		if y0 < 10 {
// 			fmt.Printf(" %d ", y0)
// 		} else {
// 			fmt.Printf("%d ", y0)
// 		}
// 	}

// 	print("\n")
// 	fmt.Print(strings.Repeat("-+-", bf.size[1]+2))
// 	print("\n")
// 	for x := 1; x < bf.size[0]+1; x++ {
// 		if x < 10 {
// 			fmt.Printf("%d  | ", x)
// 		} else {
// 			fmt.Printf("%d | ", x)
// 		}
// 		for y := 1; y < bf.size[1]+1; y++ {
// 			// TODO
// 			// for _, p := range points {
// 			// 	if reflect.DeepEqual([2]int{x, y}, p) {
// 			// 		fmt.Print("*")
// 			// 	}
// 			// }
// 		}
// 	}
// }

func (bf *Battlefield) GetInfo() {
	// log.
}

func (bf *Battlefield) GetPlayers() []Player {
	return bf.playerList
}

func (bf *Battlefield) GetBases() []Base {
	var arr []Base
	for _, player := range bf.playerList {
		arr = append(arr, player.Base)
	}
	return arr
}

func (bf *Battlefield) GetResources() []Resource {
	return bf.resources
}