package company

// Company struct
type Company struct {
	Version int

	// Owning objects
	Machines []*Machine
}

// CreateMachine method
func (company *Company) CreateMachine(MachineName string, MachineType byte) *Machine {
	machine := &Machine{
		MachineName: MachineName,
		MachineType: MachineType,
		Owner:       company,
	}

	// Add it to company unsorted
	company.Machines = append(company.Machines, machine)

	return machine
}
