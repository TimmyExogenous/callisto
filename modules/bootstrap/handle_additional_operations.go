package bootstrap

// RunAdditionalOperations implements modules.AdditionalOperationsModule
func (m *Module) RunAdditionalOperations() error {
	// Fetch all states at indexer startup.
	return m.refetchETHStates()
}
