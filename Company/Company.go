package Company

type Company struct {
	Version int
	
	// Owning objects
	Machines[] *machine
}

func (company *Company) CreateMachine(MachineName string, MachineType byte) *machine {
	machine := &machine{
		MachineName: MachineName,
		MachineType: MachineType,
		Owner: company,
	}
	
	// Add it to company unsorted
	company.Machines = append(company.Machines, machine)
	
	return machine
}