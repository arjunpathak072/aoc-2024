package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Part One:", partOne([]uint64{0, 1, 5, 4, 3, 0}, 729))
	fmt.Println("Part Two:", partTwo([]uint64{0, 3, 5, 4, 3, 0}))
}

func partTwo(program []uint64) (seed uint64) {
	for itr := len(program) - 1; itr >= 0; itr-- {
		seed <<= 3
		for !slices.Equal(partOne(program, seed), program[itr:]) {
			seed++
		}
	}
	return
}

func partOne(program []uint64, seed uint64) (res []uint64) {
	var instructionPointer int
	registers := map[rune]uint64{
		'A': seed,
		'B': 0,
		'C': 0,
	}

	for instructionPointer < len(program)-1 {
		operand := program[instructionPointer+1]

		switch operator := program[instructionPointer]; operator {
		case 0: // adv
			registers['A'] >>= getValue(operand, registers)
		case 1: // bxl
			registers['B'] ^= operand
		case 2: // bst
			registers['B'] = getValue(operand, registers) & 7
		case 3: // jnz
			if registers['A'] != 0 {
				instructionPointer = int(operand)
				continue
			}
		case 4: // bxc
			registers['B'] ^= registers['C']
		case 5: // out
			val := getValue(operand, registers) & 7
			res = append(res, val)
		case 6: // bdv
			registers['B'] = registers['A'] >> getValue(operand, registers)
		case 7: // cdv
			registers['C'] = registers['A'] >> getValue(operand, registers)
		}
		instructionPointer += 2
	}

	return
}

func getValue(comboOperand uint64, registers map[rune]uint64) (value uint64) {
	if comboOperand >= 0 && comboOperand <= 3 {
		value = comboOperand
	} else if comboOperand == 4 {
		value = registers['A']
	} else if comboOperand == 5 {
		value = registers['B']
	} else if comboOperand == 6 {
		value = registers['C']
	} else {
		panic("Invalid combo operand!")
	}
	return
}
