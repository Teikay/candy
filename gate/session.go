package gate

import (
	"sync"

	"github.com/dearcode/candy/util/log"
)

type session struct {
	user  int64         // 用户ID
	conns []*connection // 来自不同设备的所有连接
	sync.Mutex
}

func newSession(id int64, c *connection) *session {
	return &session{user: id, conns: []*connection{c}}
}

func (s *session) addConnection(conn *connection) {
	log.Debugf("%d session:%v add conn:%+v", s.user, s, conn)
	s.Lock()
	s.conns = append(s.conns)
	s.Unlock()
}

// delConnection 遍历session的conns，删除当前连接
func (s *session) delConnection(conn *connection) bool {
	log.Debugf("%d session:%v del conn:%+v", s.user, s, conn)
	empty := false
	s.Lock()
	for i := 0; i < len(s.conns); {
		if s.conns[i].getAddr() == conn.getAddr() {
			copy(s.conns[i:], s.conns[i+1:])
			s.conns = s.conns[:len(s.conns)-1]
			continue
		}
		i++
	}
	if len(s.conns) == 0 {
		empty = true
	}
	s.Unlock()

	return empty
}

//walkConnection 复制遍历
func (s *session) walkConnection(call func(*connection) bool) {
	log.Debugf("%d walkConnection", s.user)
	s.Lock()
	conns := append([]*connection{}, s.conns...)
	s.Unlock()

	for _, c := range conns {
		if call(c) {
			break
		}
	}
}
