package pstruct

import "sort"

type  Student struct {
	ID string
	Name string
	Age int
}
type Students []Student

func (s Students) Len() int{
	return len(s)
}
func (s Students)Less(i,j int) bool{
	ids := []string{s[i].ID,s[j].ID}
	sort.Strings(ids)
	if ids[0]==s[i].ID {
		return true
	}else {
		return false
	}
}
func (s Students)Swap (i,j int){
	s[i],s[j]=s[j],s[i]
}

