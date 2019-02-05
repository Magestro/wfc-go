package wfc

import "math"


type Model struct {
	wave [][]bool

	propagator [][][]int
	Compatible [][][]int
	observed   []int

	Stack     [][1]int //(int, int)[] stack;
	StackSize int

	random      interface{} // protected Random random;
	FMX, FMY, T int
	periodic    bool

	weights          []float64
	WeightLogWeights []float64

	SumsOfOnes                                           []int
	SumOfWeights, SumOfWeightLogWeights, StartingEntropy float64
	SumsOfWeights, SumsOfWeightLogWeights, Entropies     []float64
}

func New(width, height int) *Model {
	return &Model{
		FMX: width,
		FMY: height,
	}
}

func (m *Model) Init() {
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

	m.StartingEntropy = math.Log(m.SumOfWeights) - m.SumOfWeightLogWeights / m.SumOfWeights

	m.SumsOfOnes = make([]int, m.FMX * m.FMY)
	m.SumsOfWeights = make([]float64, m.FMX * m.FMY)
	m.SumsOfWeightLogWeights = make([]float64, m.FMX * m.FMY)
	m.Entropies = make([]float64, m.FMX * m.FMY)

	m.Stack = make([][1]int, len(m.wave) * m.T)
	m.StackSize = 0
}
