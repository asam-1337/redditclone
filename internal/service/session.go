package service

import (
	"encoding/json"
	"fmt"
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
)

const sessKeyLen = 10

type SessionManagerInterface interface {
	Create(in *entity.Session) (*entity.SessionID, error)
	Check(in *entity.SessionID) (*entity.Session, error)
	Delete(in *entity.SessionID) error
}

type SessionManager struct {
	redisConn redis.Conn
}

func NewSessionManager(conn redis.Conn) *SessionManager {
	return &SessionManager{
		redisConn: conn,
	}
}

func (s *SessionManager) Create(in *entity.Session) (*entity.SessionID, error) {
	id := entity.SessionID{ID: getRandomSeed(sessKeyLen)}
	dataSerialized, _ := json.Marshal(in)
	mkey := "sessions:" + id.ID
	result, err := redis.String(s.redisConn.Do("SET", mkey, dataSerialized, "EX", 86400))
	if err != nil {
		return nil, err
	}
	if result != "OK" {
		return nil, fmt.Errorf("result is not OK")
	}

	return &id, nil
}

func (s *SessionManager) Check(in *entity.SessionID) (*entity.Session, error) {
	mkey := "session:" + in.ID
	data, err := redis.Bytes((s.redisConn.Do("GET", mkey)))
	if err != nil {
		logrus.Errorf("cant get data")
		return nil, err
	}
	sess := &entity.Session{}
	err = json.Unmarshal(data, sess)
	if err != nil {
		logrus.Errorf("cant unpack session data")
		return nil, err
	}
	return sess, nil
}

func (s *SessionManager) Delete(in *entity.SessionID) error {
	mkey := "session:" + in.ID
	_, err := s.redisConn.Do("DEL", mkey)
	return err
}
