package service

import (
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/user"
)

type SessionServiceImpl struct {
	sessionRepo user.SessionRepository
}

// NewSessionService  returns a new SessionService object
func NewSessionService(sessRepository user.SessionRepository) user.SessionService {
	return &SessionServiceImpl{sessionRepo: sessRepository}
}

// Session returns a given stored session
func (ss *SessionServiceImpl) Session(sessionId string) (*model.Session, []error) {
	return ss.sessionRepo.Session(sessionId)
}

// Returns all the sessions
func (ss *SessionServiceImpl) Sessions() ([]model.Session, []error) {
	return ss.sessionRepo.Sessions()
}

// StoreSession stores a given session
func (ss *SessionServiceImpl) StoreSession(session *model.Session) (*model.Session, []error) {
	return ss.sessionRepo.StoreSession(session)
}

// DeleteSession deletes a given session
func (ss *SessionServiceImpl) DeleteSession(sessionId string) (*model.Session, []error) {
	return ss.sessionRepo.DeleteSession(sessionId)
}
