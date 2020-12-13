package day8

import (
	"errors"
	"fmt"
	aoc "github.com/TipsyPixie/advent-of-code-2020"
	"strconv"
	"strings"
)

type opcode string

const (
	NOP = opcode("nop")
	ACC = opcode("acc")
	JMP = opcode("jmp")
)

type operation struct {
	operator opcode
	operand  int
}

type stateMachine struct {
	operations            []operation
	programCounter        int
	accumulator           int
	processedOperationIds map[int]bool
}

func (thisMachine *stateMachine) proceed() (bool, error) {
	if thisMachine.programCounter > len(thisMachine.operations)-1 {
		return false, nil
	}
	currentOperation, nextProgramCounter := thisMachine.operations[thisMachine.programCounter], thisMachine.programCounter+1
	switch currentOperation.operator {
	case NOP:
	case ACC:
		thisMachine.accumulator += currentOperation.operand
	case JMP:
		nextProgramCounter = thisMachine.programCounter + currentOperation.operand
	default:
		return false, fmt.Errorf("invalid opcode %s", currentOperation.operator)
	}
	thisMachine.processedOperationIds[thisMachine.programCounter] = true
	thisMachine.programCounter = nextProgramCounter
	return true, nil
}

func (thisMachine *stateMachine) isOnSecondRun() bool {
	return thisMachine.processedOperationIds[thisMachine.programCounter]
}

func (thisMachine *stateMachine) clone() *stateMachine {
	cloneMachine := stateMachine{
		operations:            make([]operation, 0, len(thisMachine.operations)),
		programCounter:        thisMachine.programCounter,
		accumulator:           thisMachine.accumulator,
		processedOperationIds: map[int]bool{},
	}
	for _, op := range thisMachine.operations {
		cloneMachine.operations = append(cloneMachine.operations, op)
	}
	for operationId := range thisMachine.processedOperationIds {
		cloneMachine.processedOperationIds[operationId] = true
	}
	return &cloneMachine
}

func (thisMachine *stateMachine) getNextOperator() opcode {
	return thisMachine.operations[thisMachine.programCounter].operator
}

func (thisMachine *stateMachine) setNextOperator(operator opcode) *stateMachine {
	thisMachine.operations[thisMachine.programCounter].operator = operator
	return thisMachine
}

func appendOperation(operations *[]operation, opLine string) error {
	splitLine := strings.SplitN(opLine, " ", 2)
	operatorText, operandText := splitLine[0], splitLine[1]
	operand, err := strconv.Atoi(operandText)
	if err != nil {
		return err
	}
	*operations = append(*operations, operation{
		operator: opcode(operatorText),
		operand:  operand,
	})
	return nil
}

func setupStateMachine(inputPath string) (*stateMachine, error) {
	input, err := aoc.FromFile(inputPath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = input.Close() }()

	machine := stateMachine{
		operations:            make([]operation, 0, 256),
		programCounter:        0,
		accumulator:           0,
		processedOperationIds: map[int]bool{},
	}
	for opLine, ok, err := input.ReadLine(); ok || err != nil; opLine, ok, err = input.ReadLine() {
		if err != nil {
			return nil, err
		}
		err = appendOperation(&machine.operations, opLine)
		if err != nil {
			return nil, err
		}
	}
	return &machine, nil
}

func solvePart1(inputPath string) (int, error) {
	machine, err := setupStateMachine(inputPath)
	if err != nil {
		return 0, nil
	}
	for ok, err := machine.proceed(); ok || err != nil; ok, err = machine.proceed() {
		if err != nil {
			return 0, err
		}
		if machine.isOnSecondRun() {
			return machine.accumulator, nil
		}
	}
	return 0, errors.New("second run not found")
}

func solvePart2(inputPath string) (int, error) {
	machines := make([]*stateMachine, 0, 64)
	opSwap := map[opcode]opcode{NOP: JMP, JMP: NOP}

	primitiveMachine, err := setupStateMachine(inputPath)
	if err != nil {
		return 0, nil
	}
	machines = append(machines, primitiveMachine)
	swapOpAndClone := func(m *stateMachine) {
		if m.programCounter == len(m.operations) {
			return
		}
		if _, swappableNextOperator := opSwap[m.getNextOperator()]; swappableNextOperator {
			cloneMachine := m.clone().setNextOperator(opSwap[m.getNextOperator()])
			_, _ = cloneMachine.proceed()
			machines = append(machines, cloneMachine)
		}
	}
	swapOpAndClone(primitiveMachine)

	for i := 0; i < len(machines); i++ {
		machine := machines[i]
		for ok, err := machine.proceed(); ok || err != nil; ok, err = machine.proceed() {
			if err != nil {
				return 0, err
			}
			if machine.isOnSecondRun() {
				break
			}
			swapOpAndClone(machine)
		}
		if !machine.isOnSecondRun() {
			return machine.accumulator, nil
		}
	}
	return 0, errors.New("not found")
}
