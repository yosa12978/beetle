package machine

type Machine struct {
	mem []uint8
	sp  uint8
	pc  uint8
	cf  bool
	zf  bool
}

func NewMachine() *Machine {
	return &Machine{
		mem: make([]uint8, 0xFF),
		sp:  0x00,
		pc:  0x40,
		cf:  false,
		zf:  false,
	}
}

func (m *Machine) Run() {

}

func (m *Machine) LoadProgram() {

}

func (m *Machine) fetch() (op uint8) {
	op = m.mem[m.pc]
	m.pc++
	return
}

func (m *Machine) Cycle() {
	opcode := m.fetch()
	switch opcode {
	case PUSH:
		operand := m.fetch()
		m.mem[m.sp] = operand
		m.sp++
	case POP:
		m.sp--
	case DUP:
		m.mem[m.sp] = m.mem[m.sp-1]
		m.sp++
	case STORE:
		operand := m.fetch()
		m.mem[operand] = m.mem[m.sp-1]
	case LOAD:
		operand := m.fetch()
		m.mem[m.sp] = m.mem[operand]
		m.sp++
	}
}
