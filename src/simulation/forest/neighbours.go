package forest

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"gonum.org/v1/gonum/spatial/vptree"
)

// numcpu := runtime.NumCPU()
const workerNum = 2

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

// DistributedGetNeighbours finds adjacent trees over const workerNum of cores
func (f *Forest) DistributedGetNeighbours(d float64) {

	neighbourMap := make([][][]int, workerNum)
	for i := range neighbourMap {
		neighbourMap[i] = make([][]int, len(f.Tree_lst))
		// for j := range neighbourMap[i] {
		// 	neighbourMap[i][j] = make([]int, 10)
		// }
	}

	wg := sync.WaitGroup{}

	for worker := 0; worker < workerNum; worker++ {
		wg.Add(1)
		time.Sleep(1)
		go f.run(worker, d, &neighbourMap[worker], &wg)
	}
	wg.Wait()

	for worker := 0; worker < workerNum; worker++ { // iterate over all workers
		// fmt.Println(neighbourMap[worker])
		neighbourArray := neighbourMap[worker]
		for idx := range neighbourArray { // iterate over list of neighbours calculate by each worker
			if allSame(neighbourArray[idx]) {
				continue // no neighbours for this tree = idx, by this worker
			}
			f.Tree_lst[idx].Neighbours = append(f.Tree_lst[idx].Neighbours, neighbourArray[idx]...)
		}
	}
	fmt.Println("Distributed Get Neighbours done")
}

func (f *Forest) run(worker int, d float64, neighbourArray *[][]int, wg *sync.WaitGroup) {
	// defer wg.Done()
	VP, _ := vptree.New(f.Tree_Coords, 0, nil)

	for i, q := range f.Tree_Coords {
		// each worker has a predefined collection of tree idxs
		if i%runtime.NumCPU() != worker {
			continue
		}

		var keep vptree.Keeper
		keep = vptree.NewDistKeeper(d)
		VP.NearestSet(keep, q)
		for _, neighbourTree := range keep.(*vptree.DistKeeper).Heap {
			tree := neighbourTree.Comparable.(TreeCoord)
			// fmt.Println(tree.ID)
			(*neighbourArray)[i] = append((*neighbourArray)[i], tree.ID)
		}
		// (*neighbourArray)[i] = append((*neighbourArray)[i], 111115555)
	}
	wg.Done()
}

func allSame(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}
