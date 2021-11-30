package main

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}

	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) > 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}

	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

type point struct {
	x, y int
}

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	////////////////////////////////
	// Collection Functions
	////////////////////////////////
	/*
		var strs = []string{"peach", "apple", "pear", "plum"}
		fmt.Println(Index(strs, "pear"))
		fmt.Println(Index(strs, "grape"))

		fmt.Println(Any(strs, func(v string) bool {
			return strings.HasPrefix(v, "p")
		}))

		fmt.Println(All(strs, func(v string) bool {
			return strings.HasPrefix(v, "p")
		}))

		fmt.Println(Filter(strs, func(v string) bool {
			return strings.Contains(v, "e")
		}))

		fmt.Println(Map(strs, strings.ToUpper))
	*/

	////////////////////////////////
	// String Functions
	////////////////////////////////
	/*
		var p = fmt.Println

		p("Contains: ", s.Contains("test", "e"))
		p("Count: ", s.Count("test", "t"))
		p("HasPrefix: ", s.HasPrefix("test", "t"))
		p("HasSuffix: ", s.HasSuffix("test", "st"))
		p("Index: ", s.Index("test", "t"))
		p("Join: ", s.Join([]string{"a", "b"}, "-"))
		p("Repeat: ", s.Repeat("a", 5))
		p("Replace: ", s.Replace("foo", "o", "0", -1)) //-1 means theres no limit to replace
		p("Replace: ", s.Replace("foo", "o", "0", 1))
		p("Split: ", s.Split("a-b-c-d-e", "-"))
		p("ToLower: ", s.ToLower("TEst"))
		p("ToUpper: ", s.ToUpper("test"))

		p("Len: ", len("hello"))
		p("Char:", "hello"[1])
	*/
	////////////////////////////////
	// String Formatting
	////////////////////////////////
	/*
		p := point{1, 2}
		fmt.Printf("struct1: %v\n", p)
		fmt.Printf("struct2: %+v\n", p)
		fmt.Printf("struct3: %#v\n", p)

		fmt.Printf("type: %T\n", p)
		fmt.Printf("bool: %t\n", true)
		fmt.Printf("int: %d\n", 123)
		fmt.Printf("bin: %b\n", 14)
		fmt.Printf("char: %c\n", 33)
		fmt.Printf("hex: %x\n", 456)

		fmt.Printf("float1: %f\n", 78.9)
		fmt.Printf("float2: %e\n", 123400000.0)
		fmt.Printf("float3: %E\n", 123400000.0)

		fmt.Printf("str1: %s\n", "\"string\"")
		fmt.Printf("str2: %q\n", "\"string\"")
		fmt.Printf("str3: %x\n", "hex this")

		fmt.Printf("pointer: %p\n", &p)

		fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
		fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
		fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

		fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
		fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

		s := fmt.Sprintf("sprintf: a %s", "string")
		fmt.Println(s)

		fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
	*/

	///////////////////////////////
	// Regex
	///////////////////////////////
	/*
		match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
		fmt.Println(match)

		r, _ := regexp.Compile("p([a-z]+)ch")

		fmt.Println(r.MatchString("peach"))
		fmt.Println(r.FindString("peach punch"))
		fmt.Println("idx:", r.FindStringIndex("peach punch"))
		fmt.Println(r.FindStringSubmatch("peach punch"))
		fmt.Println(r.FindStringSubmatch("punch peach"))
		fmt.Println(r.FindStringSubmatchIndex("peach punch"))

		fmt.Println(r.FindAllString("peach punch pinch", -1))
		fmt.Println("all:", r.FindAllStringSubmatchIndex(
			"peach punch pinch", -1))

		fmt.Println(r.FindAllString("peach punch pinch", 2))
		fmt.Println(r.Match([]byte("peach")))

		r = regexp.MustCompile("p([a-z]+)ch")
		fmt.Println("regexp:", r)

		fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

		in := []byte("a peach")
		out := r.ReplaceAllFunc(in, bytes.ToUpper)
		fmt.Println(string(out))
	*/
	////////////////////////////////
	// JSON
	////////////////////////////////
	/*

		bolB, _ := json.Marshal(true)
		fmt.Println(string(bolB))

		intB, _ := json.Marshal(1)
		fmt.Println(string(intB))

		fltB, _ := json.Marshal(2.34)
		fmt.Println(string(fltB))
		// fmt.Println(fltB_) captures error

		strB, _ := json.Marshal("gopher")
		fmt.Println(string(strB))

		slcD := []string{"apple", "peach", "pear"}
		slcB, _ := json.Marshal(slcD)
		fmt.Println(string(slcB))

		mapD := map[string]int{"apple": 5, "lettuce": 7}
		mapB, _ := json.Marshal(mapD)
		fmt.Println(string(mapB))

		res1D := &response1{
			Page:   1,
			Fruits: []string{"apple", "peach", "pear"}}
		res1B, _ := json.Marshal(res1D)
		fmt.Println(string(res1B))

		res2D := &response2{
			Page:   1,
			Fruits: []string{"apple", "peach", "pear"}}
		res2B, _ := json.Marshal(res2D)
		fmt.Println(string(res2B))

		byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
		var dat map[string]interface{}

		if err := json.Unmarshal(byt, &dat); err != nil {
			panic(err)
		}

		fmt.Println(dat)
		num := dat["num"].(float64)
		fmt.Println(num)

		strs := dat["strs"].([]interface{})
		str1 := strs[0].(string)
		fmt.Println(str1)

		str := `{"page": 1, "fruits": ["apple", "peach"]}`
		res := response2{}
		json.Unmarshal([]byte(str), &res)
		fmt.Println(res)
		fmt.Println(res.Fruits[0])

		enc := json.NewEncoder(os.Stdout)
		d := map[string]int{"apple": 5, "lettuce": 7}
		enc.Encode(d)
	*/
	////////////////////////////////
	// Time
	////////////////////////////////
	/*

		p := fmt.Println
		now := time.Now()

		p(now)

		then := time.Date(
			2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
		p(then)

		p(then.Year())
		p(then.Month())
		p(then.Day())
		p(then.Hour())
		p(then.Minute())
		p(then.Second())
		p(then.Nanosecond())
		p(then.Location())

		p(then.Weekday())

		p(then.Before(now))
		p(then.After(now))
		p(then.Equal(now))

		diff := now.Sub(then)
		p(diff)

		p(diff.Hours())
		p(diff.Minutes())
		p(diff.Seconds())
		p(diff.Nanoseconds())

		p(then.Add(diff))
		p(then.Add(-diff))
	*/

	///////////////////////////////
	// Random Numbers
	///////////////////////////////
	/*
		fmt.Print(rand.Intn(100), ",")
		fmt.Print(rand.Intn(100))
		fmt.Println()

		fmt.Println(rand.Float64())
		fmt.Print((rand.Float64()*5)+5, ",")
		fmt.Print((rand.Float64() * 5) + 5)
		fmt.Println()

		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		fmt.Print(r1.Intn(100), ",")
		fmt.Print(r1.Intn(100))
		fmt.Println()

		s2 := rand.NewSource(42)
		r2 := rand.New(s2)
		fmt.Print(r2.Intn(100), ",")
		fmt.Print(r2.Intn(100))
		fmt.Println()
		s3 := rand.NewSource(42)
		r3 := rand.New(s3)
		fmt.Print(r3.Intn(100), ",")
		fmt.Print(r3.Intn(100))
	*/

	////////////////////////////////////////////////////////////////
	// Number Parsing
	////////////////////////////////////////////////////////////////
	/*
		f, _ := strconv.ParseFloat("1.2345", 64)
		fmt.Println(f)

		i, _ := strconv.ParseInt("123", 0, 64)
		fmt.Println(i)

		d, _ := strconv.ParseInt("0x1c8", 0, 64)
		fmt.Println(d)

		u, _ := strconv.ParseUint("789", 0, 64)
		fmt.Println(u)

		k, _ := strconv.Atoi("135")
		fmt.Println(k)

		_, e := strconv.Atoi("wat")
		fmt.Println(e)
	*/

	///////////////////////////////
	// Url Parsing
	///////////////////////////////
	/*

		s := "postgres://user:pass@host.com:5432/path?k=v#f"

		u, err := url.Parse(s)
		if err != nil {
			panic(err)
		}

		fmt.Println(u)
		fmt.Println(u.Scheme)

		fmt.Println(u.User)
		fmt.Println(u.User.Username())
		p, _ := u.User.Password()
		fmt.Println(p)

		fmt.Println(u.Host)

		host, port, _ := net.SplitHostPort(u.Host)
		fmt.Println(host)
		fmt.Println(port)

		fmt.Println(u.Path)
		fmt.Println(u.Fragment)

		fmt.Println(u.RawQuery)
		m, _ := url.ParseQuery(u.RawQuery)
		fmt.Println(m)
		fmt.Println(m["k"][0])
	*/

	///////////////////////////////
	// Reading a File
	///////////////////////////////
	/*
		dat, err := os.ReadFile("/tmp/dat")
		check(err)
		fmt.Println(dat)

		f, err := os.Open("/tmp/dat")
		defer f.Close()
		check(err)

		b1 := make([]byte, 5)
		n1, err := f.Read(b1)
		check(err)
		fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

		o2, err := f.Seek(6, 0)
		check(err)
		b2 := make([]byte, 2)
		n2, err := f.Read(b2)
		check(err)
		fmt.Printf("%d bytes @ %d: ", n2, o2)
		fmt.Printf("%v\n", string(b2[:n2]))

		o3, err := f.Seek(6, 0)
		check(err)
		b3 := make([]byte, 2)
		n3, err := io.ReadAtLeast(f, b3, 2)
		check(err)
		fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

		_, err = f.Seek(0, 0)
		check(err)

		r4 := bufio.NewReader(f)
		b4, err := r4.Peek(5)
		check(err)
		fmt.Printf("5 bytes: %s\n", string(b4))
	*/
	////////////////////////////////
	// Write File
	////////////////////////////////
	/*
		d1 := []byte("hello\ngo\n")
		err := os.WriteFile("/tmp/dat1", d1, 0644)
		check(err)

		f, err := os.Create("/tmp/dat2")
		check(err)
		defer f.Close()

		d2 := []byte{115, 111, 109, 101, 10}
		n2, err := f.Write(d2)
		check(err)
		fmt.Printf("wrote %d bytes\n", n2)

		n3, err := f.WriteString("writes\n")
		check(err)
		fmt.Printf("wrote %d bytes\n", n3)

		f.Sync()

		w := bufio.NewWriter(f)
		n4, err := w.WriteString("buffered\n")
		check(err)
		fmt.Printf("wrote %d bytes\n", n4)

		w.Flush()
	*/
}
