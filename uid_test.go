package uid

import(
	"fmt"
	"sync"
	"testing"
)

func TestUID(t *testing.T) {
	t.Log("begun")
	var shutdownWg = &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		//time.Sleep(time.Second * 1)
		shutdownWg.Add(1)
		go func(i int) {
			defer shutdownWg.Done()
			idGen, err := NewUIDGenerator("ait", int32(i))
			if err != nil {
				t.Error("initialization failed")
				return
			}
			fmt.Println(idGen.ID())
		}(i)
	}
	shutdownWg.Wait()
	//time.Sleep(time.Second * 0)
	t.Log("completed")
}

func BenchmarkUID(b *testing.B) {
	idGen, err := NewUIDGenerator("ait", int32(1))
	if err != nil {
		return
	}
	for j := 1; j < b.N; j++ {
		idGen.ID()
	}
}
