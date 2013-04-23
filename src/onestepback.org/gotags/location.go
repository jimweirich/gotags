package main

type Location struct {
	LineCount, ByteCount int
}

func NewLocation() Location {
	return Location {
		ByteCount: 0,
		LineCount: 1,
	}
}

func (self Location) Bump(line string) Location {
	return Location {
		ByteCount: self.ByteCount + len(line),
		LineCount: self.LineCount + 1,
	}
}
