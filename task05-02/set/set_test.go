package set

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

// 3. Протестируйте производительность операций чтения и записи на множестве действительных чисел,
// безопасность которого обеспечивается sync.Mutex и sync.RWMutex для разных вариантов использования:
// 10% запись, 90% чтение; 50% запись, 50% чтение; 90% запись, 10% чтение
func MutexSetRWTest(s Set, iterNum int, readOps int, writeOps int) {
	// We will use random numbers for (read from)/(write into) sets
	rand.Seed(time.Now().Unix())

	wg := &sync.WaitGroup{}

	for i := 0; i < iterNum; i++ {
		for writeOp := 0; writeOp < writeOps; writeOp++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup, number float64) {
				defer wg.Done()
				s.Add(number)
			}(wg, rand.Float64())
		}

		for readOp := 0; readOp < readOps; readOp++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup, number float64) {
				defer wg.Done()
				s.Has(number)
			}(wg, rand.Float64())
		}
	}
	wg.Wait()

}

func BenchmarkMutextSetR10W90(t *testing.B) {
	s := NewMutexSet()
	MutexSetRWTest(s, t.N, 100, 900)
}

func BenchmarkRWMutextSetR10W90(t *testing.B) {
	s := NewRWMutexSet()
	MutexSetRWTest(s, t.N, 100, 900)
}

func BenchmarkMutextSetR50W50(t *testing.B) {
	s := NewMutexSet()
	MutexSetRWTest(s, t.N, 500, 500)
}

func BenchmarkRWMutextSetR50W50(t *testing.B) {
	s := NewRWMutexSet()
	MutexSetRWTest(s, t.N, 500, 500)
}

func BenchmarkMutextSetR90W10(t *testing.B) {
	s := NewMutexSet()
	MutexSetRWTest(s, t.N, 900, 100)
}

func BenchmarkRWMutextSetR90W10(t *testing.B) {
	s := NewRWMutexSet()
	MutexSetRWTest(s, t.N, 90, 10)
}
