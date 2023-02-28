# Learn Go with Tests

책 [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/) 을 학습합니다.  
Repo : https://github.com/quii/learn-go-with-tests/releases  

---

### Test 생성

```go
import "testing" // test packages

func TestAdder(t *testing.T) { // Method 이름이 Test으로 시작
	sut := Add(2, 2)
	expected := 4

	if sut != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sut)
	}
}
```

```bash
# test 실행
go test
```


### 문서 생성, godoc
Go 테스트 코드 중 Example 메서드 시작으로 문서를 작성할 수 있다.

```go
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

```

```bash
# godoc 생성
godoc -http=localhost:6060
```

### BenchMark 
Go 테스트 패키지에서 밴치마크를 손쉽게 지원한다.   
BenchMark 메서드 이름으로 시작하면 된다.
```go
import "testing"

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
```

```go
# 벤치마크 테스트 실행
> go test -bench=.

goos: darwin
goarch: amd64
pkg: learn-go-with-tests/iteration
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkRepeat-12       9213703               122.8 ns/op
PASS
ok      learn-go-with-tests/iteration   1.502s


```

### Test coverage
go test -cover 명령을 통해 커버리지을 확인할 수 있다.   
```bash
> go test -cover

PASS
        learn-go-with-tests/arraysAndSlices     coverage: 100.0% of statements
ok      learn-go-with-tests/arraysAndSlices     0.247s

```

---

[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fku-kim%2Flearn-go-with-tests&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)