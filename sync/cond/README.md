## sync.Cond



```go
type Server struct {
	opts serverOptions

	mu sync.Mutex // guards following

	// net.Listener map
	lis map[net.Listener]bool
	// conns contains all active server transports. It is a map keyed on a
	// listener address with the value being the set of active transports
	// belonging to that listener.

	conns    map[string]map[transport.ServerTransport]bool
	serve    bool
	drain    bool
	cv       *sync.Cond              // signaled when connections close for GracefulStop
	services map[string]*serviceInfo // service name -> service info
	events   trace.EventLog

	quit               *grpcsync.Event
	done               *grpcsync.Event
	channelzRemoveOnce sync.Once
	serveWG            sync.WaitGroup // counts active Serve goroutines for GracefulStop

	channelzID int64 // channelz unique identification number
	czData     *channelzData

	serverWorkerChannels []chan *serverWorkerData
}
```

cv       *sync.Cond 定义为 *sync.Cond

// 连接数不为0等 处于Wait状态
for len(s.conns) != 0 {
		s.cv.Wait()
}


```go 
func (s *Server) Stop() {
	s.quit.Fire()

	defer func() {
		s.serveWG.Wait()
		s.done.Fire()
	}()

	s.channelzRemoveOnce.Do(func() {
		if channelz.IsOn() {
			channelz.RemoveEntry(s.channelzID)
		}
	})

	s.mu.Lock()
	listeners := s.lis
	s.lis = nil
	conns := s.conns
	s.conns = nil
	// interrupt GracefulStop if Stop and GracefulStop are called concurrently.

    // 服务关闭时候 conns置为空
    // 通知所有的协程，如果有事件
    // 触发Finish()
	s.cv.Broadcast()
	s.mu.Unlock()
    ...
}
```