package repository

import (
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/user"
	"github.com/jinzhu/gorm"
)

type SessionGormRepo struct {
	conn *gorm.DB
}

// NewSessionGormRepo  returns a new SessionGormRepo object
func NewSessionGormRepo(db *gorm.DB) user.SessionRepository {
	return &SessionGormRepo{conn: db}
}

// Session returns a given stored session
func (sr *SessionGormRepo) Session(sessionId string) (*model.Session, []error) {
	session := model.Session{}
	errs := sr.conn.Find(&session, "session_id=?", sessionId).GetErrors()
	return &session, errs
}

// Returns all the sessions
func (sr *SessionGormRepo) Sessions() ([]model.Session, []error) {
	sessions := []model.Session{}
	errs := sr.conn.Find(&sessions).GetErrors()
	return sessions, errs
}

// StoreSession stores a given session
func (sr *SessionGormRepo) StoreSession(session *model.Session) (*model.Session, []error) {
	sess := session
	errs := sr.conn.Save(sess).GetErrors()
	return sess, errs
}

// DeleteSession deletes a given session
func (sr *SessionGormRepo) DeleteSession(sessionId string) (*model.Session, []error) {
	sess, errs := sr.Session(sessionId)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = sr.conn.Delete(sess, sessionId).GetErrors()
	return sess, errs
}
