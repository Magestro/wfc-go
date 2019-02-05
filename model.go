package wfc

import "math"

var (
	dx       = []int{-1, 0, 1, 0}
	dy       = []int{0, 1, 0, -1}
	opposite = []int{2, 3, 0, 1}
)

type Model struct {
	wave [][]bool

	propagator [][][]int
	Compatible [][][]int
	observed   []int

	Stack     [][1]int //(int, int)[] stack; todo
	StackSize int

	random      interface{} // protected Random random; todo
	FMX, FMY, T int
	periodic    bool

	weights          []float64
	WeightLogWeights []float64

	SumsOfOnes                                           []int
	SumOfWeights, SumOfWeightLogWeights, StartingEntropy float64
	SumsOfWeights, SumsOfWeightLogWeights, Entropies     []float64

	OnBoundary func(x, y int) bool // todo
	//public abstract System.Drawing.Bitmap Graphics(); todo
}

func New(width, height int) *Model {
	return &Model{
		FMX: width,
		FMY: height,
	}
}

// todo private
func (m *Model) init() {
	m.wave = make([][]bool, m.FMX*m.FMY)
	m.Compatible = make([][][]int, len(m.wave))

	for i := 0; i < len(m.wave); i++ {
		m.wave[i] = make([]bool, m.T)
		m.Compatible[i] = make([][]int, m.T)
		for t := 0; t < m.T; t++ {
			m.Compatible[i][t] = make([]int, 4)
		}
	}

	m.WeightLogWeights = make([]float64, m.T)
	m.SumOfWeights = 0
	m.SumOfWeightLogWeights = 0

	for t := 0; t < m.T; t++ {
		m.WeightLogWeights[t] = m.weights[t] * math.Log(m.weights[t])
		m.SumOfWeights += m.weights[t]
		m.SumOfWeightLogWeights += m.WeightLogWeights[t]
	}

	m.StartingEntropy = math.Log(m.SumOfWeights) - m.SumOfWeightLogWeights/m.SumOfWeights

	m.SumsOfOnes = make([]int, m.FMX*m.FMY)
	m.SumsOfWeights = make([]float64, m.FMX*m.FMY)
	m.SumsOfWeightLogWeights = make([]float64, m.FMX*m.FMY)
	m.Entropies = make([]float64, m.FMX*m.FMY)

	m.Stack = make([][1]int, len(m.wave)*m.T)
	m.StackSize = 0
}

// todo private
func (m *Model) observe() (ok bool, err error) {
	var (
		min    = 1E+3
		argmin = -1
	)

	for i := 0; i < len(m.wave); i++ {
		if m.OnBoundary(i%m.FMX, i/m.FMX) {
			continue
		}

		if m.SumsOfOnes[i] == 0 {
			return
		}

		if m.SumsOfOnes[i] > 1 && m.Entropies[i] <= min {
			noise := 1E-6 // * random.NextDouble();
			if m.Entropies[i]+noise < min {
				min = m.Entropies[i] + noise
				argmin = i
			}
		}
	}

	if argmin == -1 {
		ok = true
		m.observed = make([]int, m.FMX*m.FMY)
		for i := 0; i < len(m.wave); i++ {
			for t := 0; t < m.T; t++ {
				if m.wave[i][t] {
					m.observed[i] = t
					return
				}
			}
		}

		return
	}

	distribution := make([]float64, m.T)
	for i := 0; i < len(m.wave); i++ {
		if m.wave[argmin][i] {
			distribution[i] = m.weights[i]
		} else {
			distribution[i] = 0
		}
	}
	var r float64 //random element from distribution todo clarify
	// int r = distribution.Random(random.NextDouble()); todo

	for i := 0; i < m.T; i++ {
		if m.wave[argmin][i] != (float64(i) == r) {
			m.ban(argmin, i)
		}
	}

	return
}

// todo protected
func (m *Model) propagate() {
	panic("implement me")
}

func (m *Model) Run(seed, limit int) bool {
	panic("implement me")
}

// todo protected
func (m *Model) ban(i, t int) {
	panic("implement me")
}

// todo protected virtual
func (m *Model) clear() {
	panic("implement me")
}
