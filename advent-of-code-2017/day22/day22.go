package day22

import (
	"strings"

	"github.com/kvrhdn/advent-of-code/advent-of-code-2017/day22/grid"
)

func SolvePart1(input string) interface{} {
	g := readInput(input)
	c := &VirusCarrier{
		pos:       g.Center(),
		dir:       grid.North,
		virusImpl: virusV1,
	}

	return infectionsAfterBursts(10_000, g, c)
}

func SolvePart2(input string) interface{} {
	g := readInput(input)
	c := &VirusCarrier{
		pos:       g.Center(),
		dir:       grid.North,
		virusImpl: virusV2,
	}

	return infectionsAfterBursts(10_000_000, g, c)
}

const (
	Clean    = 0
	Infected = 1
)

func readInput(input string) grid.InfiniteGrid {
	infGrid := grid.New()

	for y, l := range strings.Split(input, "\n") {
		for x, r := range l {
			if r == '#' {
				infGrid.Set(grid.Pos{x, y}, Infected)
			} else {
				infGrid.Set(grid.Pos{x, y}, Clean)
			}
		}
	}

	return infGrid
}

func infectionsAfterBursts(bursts int, g grid.InfiniteGrid, c *VirusCarrier) (infections int) {
	for i := 0; i < bursts; i++ {
		newState := c.process(g.Get(c.pos))

		g.Set(c.pos, newState)
		if newState == Infected {
			infections += 1
		}

		c.step()
	}
	return
}

type VirusCarrier struct {
	pos       grid.Pos
	dir       grid.Dir
	virusImpl func(*VirusCarrier, int) int
}

func (c *VirusCarrier) process(current int) (new int) {
	return c.virusImpl(c, current)
}

func (c *VirusCarrier) step() {
	c.pos = grid.Step(c.pos, c.dir)
}

func virusV1(c *VirusCarrier, current int) (result int) {
	switch current {
	case Clean:
		c.dir = grid.LeftOf(c.dir)
		result = Infected

	case Infected:
		c.dir = grid.RightOf(c.dir)
		result = Clean
	}
	return
}

const (
	Weakened = 2
	Flagged  = 3
)

func virusV2(c *VirusCarrier, current int) (result int) {
	switch current {
	case Clean:
		c.dir = grid.LeftOf(c.dir)
		result = Weakened

	case Weakened:
		result = Infected

	case Infected:
		c.dir = grid.RightOf(c.dir)
		result = Flagged

	case Flagged:
		c.dir = grid.ReverseOf(c.dir)
		result = Clean
	}
	return
}

var input = `...###.#.#.##...##.#..##.
.#...#..##.#.#..##.#.####
#..#.#...######.....#####
.###.#####.#...#.##.##...
.#.#.##......#....#.#.#..
....##.##.#..##.#...#....
#...###...#.###.#.#......
..#..#.....##..####..##.#
#...#..####.#####...#.##.
###.#.#..#..#...##.#..#..
.....##..###.##.#.....#..
#.....#...#.###.##.##...#
.#.##.##.##.#.#####.##...
##.#.###..#.####....#.#..
#.##.#...#.###.#.####..##
#.##..#..##..#.##.####.##
#.##.#....###.#.#......#.
.##..#.##..###.#..#...###
#..#.#.#####.....#.#.#...
.#####..###.#.#.##..#....
###..#..#..##...#.#.##...
..##....##.####.....#.#.#
..###.##...#..#.#####.###
####.########.#.#..##.#.#
#####.#..##...####.#..#..`
