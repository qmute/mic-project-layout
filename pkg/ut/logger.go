package ut

import (
	"fmt"
	"net/url"
	"sync"

	"github.com/cockroachdb/errors"

	"github.com/sirupsen/logrus"
)

func Recover() {
	if err := recover(); err != nil {
		msg := fmt.Sprintf("%+v", err)
		if v, ok := err.(error); ok && v != nil {
			msg = fmt.Sprintf("%+v", errors.WithStack(v))
		}

		logrus.WithField("error", msg).Errorln("panic for recover ")
	}
}

func Log(entry ...*logrus.Entry) *Lg {
	rs := logrus.WithFields(logrus.Fields{})
	if len(entry) > 0 {
		rs = entry[0]
	}
	return &Lg{
		rs: rs,
		m:  map[string]interface{}{},
	}
}

type Lg struct {
	sync.Mutex
	m  map[string]interface{}
	rs *logrus.Entry
}

func (p *Lg) WithErr(err error) *Lg {
	msg := fmt.Sprintf("%+v", err)
	p.rs = p.rs.WithField("error", msg)
	return p
}

func (p *Lg) WithRequestId(requestId string) *Lg {
	p.rs = p.rs.WithField("requestId", requestId)
	return p
}

func (p *Lg) WithOn(on string) *Lg {
	p.rs = p.rs.WithField("on", on)
	return p
}

func (p *Lg) WithFrom(from string) *Lg {
	p.rs = p.rs.WithField("from", from)
	return p
}

func (p *Lg) WithDebug(k string, v interface{}) *Lg {
	p.Lock()
	defer p.Unlock()
	p.m[k] = v
	return p
}

func (p *Lg) WithQuery(q url.Values) *Lg {
	p.rs = p.rs.WithField("query", q.Encode())
	return p
}

func (p *Lg) WithPayload(b []byte) *Lg {
	p.rs = p.rs.WithField("payload", string(b))
	return p
}

func (p *Lg) debug() *Lg {
	if len(p.m) == 0 {
		return p
	}

	p.rs = p.rs.WithField("debug", JsonString(p.m))
	return p
}

func (p *Lg) Entry() *logrus.Entry {
	return p.debug().rs
}

func (p *Lg) WithCode(code string) *Lg {
	p.rs = p.rs.WithField("code", code)
	return p
}

func (p *Lg) WithUserId(userId int) *Lg {
	p.rs = p.rs.WithField("userId", userId)
	return p
}

func (p *Lg) WithGameId(gameId int) *Lg {
	p.rs = p.rs.WithField("gameId", gameId)
	return p
}

func (p *Lg) WithClientType(clt string) *Lg {
	p.rs = p.rs.WithField("clientType", clt)
	return p
}
