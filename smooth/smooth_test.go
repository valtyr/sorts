package smooth

import (
	"math"
	"rand"
	"testing"
)

import (
	"github.com/runningwild/sorts/smooth"
)

/*
 * You do not need to edit anything below here.
 */

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
var float64s = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}
var strings = [...]string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

func TestSortIntSlice(t *testing.T) {
	data := ints
	a := smooth.IntSlice(data[0:])
	smooth.Sort(a)
	if !smooth.IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestSortFloat64Slice(t *testing.T) {
	data := float64s
	a := smooth.Float64Slice(data[0:])
	smooth.Sort(a)
	if !smooth.IsSorted(a) {
		t.Errorf("sorted %v", float64s)
		t.Errorf("   got %v", data)
	}
}

func TestSortStringSlice(t *testing.T) {
	data := strings
	a := smooth.StringSlice(data[0:])
	smooth.Sort(a)
	if !smooth.IsSorted(a) {
		t.Errorf("sorted %v", strings)
		t.Errorf("   got %v", data)
	}
}

func TestInts(t *testing.T) {
	data := ints
	smooth.Ints(data[0:])
	if !smooth.IntsAreSorted(data[0:]) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestFloat64s(t *testing.T) {
	data := float64s
	smooth.Float64s(data[0:])
	if !smooth.Float64sAreSorted(data[0:]) {
		t.Errorf("sorted %v", float64s)
		t.Errorf("   got %v", data)
	}
}

func TestStrings(t *testing.T) {
	data := strings
	smooth.Strings(data[0:])
	if !StringsAreSorted(data[0:]) {
		t.Errorf("sorted %v", strings)
		t.Errorf("   got %v", data)
	}
}

func TestSortLarge_Random(t *testing.T) {
	n := 1000000
	if testing.Short() {
		n /= 100
	}
	data := make([]int, n)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(100)
	}
	if smooth.IntsAreSorted(data) {
		t.Fatalf("terrible rand.rand")
	}
	Ints(data)
	if !smooth.IntsAreSorted(data) {
		t.Errorf("sort didn't sort - 1M ints")
	}
}

func shuffle(data []int) {
	var n int
	for i := len(data)-1; i > 0; i-- {
		n = rand.Intn(i)
		v[i], v[n] = v[n], v[i]
	}
}

func random(data []int) {
	for i := range data {
		data[i] = rand.Int()
	}
}

func reverse(data []int) {
	n := len(data) - 1
	for i := range data {
		data[i] = n
		n--
	}
}

func inOrder(data []int) {
	for i := range data {
		data[i] = i
	}
}

func BenchmarkSorted10(b *testing.B) {
	b.StopTimer()
	data := make([]int, 10)
	inOrder(data)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		smooth.Ints(data)
	}
}

func BenchmarkRandom10(b *testing.B) {
	b.StopTimer()
	data := make([]int, 10)
	inOrder(data)
	for i := 0; i < b.N; i++ {
		random(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkShuffled10(b *testing.B) {
	b.StopTimer()
	data := make([]int, 10)
	inOrder(data)
	for i := 0; i < b.N; i++ {
		shuffle(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkReversed10(b *testing.B) {
	b.StopTimer()
	data := make([]int, 10)
	for i := 0; i < b.N; i++ {
		reverse(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkSorted100(b *testing.B) {
	b.StopTimer()
	data := make([]int, 100)
	inOrder(data)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		smooth.Ints(data)
	}
}

func BenchmarkRandom100(b *testing.B) {
	b.StopTimer()
	data := make([]int, 100)
	inOrder(data)
	for i := 0; i < b.N; i++ {
		random(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkShuffled100(b *testing.B) {
	b.StopTimer()
	data := make([]int, 100)
	inOrder(data)
	for i := 0; i < b.N; i++ {
		shuffle(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkReversed100(b *testing.B) {
	b.StopTimer()
	data := make([]int, 100)
	for i := 0; i < b.N; i++ {
		reverse(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkSorted1k(b *testing.B) {
	b.StopTimer()
	data := make([]int, 1000)
	inOrder(data)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		smooth.Ints(data)
	}
}

func BenchmarkRandom1k(b *testing.B) {
	b.StopTimer()
	data := make([]int, 1000)
	for i := 0; i < b.N; i++ {
		random(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkShuffled1k(b *testing.B) {
	b.StopTimer()
	data := make([]int, 1000)
	inOrder(data)
	for i := 0; i < b.N; i++ {
		shuffle(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkReversed1k(b *testing.B) {
	b.StopTimer()
	data := make([]int, 1000)
	for i := 0; i < b.N; i++ {
		reverse(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkSorted10k(b *testing.B) {
	b.StopTimer()
	data := make([]int, 10000)
	inOrder(data)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		smooth.Ints(data)
	}
}

func BenchmarkRandom10k(b *testing.B) {
	b.StopTimer()
	data := make([]int, 10000)
	inOrder(data)
	for i := 0; i < b.N; i++ {
		random(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkShuffled10k(b *testing.B) {
	b.StopTimer()
	data := make([]int, 10000)
	inOrder(data)
	for i := 0; i < b.N; i++ {
		shuffle(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkReversed10k(b *testing.B) {
	b.StopTimer()
	data := make([]int, 10000)
	for i := 0; i < b.N; i++ {
		reverse(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkSorted100k(b *testing.B) {
	b.StopTimer()
	data := make([]int, 100000)
	inOrder(data)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		smooth.Ints(data)
	}
}

func BenchmarkRandom100k(b *testing.B) {
	b.StopTimer()
	data := make([]int, 100000)
	inOrder(data)
	for i := 0; i < b.N; i++ {
		random(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkShuffled100k(b *testing.B) {
	b.StopTimer()
	data := make([]int, 100000)
	inOrder(data)
	for i := 0; i < b.N; i++ {
		shuffle(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkReversed100k(b *testing.B) {
	b.StopTimer()
	data := make([]int, 100000)
	for i := 0; i < b.N; i++ {
		reverse(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkSorted1M(b *testing.B) {
	b.StopTimer()
	data := make([]int, 1000000)
	inOrder(data)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		smooth.Ints(data)
	}
}

func BenchmarkRandom1M(b *testing.B) {
	b.StopTimer()
	data := make([]int, 1000000)
	inOrder(data)
	for i := 0; i < b.N; i++ {
		random(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkShuffled1M(b *testing.B) {
	b.StopTimer()
	data := make([]int, 1000000)
	inOrder(data)
	for i := 0; i < b.N; i++ {
		shuffle(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}

func BenchmarkReversed1M(b *testing.B) {
	b.StopTimer()
	data := make([]int, 1000000)
	for i := 0; i < b.N; i++ {
		reverse(data)
		b.StartTimer()
		smooth.Ints(data)
		b.StopTimer()
	}
}
