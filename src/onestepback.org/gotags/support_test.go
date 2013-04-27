package main

type fauxRecorder struct {
	name, def, names string
	loc Location
	count int
}
func (self *fauxRecorder) Add(tagname, defstring string, loc Location) {
	self.count += 1
	self.name = tagname
	self.def = defstring
	self.loc = loc
	self.names = self.names + "<" + tagname + ">"
}
