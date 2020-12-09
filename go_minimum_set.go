package go_minimum_set

type costque struct{
	key int64
	num float64
}

type MinSet struct{
	que map[int64][]costque
	last_num int64
	n int64
}

func NewMinSet() *MinSet{
	s := new(MinSet)
	s.last_num=0
	s.que = map[int64][]costque{}
	s.n=0
	return s
}

func (s *MinSet) Add_val(key int64,val float64){
	var t costque
	t.key = key
	t.num = val
	s.n++
	s.que[int64(val)] = append(s.que[int64(val)],t)
}

func (s *MinSet) Get_min()(int64){
	if _, ok := s.que[s.last_num]; !ok {
		s.last_num = 1<<30
		for k,_ := range s.que{
			if k < s.last_num && len(s.que[k]) > 0 {
				s.last_num = k
			}
		}
		if s.last_num == -1{
			return -1
		}	
	}
	min_id := 0
/*	min_cost := s.que[s.last_num][0]

	for i:=1;i<len(s.que[s.last_num]);i++{
		if s.que[s.last_num][i].num < min_cost.num{
			min_cost = s.que[s.last_num][i]
			min_id = i
		}
	}/**/

	ans := s.que[s.last_num][min_id].key

	s.que[s.last_num][min_id] = s.que[s.last_num][0]
	s.que[s.last_num] = s.que[s.last_num][1:]

	if len(s.que[s.last_num])==0{
		delete(s.que,s.last_num)
	}
	s.n--
	return ans
}

func (s *MinSet) Get_mins()([]int64){
	if _, ok := s.que[s.last_num]; !ok {
		s.last_num = 1<<30
		for k,_ := range s.que{
			if k < s.last_num && len(s.que[k]) > 0 {
				s.last_num = k
			}
		}
		if s.last_num == -1{
			return []int64{-1}
		}	
	}
	min_costs := s.que[s.last_num]

	mins := []int64{}
	for _,v := range min_costs{
		mins = append(mins,v.key)
	}
	delete(s.que,s.last_num)
	return mins
}

func (s *MinSet) Len()(int64){
	return s.n
}