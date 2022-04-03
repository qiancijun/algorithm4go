package week287

type Encrypter struct {
	k      []byte
	v      []string
	d      []string
	cache  map[byte]int
	cache2 map[string]byte
}

func Constructor(keys []byte, values []string, dictionary []string) Encrypter {
	encrypter := Encrypter{
		k:      keys,
		v:      values,
		d:      dictionary,
		cache:  make(map[byte]int),
		cache2: make(map[string]byte),
	}
	for _, v := range dictionary {
		encrypter.cache2[v] = 1
	}
	return encrypter
}

func (this *Encrypter) Encrypt(word1 string) string {
	ans := ""
	c := this.cache
	for i := range word1 {
		ch := word1[i]
		if idx, has := c[ch]; has {
			ans += this.v[idx]
			continue
		} else {
			// 去key寻找
			for m, n := range this.k {
				if n == ch {
					c[n] = m
					ans += this.v[m]
					break
				}
			}
		}
	}
	return ans
}

func (this *Encrypter) Decrypt(word2 string) int {
	d := this.d
	s := make(map[string]int)
	for _, v := range d {
		s[this.Encrypt(v)]++
	}
	return s[word2]
}
// func (this *Encrypter) Decrypt(word2 string) int {
// 	n := len(word2)
// 	list := [][]int{}
// 	for i, j := 2, 0; i <= n; i += 2 {
// 		substr := word2[j:i]
// 		tmp := []int{}
// 		for m, n := range this.v {
// 			if n == substr {
// 				tmp = append(tmp, m)
// 			}
// 		}
// 		list = append(list, tmp)
// 		j = i
// 	}
// 	end := len(list)
// 	s := make(map[string]byte)
// 	cnt := 0
// 	var dfs func(pos int, tmp string)
// 	dfs = func (pos int, tmp string)  {
// 		if pos == end {
// 			s[tmp] = 1
// 			return
// 		}
// 		for i := 0; i < len(list[pos]); i++ {
// 			idx := list[pos][i]
// 			dfs(pos+1, tmp+string(this.k[idx]))
// 		}
// 	}
// 	dfs(0, "")
// 	for k, _ := range s {
// 		if _, has := this.cache2[k]; has {
// 			cnt++
// 		}
// 	}
// 	return cnt
// }

/**
 * Your Encrypter object will be instantiated and called as such:
 * obj := Constructor(keys, values, dictionary);
 * param_1 := obj.Encrypt(word1);
 * param_2 := obj.Decrypt(word2);
 */