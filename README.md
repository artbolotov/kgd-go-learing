# kgd-go-learing

[Learn go with tests](https://quii.gitbook.io/learn-go-with-tests/) exercises repository template.

# About me

Short bio and motivation in learning golang.

# Learned lessons

### arraysgo - 36.4%

<details>
  <!-- <summary><code>func Sum(numbers []int) int</code></summary> -->

</details>

<details>
  <!-- <summary><code>func SumAllTails(numbersToSum ...[]int) []int</code></summary> -->

</details>

### hellogo - 88.9%

<details>
  <!-- <summary><code>func SayHello()</code></summary> -->

</details>

### integersgo - 100.0%

<details>
  <!-- <summary><code>func Add(x, y int) int</code></summary> -->

</details>

<details>
  <!-- <summary><code>func Multiply(x, y int) int</code></summary> -->

</details>

### concurrencygo - 60.0%

<details>
  <!-- <summary><code>func CheckWebsite(url string) bool</code></summary> -->

</details>

<details>
  <!-- <summary><code>func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool</code></summary> -->

    TYPES
    type WebsiteChecker func(string) bool
</details>

### dependencyinjection - 66.7%

<details>
  <!-- <summary><code>func Greet(writer io.Writer, name string)</code></summary> -->

</details>

<details>
  <!-- <summary><code>func MyGreeterHandler(w http.ResponseWriter, r *http.Request)</code></summary> -->

</details>

### iterationgo - 100.0%

<details>
  <!-- <summary><code>func Repeat(character string) string</code></summary> -->

</details>

<details>
  <!-- <summary><code>func SumAllNumbers(numbers ...int) int</code></summary> -->

</details>

<details>
  <!-- <summary><code>func SumPositiveNumbers(numbers ...int) int</code></summary> -->

</details>

### reflectiongo - 100.0%

### pointersgo - 54.5%
VARIABLES

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

TYPES

type Bitcoin int

func (b Bitcoin) String() string

type Dollar int

func (d Dollar) StringDollar() string

type Wallet struct {
	// Has unexported fields.
}

func (w *Wallet) Balance() Bitcoin

func (w *Wallet) Deposit(amount Bitcoin)

func (w *Wallet) Withdraw(amount Bitcoin) error

type WalletDol struct {
	// Has unexported fields.
}

func (d *WalletDol) BalanceDollar() Dollar

func (d *WalletDol) Change(w *Wallet) int
    exchange

func (d *WalletDol) Insert(amount Dollar)
    dollar
### mockgo - 55.6%

<details>
  <!-- <summary><code>func Countdown(out io.Writer, sleeper Sleeper)</code></summary> -->

    Countdown prints a countdown from 3 to out.
</details>

<details>
  <!-- <summary><code>func Mocking()</code></summary> -->

    TYPES
    type ConfigurableSleeper struct {
    // Has unexported fields.
    }
</details>

<details>
  <!-- <summary><code>func (c *ConfigurableSleeper) Sleep()</code></summary> -->

    type DefaultSleeper struct{}
</details>

<details>
  <!-- <summary><code>func (d *DefaultSleeper) Sleep()</code></summary> -->

    type Sleeper interface {
    Sleep()
    }
</details>

### selectgo - 90.0%

<details>
  <!-- <summary><code>func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error)</code></summary> -->

</details>

<details>
  <!-- <summary><code>func Racer(a, b string) (winner string, error error)</code></summary> -->

</details>

### structgo - 100.0%

<details>
  <!-- <summary><code>func Perimeter(rectangle Rectangle) float64</code></summary> -->

    TYPES
    type Circle struct {
    Radius float64
    }
</details>

<details>
  <!-- <summary><code>func (c Circle) Area() float64</code></summary> -->

    type Rectangle struct {
    Width  float64
    Height float64
    }
</details>

<details>
  <!-- <summary><code>func (r Rectangle) Area() float64</code></summary> -->

    type Shape interface {
    Area() float64
    }
    type Triangle struct {
    Base   float64
    Height float64
    }
</details>

<details>
  <!-- <summary><code>func (t Triangle) Area() float64</code></summary> -->

</details>
