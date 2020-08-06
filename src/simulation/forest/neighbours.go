package forest

import (
	"fmt"
	"sync"

	"gonum.org/v1/gonum/spatial/vptree"
)

const workerNum = 4

// GetNeighbours finds trees closer than a distance d
func (f *Forest) GetNeighbours(d float64) {
	// handle vp, err :=....
	VP, _ := vptree.New(f.Tree_Coords, 0, nil)
	fmt.Println("VP found")

	for i, q := range f.Tree_Coords {
		var keep vptree.Keeper
		keep = vptree.NewDistKeeper(d)
		VP.NearestSet(keep, q)

		for _, neighbourTree := range keep.(*vptree.DistKeeper).Heap {
			tree := neighbourTree.Comparable.(TreeCoord)
			f.Tree_lst[i].Neighbours = append(f.Tree_lst[i].Neighbours, tree.ID)
		}
	}
	fmt.Println("Neighbours found")
}

// DistributedNeighbour finds adjacent trees over const workerNum of cores
func (f *Forest) DistributedNeighbour(d float64) {
	// VP, _ := vptree.New(f.Tree_Coords, 0, nil)
	fmt.Println("VP found")

	wg := sync.WaitGroup{}
	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go f.run(i, 3.0, &wg)
	}

}

func (f *Forest) run(worker int, d float64, wg *sync.WaitGroup) {
	VP, _ := vptree.New(f.Tree_Coords, 0, nil)

	for i, q := range f.Tree_Coords {
		if i%workerNum != worker {
			continue
		}

		var keep vptree.Keeper
		keep = vptree.NewDistKeeper(d)
		VP.NearestSet(keep, q)

		for _, neighbourTree := range keep.(*vptree.DistKeeper).Heap {
			tree := neighbourTree.Comparable.(TreeCoord)
			f.Tree_lst[i].Neighbours = append(f.Tree_lst[i].Neighbours, tree.ID)
		}
	}
	// let the wait group know we finished
	wg.Done()

}
