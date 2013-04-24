type Strategy interface {
}
type Rule struct {
}
func (self *Rule) Apply() bool {
}
func NewRule(pattern string, tagIndex, defIndex int, isMulti bool) *Rule {
}
var RubyRulesList = []*Rule {
}
