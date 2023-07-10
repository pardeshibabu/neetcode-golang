package bst

type TimeStamp struct {
	Value          string
	TimeStampValue int
}

type TimeMap struct {
	HashMap map[string][]TimeStamp
}

func Constructor() TimeMap {
	obj := TimeMap{}
	obj.HashMap = make(map[string][]TimeStamp)
	return obj
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if len(this.HashMap) == 0 {
		obj := TimeStamp{}
		obj.TimeStampValue = -1
		obj.Value = ""
		this.HashMap[key] = append(this.HashMap[key], obj)
	}
	obj := TimeStamp{}
	obj.TimeStampValue = timestamp
	obj.Value = value
	this.HashMap[key] = append(this.HashMap[key], obj)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if len(this.HashMap[key]) == 0 {
		return ""
	}
	if this.HashMap[key][len(this.HashMap[key])-1].TimeStampValue <= timestamp {
		return this.HashMap[key][len(this.HashMap[key])-1].Value
	}
	return searchValue(this.HashMap[key], timestamp)
}

func searchValue(mat []TimeStamp, target int) string {
	lm, rm := 0, len(mat)-1
	for lm <= rm {
		mid_m := (rm + lm) / 2
		// fmt.Println(mid_m)
		if mat[mid_m].TimeStampValue == target {
			return mat[mid_m].Value
		} else if mat[mid_m].TimeStampValue > target {
			rm = mid_m - 1
		} else {
			lm = mid_m + 1
		}
	}
	if lm-1 < 0 {
		return ""
	}
	return mat[lm-1].Value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
