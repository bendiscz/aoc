package main

import (
	"fmt"
	. "github.com/bendiscz/aoc/aoc"
	"maps"
	"slices"
	"strings"
)

const (
	year    = 2024
	day     = 24
	example = `

x00: 1
x01: 0
x02: 1
x03: 1
x04: 0
y00: 1
y01: 1
y02: 1
y03: 1
y04: 1

ntg XOR fgs -> mjb
y02 OR x01 -> tnw
kwq OR kpj -> z05
x00 OR x03 -> fst
tgd XOR rvg -> z01
vdt OR tnw -> bfw
bfw AND frj -> z10
ffh OR nrd -> bqk
y00 AND y03 -> djm
y03 OR y00 -> psh
bqk OR frj -> z08
tnw OR fst -> frj
gnj AND tgd -> z11
bfw XOR mjb -> z00
x03 OR x00 -> vdt
gnj AND wpb -> z02
x04 AND y00 -> kjc
djm OR pbm -> qhw
nrd AND vdt -> hwm
kjc AND fst -> rvg
y04 OR y02 -> fgs
y01 AND x02 -> pbm
ntg OR kjc -> kwq
psh XOR fgs -> tgd
qhw XOR tgd -> z09
pbm OR djm -> kpj
x03 XOR y03 -> ffh
x00 XOR y04 -> ntg
bfw OR bqk -> z06
nrd XOR fgs -> wpb
frj XOR qhw -> z04
bqk OR frj -> z07
y03 OR x01 -> nrd
hwm AND bqk -> z03
tgd XOR rvg -> z12
tnw OR pbm -> gnj

`
)

func main() {
	Run(year, day, example, solve)
}

/*

Full adder:

C_in--------------------------+-------+---+
                               \      +XOR >------------------S_out
                                \   +-+---+
                                 \ /
                                  X
                                 / \
                                /   +-+---+
X_i---+-------+---+            /      +AND >--C_add---+---+
       \      |XOR >--S_mid---+-------+---+           +OR  >--C_out
        \   +-+---+                                 +-+---+
         \ /                                       /
          X                                       /
         / \                                     /
        /   +-+---+                             /
       /      |AND >--C_mid--------------------+
Y_i---+-------+---+

Half adder:

X_i---+-------+---+
       \      |XOR >--S_out
        \   +-+---+
         \ /
          X
         / \
        /   +-+---+
       /      |AND >--C_out
Y_i---+-------+---+

*/

type gate struct {
	in1, in2 string
	op       string
	out      string
}

type circuit struct {
	gates []gate
	swaps map[string]string
}

func newCircuit() *circuit {
	return &circuit{swaps: map[string]string{}}
}

func (c *circuit) addGate(in1, in2, op, out string) {
	c.gates = append(c.gates, gate{in1, in2, op, out})
}

func (c *circuit) swap(out1, out2 string) {
	c.swaps[out1] = out2
	c.swaps[out2] = out1
}

func (c *circuit) findGate(in1, in2, op string) (gate, bool) {
	for _, g := range c.gates {
		if g.op != op {
			continue
		}
		if g.in1 == in1 && g.in2 == in2 || g.in1 == in2 && g.in2 == in1 {
			if out, ok := c.swaps[g.out]; ok {
				g.out = out
			}
			return g, true
		}
	}
	return gate{}, false
}

func wire(x string, num int) string {
	return fmt.Sprintf("%s%02d", x, num)
}

func solve(p *Problem) {
	wires := map[string]bool{}
	for p.NextLine() && p.Line() != "" {
		f := SplitFieldsDelim(p.Line(), ": ")
		wires[f[0]] = f[1] == "1"
	}

	c := newCircuit()
	for p.NextLine() {
		f := SplitFieldsDelim(p.Line(), " ->")
		c.addGate(f[0], f[2], f[1], f[3])
	}

	p.PartOne(solvePartOne(c, wires))

	if p.Example() {
		return
	}
	p.PartTwo(solvePartTwo(c, wires))
}

func solvePartOne(c *circuit, wires map[string]bool) int {
	for done := false; !done; {
		done = true
		for _, g := range c.gates {
			if _, ok := wires[g.out]; ok {
				continue
			}

			in1, ok1 := wires[g.in1]
			in2, ok2 := wires[g.in2]
			if !ok1 || !ok2 {
				continue
			}

			done = false

			switch g.op {
			case "AND":
				wires[g.out] = in1 && in2
			case "OR":
				wires[g.out] = in1 || in2
			case "XOR":
				wires[g.out] = in1 != in2
			}
		}
	}

	for s, i := 0, 0; ; i++ {
		if w, ok := wires[wire("z", i)]; ok {
			if w {
				s |= 1 << i
			}
		} else {
			return s
		}
	}
}

func solvePartTwo(c *circuit, wires map[string]bool) string {
loop:
	for iter := 0; iter < 1_000_000; iter++ {
		carry := ""
		for i := 0; ; i++ {
			x, y, z := wire("x", i), wire("y", i), wire("z", i)
			if _, ok := wires[x]; !ok {
				break
			}

			var ok bool
			if i == 0 {
				carry, ok = checkHalfAdder(c, x, y, z)
			} else {
				carry, ok = checkFullAdder(c, x, y, z, carry)
			}

			if !ok {
				// restart with some new swaps
				continue loop
			}
		}

		return strings.Join(slices.Sorted(maps.Keys(c.swaps)), ",")
	}

	fail("never pick up an infinite loop!")
	return ""
}

func fail(format string, a ...any) {
	panic(fmt.Sprintf(format, a...))
}

func checkHalfAdder(c *circuit, x, y, z string) (string, bool) {
	sOutGate, sOutOk := c.findGate(x, y, "XOR")
	cOutGate, cOutOk := c.findGate(x, y, "AND")
	if !sOutOk || !cOutOk {
		fail("could not find half gate for %s", z)
	}

	if sOutGate.out != z {
		if cOutGate.out != z {
			fail("broken half gate for %s", z)
		}
		c.swap(cOutGate.out, z)
		return "", false
	}

	return cOutGate.out, true
}

func checkFullAdder(c *circuit, x, y, z, carry string) (string, bool) {
	sMidGate, sMidOk := c.findGate(x, y, "XOR")
	cMidGate, cMidOk := c.findGate(x, y, "AND")
	if !sMidOk || !cMidOk {
		fail("could not find full gate for %s", z)
	}

	sOutGate, sOutOk := c.findGate(sMidGate.out, carry, "XOR")
	if !sOutOk {
		c.swap(sMidGate.out, cMidGate.out)
		return "", false
	}

	if sOutGate.out != z {
		c.swap(sOutGate.out, z)
		return "", false
	}

	cAddGate, cAddOk := c.findGate(carry, sMidGate.out, "AND")
	if !cAddOk {
		// should not happen, changing sMid would break sOut
		fail("failure at cAdd gate for %s", z)
	}

	cOutGate, cOutOk := c.findGate(cMidGate.out, cAddGate.out, "OR")
	if !cOutOk {
		// should not happen, we can swap only cMid-cAdd and both are
		// connected to cOut gate
		fail("failure at cOut gate for %s", z)
	}

	return cOutGate.out, true
}
