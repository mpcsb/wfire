package forest



import (
	"fmt" 
	"math/rand"
	"testing" 
)


func TestWriteCsv(t *testing.T) { 
	rand.NewSource(1999)

	for frame:=0; frame < 36; frame++{

		m := make(map[int]Record, 100)
		for i:=0; i < 100; i++ {
			m[i] = Record{Lat:rand.Float64(), Lon:rand.Float64(), Alt:rand.Float64()*10, State:"tree"}
		}
		
		WriteCsv("../viz/" + fmt.Sprintf("%d", frame), m) 
	}

}


func TestSampler(t *testing.T){

	size := 10
	selected := make(map[int]bool, size)
	var Sample_trees []int
	for tries :=0; len(Sample_trees) < size; tries ++{
		fmt.Println(len(Sample_trees))

		i := rand.Intn(20)
	 
		if selected[i]{
			continue
		} else {
			Sample_trees = append(Sample_trees, i)
			selected[i] = true
		}
	}
	fmt.Println(Sample_trees)
	t.Errorf("Height should not be zero: ")
}