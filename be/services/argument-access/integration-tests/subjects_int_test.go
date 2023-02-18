package integration_tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsertSubjectSuccessfully(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbCleanUp(ctx)

	res, err := argumentAccessClient.CreateSubject(ctx, subject)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotNil(t, res.Subject)
	assert.Greater(t, len(res.Subject.Id), 1)
}

func TestInsertWithId(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbCleanUp(ctx)

	res, err := argumentAccessClient.CreateSubject(ctx, subject2)
	assert.NotNil(t, err)
	assert.Nil(t, res)
}
