package database

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestBaseModel(t *testing.T) {
	testDBId := "TestBaseModel.db"
	db, err := Connect(testDBId)
	defer os.Remove(testDBId)
	require.NoError(t, err)

	err = db.AutoMigrate(BaseModel{})
	require.NoError(t, err)

	startTime := time.Now()
	model := BaseModel{}
	db.Create(&model)
	require.Greater(t, model.ID, uint(0))
	require.Equal(t, 36, len(model.UUID.String()))
	require.GreaterOrEqual(t, model.CreatedAt, startTime)
	require.GreaterOrEqual(t, model.UpdatedAt, startTime)
	require.True(t, model.DeletedAt.Time.IsZero())

	var models []BaseModel
	db.Where("UUID = ?", model.UUID).Find(&models)
	require.NotEmpty(t, models)
	require.Equal(t, model.ID, models[0].ID)

	db.Delete(&model)
	db.Where("UUID = ?", model.UUID).Find(&models)
	require.Empty(t, models)

	db.Unscoped().Where("UUID = ?", model.UUID).Find(&models)
	require.NotEmpty(t, models)
	require.GreaterOrEqual(t, models[0].DeletedAt.Time, startTime)
}
