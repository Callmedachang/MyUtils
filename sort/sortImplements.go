package sort

import (
	"sort"
	"zcm_tools/amqp"
)
type Player struct {

	name string

	level int

}

// 队伍实现

type Team []*Player

func (t *Team) Join(p *Player) {
	*t = append(*t, p)
}

func (t *Team) Quit(p *Player) {
	for i, v := range *t {
		if v.name == p.name {
			copy((*t)[i:], (*t)[i+1:])
			*t = (*t)[:len(*t)-1]
			return
		}
	}
}

func (t *Team) Sort() {
	sort.Sort(sortTeam(*t))
}

// 排序实现

type sortTeam Team

func (t sortTeam) Len() int {
	return len(t)
}

func (t sortTeam) Swap(i, j int) {

	t[i], t[j] = t[j], t[i]

}

func (t sortTeam) Less(i, j int) bool {

	return t[i].level < t[j].level

}
