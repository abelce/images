package memory

import (
	"admin/session"
	"sync"
	"container/list"
	"time"
)

var pder = &Provider{list: list.New()}

// SessionInit(sid string) (Session, error)
// SessionRead(sid string) (Session, error)
// SessionDestroy (sid string) error
// SessionGC(maxLifeTime int64)
type Provider struct {
	lock sync.Mutex;
	sessions map[string]*list.Element;
	list *list.List
}

// Set(key string, value interface{})error
// Get(key string) interface{}
// Destory(key string) error
// SessionID() string
type SessionStore struct {
	sid string;
	value map[interface{}]interface{};
	timeAccessed time.Time;
}


func (sess *SessionStore) Set(key, value interface{}) error {
	sess.value[key] = value;
	pder.SessionUpdate(sess.sid);
	return nil;	
}

func (sess *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(sess.sid);
	if value, ok := sess.value[key]; ok {
		return value;
	}
	return nil;
}

func (sess *SessionStore) Delete(key interface{}) error {
	delete(sess.value, key);	
	pder.SessionUpdate(sess.sid);
	return nil;
}

func (sess *SessionStore) SessionID() string {
	return sess.sid;
}

func (pder *Provider) SessionInit(sid string) (session.Session, error) {
	pder.lock.Lock();
	defer pder.lock.Unlock();
	v := make(map[interface{}]interface{}, 0);
	newSess := &SessionStore{
		sid: sid,
		timeAccessed: time.Now(),
		value: v,
	}

	element := pder.list.PushBack(newSess);
	pder.sessions[sid] = element;
	return newSess, nil
}

func (pder *Provider) SessionRead(sid string) (session.Session, error) {
	if element , ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := pder.SessionInit(sid);
		return sess, err;
	}
	return nil, nil;
}

func (pder *Provider) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid);
		pder.list.Remove(element);
		return nil;
	}
	return nil;
}

func (pder *Provider) SessionGC(maxlifetime int64) {
	pder.lock.Lock();
	defer pder.lock.Unlock();

	for {
		element := pder.list.Back();
		if element == nil {
			return;
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
			pder.list.Remove(element);
			delete(pder.sessions, element.Value.(*SessionStore).sid);
		} else {
			break;
		}
	}
}

func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock();
	defer pder.lock.Unlock();

	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now();
		pder.list.MoveToFront(element);
		return nil;
	}
	return nil;
}


func init(){
	pder.sessions = make(map[string]*list.Element, 0)
	session.Register("memory", pder)
}