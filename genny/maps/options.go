package maps

import "errors"

// Options for a sync.Map implementation
type Options struct {
	Maps []Map
}

// Validate that options are usuable
func (opts *Options) Validate() error {
	if len(opts.Maps) == 0 {
		return errors.New("you must supply at least one map")
	}

	for _, m := range opts.Maps {
		if err := (&m).Validate(); err != nil {
			return err
		}
	}
	return nil
}
