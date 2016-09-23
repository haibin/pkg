package pkg

type Foo struct{
    verbosity int
}

type option func(*Foo) option

// Option sets the options specified.
// It returns an option to restore the last arg's previous value.
func (f *Foo) Option(opts ...option) (previous option) {
    for _, opt := range opts {
        previous = opt(f)
    }
    return previous
}

// Verbosity sets Foo's verbosity level to v.
func Verbosity(v int) option {
    return func(f *Foo) option {
        previous := f.verbosity
        f.verbosity = v
        return Verbosity(previous)
    }
}