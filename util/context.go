package util

import (
	"context"
	"crypto/rand"
)

const SessionKey = 1

type Session struct {
	CID    string
	Logger Logger
}

func SessionCid(ctx context.Context) string {
	session, ok := ctx.Value(SessionKey).(*Session)

	// Handle if session middleware is not used
	if !ok {
		return ""
	}

	return session.CID
}

func SessionLogger(ctx context.Context, defaultLogger Logger) Logger {
	session, ok := ctx.Value(SessionKey).(*Session)

	// Handle if session middleware is not used
	if !ok {
		return defaultLogger
	}

	return session.Logger
}

func NewSessionCtx(cid string, log Logger) context.Context {
	session := Session{
		cid,
		log,
	}
	return context.WithValue(context.Background(), SessionKey, &session)
}

func NewId() []byte {
	ret := make([]byte, 20)
	if _, err := rand.Read(ret); err != nil {
		panic(err)
	}
	return ret
}
