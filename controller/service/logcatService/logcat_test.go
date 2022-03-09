package logcatService

import (
	"sirclo/project-capstone/entities/logcatEntities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLogcat(t *testing.T) {
	logcatServiceMock := NewLogcatService(mockLogcatRepository{})
	logcat, err := logcatServiceMock.CreateLogcat("logcat-123456789", "logcat-message", "user")
	assert.Nil(t, err)
	assert.Equal(t, "logcat-123456789", logcat.ID, "tidak sama")
	assert.Equal(t, "message-logcat-123456789", logcat.Message, "tidak sama")
}

func TestGetLogcats(t *testing.T) {
	logcatServiceMock := NewLogcatService(mockLogcatRepository{})
	logcat, err := logcatServiceMock.GetLogcats()

	expected, _ := []logcatEntities.Logcat{
		{
			ID:      "logcat-123456789",
			Message: "message-logcat-123456789",
		},
		{
			ID:      "logcat-987654321",
			Message: "message-logcat-987654321",
		},
	}, err

	assert.Nil(t, err)
	assert.Equal(t, expected, logcat)
}

func TestGetLogcatUser(t *testing.T) {
	logcatServiceMock := NewLogcatService(mockLogcatRepository{})
	logcat, err := logcatServiceMock.GetLogcatUser("user-for-testing")

	expected, _ := []logcatEntities.Logcat{
		{
			ID:      "logcat-123456789",
			Message: "message-logcat-123456789",
		},
		{
			ID:      "logcat-987654321",
			Message: "message-logcat-987654321",
		},
	}, err

	assert.Nil(t, err)
	assert.Equal(t, expected, logcat)
}

type mockLogcatRepository struct{}

func (mock mockLogcatRepository) CreateLogcat(lc logcatEntities.Logcat) (logcatEntities.Logcat, error) {
	return logcatEntities.Logcat{
		ID:      "logcat-123456789",
		Message: "message-logcat-123456789",
	}, nil
}

func (mock mockLogcatRepository) GetLogcats() ([]logcatEntities.Logcat, error) {
	return []logcatEntities.Logcat{
		{
			ID:      "logcat-123456789",
			Message: "message-logcat-123456789",
		},
		{
			ID:      "logcat-987654321",
			Message: "message-logcat-987654321",
		},
	}, nil
}

func (mock mockLogcatRepository) GetLogcatUser(userID string) ([]logcatEntities.Logcat, error) {
	return []logcatEntities.Logcat{
		{
			ID:      "logcat-123456789",
			Message: "message-logcat-123456789",
		},
		{
			ID:      "logcat-987654321",
			Message: "message-logcat-987654321",
		},
	}, nil
}
