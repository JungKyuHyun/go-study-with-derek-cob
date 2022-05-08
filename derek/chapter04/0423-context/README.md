# context란?

context는 하나의 맥락이고, 이 맥락을 프로세스나 API에 전달해서 모든 연계되는 작업들이 하나의 맥락안에서 이뤄지도록 하는 것이다.

# 주의할 점.

context는 struct로 갖고 있으면 안된다. Param으로 명시적으로 전달을 해야한다. 그렇다고 struct로 갖는게 안되는건 아니지만, 절대 nil을 전달하면 안된다.

# context의 구성.

`type Context interface {
// Done returns a channel that is closed when this Context is canceled
// or times out.
Done() <-chan struct{}

    // Err indicates why this context was canceled, after the Done channel
    // is closed.
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    Value(key interface{}) interface{}

}`

1. Done : cancel 혹은 타임아웃 되었을 때 닫힌 channel을 리턴
2. Err : 만약 Done이 닫혀있지 않다면 Err는 nil을 리턴한다. 만약 Done이 닫혀있다면 Err는 non-nil 에러를 리턴한다. 에러에는 왜 context가 cancel되었는지에 대한 설명이 존재한다.
3. Deadline : deadline이 존재할 때 주어진 context(맥락)의 마감기간을 리턴한다. 만약 마감기간이 정해져있지 않다면 ok값으로는 false를 리턴
4. Value : context에 key가 있으면, 그 key에 해당하는 값이다.

# ⭐️Cancel context

고루틴을 사용할 때 종료를 제대로 해주지 않으면 종료되지 않는 작업을 만들게 될 수 있다. 그 때, 이 고루틴의 종료를 관리해줄 수 있는 것이 바로 cancel 가능한 context인 것.
ex) 유저가 response를 받기 전에 request를 cancel해도, 서버는 그대로 request를 처리하기 위해 DB에서 데이터를 가져오고, 가공한 다음 response를 보내게 된다 ( 불필요한 작업.)

## WithCancel

parent context의 copy본과 새로운 Done channel을 리턴한다.

````go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
{ if parent == nil { panic("cannot create context from nil parent") } c := newCancelCtx(parent) propagateCancel(parent, &c) return &c, func() { c.cancel(true, Canceled) } }```
````
