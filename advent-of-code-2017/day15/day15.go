package day15

func SolvePart1(input string) interface{} {
	genDefA1 := GeneratorDefinition{116, 16807, nil}
	genDefB1 := GeneratorDefinition{299, 48271, nil}

	return runTheJudge(genDefA1, genDefB1, 40000000)
}

func SolvePart2(input string) interface{} {
	genDefA2 := GeneratorDefinition{116, 16807, notDivisbleBy(4)}
	genDefB2 := GeneratorDefinition{299, 48271, notDivisbleBy(8)}

	return runTheJudge(genDefA2, genDefB2, 5000000)
}

func runTheJudge(genDefA GeneratorDefinition, genDefB GeneratorDefinition, count int) (matches int) {
	genA, stopA := CreateGenerator(genDefA)
	defer close(stopA)

	genB, stopB := CreateGenerator(genDefB)
	defer close(stopB)

	matches = Judge(genA, genB, count)

	return
}

type GeneratorDefinition struct {
	firstValue uint64
	factor     uint64
	filter     func(uint64) bool
}

func CreateGenerator(genDef GeneratorDefinition) (dataChannel <-chan uint64, stopChannel chan<- struct{}) {
	dataChan := make(chan uint64, 1000)
	stopChan := make(chan struct{})

	go generate(dataChan, stopChan, genDef.firstValue, genDef.factor, genDef.filter)

	return dataChan, stopChan
}

func generate(dataChan chan<- uint64, stopChan <-chan struct{}, value uint64, factor uint64, filter func(uint64) bool) {
	defer close(dataChan)

	for {
		value = (value * factor) % 2147483647

		if filter != nil && filter(value) {
			continue
		}

		select {
		case <-stopChan:
			return
		case dataChan <- value:
		}
	}
}

func Judge(generatorA, generatorB <-chan uint64, count int) (matches int) {
	for i := 0; i < count; i++ {
		valueA := <-generatorA
		valueB := <-generatorB

		if lowestWord(valueA) == lowestWord(valueB) {
			matches += 1
		}
	}
	return
}

func lowestWord(value uint64) uint64 {
	return value & 0xFFFF
}

func notDivisbleBy(divider uint64) func(uint64) bool {
	return func(value uint64) bool {
		return value%divider != 0
	}
}
