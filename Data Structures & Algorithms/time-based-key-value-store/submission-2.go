type Entry struct {
	timestamp int
	value string
}

type TimeMap struct {
	keys map[string][]Entry
}

func Constructor() TimeMap {
	return TimeMap{
		keys: make(map[string][]Entry),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	entry := Entry{
		value: value, 
		timestamp: timestamp,
	}
	this.keys[key] = append(this.keys[key], entry)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if len(this.keys[key]) == 0 {
		return ""
	}

	largestTimestamp := -1
	val := ""
	left, right := 0, len(this.keys[key]) - 1
	for left <= right {
		middle := left + (right - left) / 2
		if this.keys[key][middle].timestamp == timestamp {
			return this.keys[key][middle].value
		} else if this.keys[key][middle].timestamp > timestamp {
			right = middle - 1
		} else {
			if this.keys[key][middle].timestamp > largestTimestamp {
				largestTimestamp = this.keys[key][middle].timestamp
				val = this.keys[key][middle].value
			}
			left = middle + 1
		}
	}

	return val
}
