package main

func distinctNames(ideas []string) (ans int64) {
	group := [26]map[string]bool{}
	for i := range group {
		group[i] = map[string]bool{}
	}
	for _, s := range ideas {
		group[s[0]-'a'][s[1:]] = true // 按照首字母分组, 如果后缀一致
	}

	for i, a := range group { // 枚举所有组对
		// 首字母不同的字符串意味着可以相互组合
		// 1. 相同前缀的互换没有新的单词
		// 2. 相同后缀前缀不同的互换也没有新词
		for _, b := range group[:i] {
			m := 0 // 交集的大小
			for s := range a {
				// 去掉相同后缀的字符串
				if b[s] {
					m++
				}
			}
			ans += int64(len(a)-m) * int64(len(b)-m)
		}
	}
	return ans * 2 // 乘 2 放到最后
}
