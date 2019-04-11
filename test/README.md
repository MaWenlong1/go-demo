golang语言中的陷阱

``` go
type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	for key, value := range m {
		fmt.Println(key,value)
	}
}

```
代码的运行结果为：
```
wang &{wang 22}
zhou &{wang 22}
li &{wang 22}
```
原因的因为每次都是range遍历会复用stu，每次都是同一个stu

2. 类型
```go
type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}
```
输出
```
showA
showB
```
原因是执行了 最后需要调用People的ShowB
```go
func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
```

3. select case
```go
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "Golang我们走，我们要做好朋友！！！"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
```
如果两个case都满足条件伪随机选择一个

4. defer 
```go
func calc(index string, a, b int) int {
    ret := a + b
    fmt.Println(index, a, b, ret)
    return ret
}

func main() {
    a := 1
    b := 2
    defer calc("1", a, calc("10", a, b))
    a = 0
    defer calc("2", a, calc("20", a, b))
    b = 1
}
```
输出结果
```
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4
```
原因是defer calc("1", a, calc("10", a, b)) 的第3个参数会在调用runtime.deferproc时确定，并不会在延时调用时才会被计算。

5. map线程安全
```go
type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}
```
可能会出现fatal error: concurrent map read and map write
6. chan缓存池
```go
func (set *threadSafeSet) Iter() <-chan interface{} {
    ch := make(chan interface{})
    //  ch := make(chan interface{},len(set.s)) 可运行
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}
```
迭代会要求set.s全部可以遍历一次。但是chan是为缓存的，那就代表这写入一次就会阻塞。