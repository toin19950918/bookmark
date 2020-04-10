package uuidutil

import uuid "github.com/iris-contrib/go.uuid"

func NewUUID() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}
