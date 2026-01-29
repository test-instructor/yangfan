package automation

import (
	"context"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"github.com/test-instructor/yangfan/server/v2/global"
	automationModel "github.com/test-instructor/yangfan/server/v2/model/automation"
	automationReq "github.com/test-instructor/yangfan/server/v2/model/automation/request"
	"gorm.io/gorm"
)

func TestAddAutoCaseStepApi_ReturnsErrorWhenRequestMissing(t *testing.T) {
	db := setupTestDB(t)
	restore := swapGlobalDB(db)
	defer restore()

	err := db.AutoMigrate(
		&automationModel.Request{},
		&automationModel.AutoStep{},
		&automationModel.AutoCaseStep{},
		&automationModel.AutoCaseStepRelation{},
	)
	require.NoError(t, err)

	caseStep := automationModel.AutoCaseStep{
		StepConfig: automationModel.StepConfig{StepName: "case-step"},
	}
	require.NoError(t, db.Create(&caseStep).Error)

	srcStep := automationModel.AutoStep{
		StepConfig: automationModel.StepConfig{StepName: "src-step"},
		StepType:   1,
	}
	require.NoError(t, db.Create(&srcStep).Error)

	svc := &AutoCaseStepService{}
	_, err = svc.AddAutoCaseStepApi(context.Background(), &automationReq.AutoCaseStepSearchApi{
		ID:    caseStep.ID,
		ApiID: srcStep.ID,
	})
	require.Error(t, err)
}

func TestAddAutoCaseStepApi_ClonesStepAndRequestAndCreatesRelation(t *testing.T) {
	db := setupTestDB(t)
	restore := swapGlobalDB(db)
	defer restore()

	err := db.AutoMigrate(
		&automationModel.Request{},
		&automationModel.AutoStep{},
		&automationModel.AutoCaseStep{},
		&automationModel.AutoCaseStepRelation{},
	)
	require.NoError(t, err)

	caseStep := automationModel.AutoCaseStep{
		StepConfig: automationModel.StepConfig{StepName: "case-step"},
	}
	require.NoError(t, db.Create(&caseStep).Error)

	req := automationModel.Request{
		Method:    "GET",
		URL:       "https://example.com",
		ProjectId: 1,
	}
	require.NoError(t, db.Create(&req).Error)

	srcStep := automationModel.AutoStep{
		StepConfig: automationModel.StepConfig{StepName: "src-step"},
		StepType:   1,
		RequestID:  req.ID,
		ProjectId:  1,
	}
	require.NoError(t, db.Create(&srcStep).Error)

	svc := &AutoCaseStepService{}
	data, err := svc.AddAutoCaseStepApi(context.Background(), &automationReq.AutoCaseStepSearchApi{
		ID:    caseStep.ID,
		ApiID: srcStep.ID,
	})
	require.NoError(t, err)
	require.NotNil(t, data)

	newIDAny, ok := data["id"]
	require.True(t, ok)
	newID, ok := newIDAny.(uint)
	require.True(t, ok)
	require.NotZero(t, newID)
	require.NotEqual(t, srcStep.ID, newID)

	var newStep automationModel.AutoStep
	require.NoError(t, db.First(&newStep, "id = ?", newID).Error)
	require.Equal(t, 11, newStep.StepType)
	require.NotZero(t, newStep.RequestID)
	require.NotEqual(t, srcStep.RequestID, newStep.RequestID)

	var newReq automationModel.Request
	require.NoError(t, db.First(&newReq, "id = ?", newStep.RequestID).Error)
	require.Equal(t, req.Method, newReq.Method)
	require.Equal(t, req.URL, newReq.URL)

	var rel automationModel.AutoCaseStepRelation
	require.NoError(t, db.First(&rel, "auto_case_step_id = ? AND auto_step_id = ?", caseStep.ID, newID).Error)
	require.Equal(t, uint(999), rel.Sort)
}

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	require.NoError(t, err)
	return db
}

func swapGlobalDB(db *gorm.DB) func() {
	old := global.GVA_DB
	global.GVA_DB = db
	return func() {
		global.GVA_DB = old
	}
}
