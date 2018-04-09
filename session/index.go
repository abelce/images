package session

import (
	"sync"
	"fmt"
	"math/rand"
	"net/url"
	"encoding/base64"
	"net/http"
	"time"
)


type Manager struct {
	cookieName string;
	provider Provider;
	maxLifeTime int64;
	lock sync.Mutex;
}

func NewManager (provideName, cookieName string, maxLifeTime int64) (*Manager, error){
	provider, ok := pdr[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unkown providename %p", provideName)
	}

	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil;
}

func (manage *Manager) SessionId() string {

	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b);
}

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {

	manager.lock.Lock();
	defer manager.lock.Unlock();
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.SessionId();
		session, _ = manager.provider.SessionInit(sid);
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid := url.QueryEscape(cookie.Value);
		session, _ = manager.provider.SessionRead(sid);
	}
	return;
}

func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName);
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.lock.Lock();
		defer manager.lock.Unlock();
		manager.provider.SessionDestroy(cookie.Value)
		expiration := time.Now();

		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}


func (manager *Manager) GC() {
	manager.lock.Lock();
	defer manager.lock.Unlock();
	manager.provider.SessionGC(manager.maxLifeTime);
	time.AfterFunc(time.Duration(manager.maxLifeTime), func () { manager.GC() })
}


type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy (sid string) error
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{})error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

var pdr = make(map[string]Provider)

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, ok := pdr[name]; ok {
		panic("session: Resigter called twice for provider " + name)
	}

	pdr[name] = provider;
}