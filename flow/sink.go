package flow

import "io"

// Sink is a writable location for output.
type Sink io.WriteCloser
