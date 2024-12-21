package asm

type Regs [6]int

type Fn func(a, b, c int, r Regs) Regs

var FnByName = map[string]Fn{
	"addr": func(a, b, c int, r Regs) Regs { r[c] = r[a] + r[b]; return r },
	"addi": func(a, b, c int, r Regs) Regs { r[c] = r[a] + b; return r },
	"mulr": func(a, b, c int, r Regs) Regs { r[c] = r[a] * r[b]; return r },
	"muli": func(a, b, c int, r Regs) Regs { r[c] = r[a] * b; return r },
	"banr": func(a, b, c int, r Regs) Regs { r[c] = r[a] & r[b]; return r },
	"bani": func(a, b, c int, r Regs) Regs { r[c] = r[a] & b; return r },
	"borr": func(a, b, c int, r Regs) Regs { r[c] = r[a] | r[b]; return r },
	"bori": func(a, b, c int, r Regs) Regs { r[c] = r[a] | b; return r },
	"setr": func(a, b, c int, r Regs) Regs { r[c] = r[a]; return r },
	"seti": func(a, b, c int, r Regs) Regs { r[c] = a; return r },
	"gtir": func(a, b, c int, r Regs) Regs { r[c] = gt(a, r[b]); return r },
	"gtri": func(a, b, c int, r Regs) Regs { r[c] = gt(r[a], b); return r },
	"gtrr": func(a, b, c int, r Regs) Regs { r[c] = gt(r[a], r[b]); return r },
	"eqir": func(a, b, c int, r Regs) Regs { r[c] = eq(a, r[b]); return r },
	"eqri": func(a, b, c int, r Regs) Regs { r[c] = eq(r[a], b); return r },
	"eqrr": func(a, b, c int, r Regs) Regs { r[c] = eq(r[a], r[b]); return r },
}

func gt(a, b int) int {
	if a > b {
		return 1
	} else {
		return 0
	}
}

func eq(a, b int) int {
	if a == b {
		return 1
	} else {
		return 0
	}
}

type Inst struct {
	Name    string
	A, B, C int
}

func (i Inst) Run(r Regs) Regs {
	return FnByName[i.Name](i.A, i.B, i.C, r)
}
