package multierror

import "go.uber.org/multierr"

var (
	Join = multierr.Combine
)
