package maps

// Options for a sync.Map implementation
type Options struct {
	Maps []Map
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	return nil
}
