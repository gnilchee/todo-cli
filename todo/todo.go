package todo

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	position int
	Done     bool
}

//ByPri implements sort.Interface for []Item based on
//the priority & position field.
type ByPri []Item

func (s ByPri) Len() int      { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return s[j].Done
	}
	if s[i].Priority == s[j].Priority {
		//fmt.Printf("DDD same pri %v \n", s[i].position < s[j].position)
		return s[i].position < s[j].position
	}
	return s[i].Priority > s[j].Priority // priority 1 > 2 > 3
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}

}

func (i *Item) PrettyPriority() string {
	var pri string
	switch i.Priority {
	case 1:
		pri = "(1)"
	case 3:
		pri = "(3)"
	default:
		pri = "   "
	}
	return pri
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "✔"
	} else {
		return " "
	}
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + ". "
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, nil
	}
	for i := range items {
		items[i].position = i + 1
	}
	return items, nil
}
