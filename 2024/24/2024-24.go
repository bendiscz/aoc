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

type gate struct {
	op       string
	in1, in2 string
	out      string
	done     bool
}

func solve(p *Problem) {
	s1, s2 := 0, 0

	//g := grid{NewMatrix[cell](Rectangle(len(p.PeekLine()), 0))}

	//s := p.ReadAll()

	wires := map[string]bool{}
	gates := []*gate(nil)
	for p.NextLine() && p.Line() != "" {
		f := SplitFieldsDelim(p.Line(), ": ")
		wires[f[0]] = f[1] == "1"
	}
	for p.NextLine() {
		f := SplitFieldsDelim(p.Line(), " ->")
		gates = append(gates, &gate{
			op:  f[1],
			in1: f[0],
			in2: f[2],
			out: f[3],
		})
	}

	for {
		done := true
		for _, g := range gates {
			if g.done {
				continue
			}

			w1, ok1 := wires[g.in1]
			w2, ok2 := wires[g.in2]
			if !ok1 || !ok2 {
				continue
			}

			var out bool
			switch g.op {
			case "AND":
				out = w1 && w2
			case "OR":
				out = w1 || w2
			case "XOR":
				out = w1 != w2
			}
			g.done = true
			wires[g.out] = out
			done = false
		}
		if done {
			break
		}
	}

	zw := []string(nil)
	for w := range wires {
		if w[0] == 'z' {
			zw = append(zw, w)
		}
	}
	slices.Sort(zw)

	for i, w := range zw {
		if wires[w] {
			s1 |= 1 << i
		}
	}
	p.PartOne(s1)

	if p.Example() {
		return
	}

	swaps := map[string]string{
		"z08": "vvr",
		"vvr": "z08",

		"bkr": "rnq",
		"rnq": "bkr",

		"z28": "tfb",
		"tfb": "z28",

		"z39": "mqh",
		"mqh": "z39",
	}
	_ = swaps

	for _, g := range gates {
		applySwap(swaps, &g.out)
	}

	adders := []adder(nil)
	for i := 0; i < 45; i++ {
		if i == 0 {
			sout := findGate(gates, "x00", "y00", "XOR")
			cout := findGate(gates, "x00", "y00", "AND")
			adders = append(adders, adder{sout: sout.out, cout: cout.out})
			continue
		}

		a := adder{}
		smid := findGate(gates, nWire("x", i), nWire("y", i), "XOR")
		cmid := findGate(gates, nWire("x", i), nWire("y", i), "AND")
		a.smid = smid.out
		a.cmid = cmid.out

		sout := findGate(gates, smid.out, adders[i-1].cout, "XOR")
		if sout == nil || sout.out[0] != 'z' {
			panic(fmt.Sprintf("sout %d", i))
		}
		a.sout = sout.out

		cadd := findGate(gates, adders[i-1].cout, a.smid, "AND")
		if cadd == nil {
			panic(fmt.Sprintf("cadd %d", i))
		}
		a.cadd = cadd.out

		cout := findGate(gates, a.cmid, a.cadd, "OR")
		if cout == nil {
			panic(fmt.Sprintf("cout %d", i))
		}
		a.cout = cout.out

		adders = append(adders, a)
	}

	p.PartTwo(s2)
	p.PartTwo(strings.Join(slices.Sorted(maps.Keys(swaps)), ","))
}

func applySwap(swaps map[string]string, s *string) {
	if x, ok := swaps[*s]; ok {
		*s = x
	}
}

func nWire(x string, num int) string {
	return fmt.Sprintf("%s%02d", x, num)
}

func findGate(gates []*gate, w1, w2, op string) *gate {
	for _, g := range gates {
		if g.op != op {
			continue
		}
		if g.in1 == w1 && g.in2 == w2 || g.in2 == w1 || g.in1 == w2 {
			return g
		}
	}
	return nil
}

type halfAdd struct {
	cout string
	sout string
}

type adder struct {
	cmid string
	smid string
	cout string
	cadd string
	sout string
}

// z08, vvr
// bkr, rnq
// tfb, z28
// mqh, z39

/*

y16 XOR x16 -> bkr
y16 AND x16 -> rnq







x09 XOR y09 -> ggf
y09 AND x09 -> nbb









x00 AND y00 -> qvn
y00 XOR x00 -> z00



y01 XOR x01 -> sgr
y01 AND x01 -> kbq

qvn XOR sgr -> z01

qvn AND sgr -> drw
drw OR kbq -> btn

*/
