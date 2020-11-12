package sllist

func From(values ...interface{}) *List{
	list := (&List{}).Init()
	for _, v := range values {
		list.append(v)
	}
	return list
}

