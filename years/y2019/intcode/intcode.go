package intcode

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
)

type Intcode struct {
    Memory []int
    InstructionPtr int
    Code int
    Err error
    input, output []int
}

func (ic *Intcode) PopOutput() (int, error) {
    if len(ic.output) == 0 {
        return -1, errors.New("output stack had no values")
    }
    out := ic.output[len(ic.output)-1]
    ic.output = ic.output[:len(ic.output)-1]
    return out, nil
}

func (ic *Intcode) GetLastOutput() int {
    return ic.output[len(ic.output)-1]
}

func (ic *Intcode) Scan() bool {
    if ic.InstructionPtr < 0 || ic.InstructionPtr >= len(ic.Memory) {
        return false
    }
    ic.Code = ic.Memory[ic.InstructionPtr]
    ic.InstructionPtr++
    return true
}

func (ic *Intcode) Text() int {
    return ic.Code
}

func (ic *Intcode) Add(parameterMode int) {
    ic.Scan(); left := ic.Text()
    ic.Scan(); right := ic.Text()
    ic.Scan(); out := ic.Text()

    leftMode := parameterMode % 10
    rightMode := (parameterMode / 10) % 10

    if leftMode == 0 { left = ic.Memory[left] }
    if rightMode == 0 { right = ic.Memory[right] }

    ic.Memory[out] = left + right
}

func (ic *Intcode) Multiply(parameterMode int) {
    ic.Scan(); left := ic.Text()
    ic.Scan(); right := ic.Text()
    ic.Scan(); out := ic.Text()
    leftMode := parameterMode % 10
    rightMode := (parameterMode / 10) % 10

    if leftMode == 0 { left = ic.Memory[left] }
    if rightMode == 0 { right = ic.Memory[right] }

    ic.Memory[out] = left * right 
}

func (ic *Intcode) Input() {
    ic.Scan(); dest := ic.Text()
    ic.Memory[dest] = ic.input[0]
    ic.input = ic.input[1:]
}

func (ic *Intcode) FeedInput(val int) {
    ic.input = append(ic.input, val)
}

func (ic *Intcode) Output() {
    ic.Scan(); out := ic.Text()
    ic.output = append(ic.output, ic.Memory[out])
}

func (ic *Intcode) JumpIfTrue(parameterMode int) {  
    ic.Scan(); left := ic.Text()
    ic.Scan(); right := ic.Text()
    leftMode := parameterMode % 10
    rightMode := (parameterMode / 10) % 10

    if leftMode == 0 { left = ic.Memory[left] }
    if rightMode == 0 { right = ic.Memory[right] }

    if left != 0 { ic.InstructionPtr = right }
}

func (ic *Intcode) JumpIfFalse(parameterMode int) {  
    ic.Scan(); left := ic.Text()
    ic.Scan(); right := ic.Text()
    leftMode := parameterMode % 10
    rightMode := (parameterMode / 10) % 10

    if leftMode == 0 { left = ic.Memory[left] }
    if rightMode == 0 { right = ic.Memory[right] }

    if left == 0 { ic.InstructionPtr = right }
}

func (ic *Intcode) Less(parameterMode int) {
    ic.Scan(); left := ic.Text()
    ic.Scan(); right := ic.Text()
    ic.Scan(); out := ic.Text()

    leftMode := parameterMode % 10
    rightMode := (parameterMode / 10) % 10

    if leftMode == 0 { left = ic.Memory[left] }
    if rightMode == 0 { right = ic.Memory[right] }

    if left < right { ic.Memory[out] = 1 } else {
        ic.Memory[out] = 0
    }
}

func (ic *Intcode) Equal(parameterMode int) {
    ic.Scan(); left := ic.Text()
    ic.Scan(); right := ic.Text()
    ic.Scan(); out := ic.Text()

    leftMode := parameterMode % 10
    rightMode := (parameterMode / 10) % 10

    if leftMode == 0 { left = ic.Memory[left] }
    if rightMode == 0 { right = ic.Memory[right] }

    if left == right { ic.Memory[out] = 1 } else {
        ic.Memory[out] = 0 
    }
}

func GetProgram(scanner *bufio.Scanner) Intcode {
    program := make([]int, 0)
    programExp := regexp.MustCompile("-?[0-9]+")
    
    for scanner.Scan() {
        line := scanner.Text()
        values := programExp.FindAllString(line, -1)
        for _, value := range values {
            val, _ := strconv.Atoi(value)
            program = append(program, val)
        }
    }
    return Intcode{Memory: program}
}

func (ic *Intcode) Run(noun, verb int) int {
    ic.Memory[1], ic.Memory[2] = noun, verb

    return ic.RunWithoutInput()
}


func (ic *Intcode) RunWithoutInput() int {
    for ic.Scan() {
        opcode := ic.Text()
        if opcode == 99 { break }
        switch opcode%100 {
        case 1:
            ic.Add(opcode / 100)
        case 2: 
            ic.Multiply(opcode / 100)
        case 3: 
            ic.Input()
        case 4: 
            ic.Output()
        case 5: 
            ic.JumpIfTrue(opcode / 100)
        case 6:
            ic.JumpIfFalse(opcode / 100)
        case 7:
            ic.Less(opcode / 100)
        case 8:
            ic.Equal(opcode / 100)
        default:
            os.Exit(1)
        }
    }
    return ic.Memory[0]
}
