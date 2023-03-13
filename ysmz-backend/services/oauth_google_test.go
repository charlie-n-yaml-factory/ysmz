package services_test

import (
	"testing"

	"github.com/charlie-n-yaml-factory/ysmz/ysmz-backend/config"
	"github.com/charlie-n-yaml-factory/ysmz/ysmz-backend/mocks"
	"github.com/charlie-n-yaml-factory/ysmz/ysmz-backend/services"
	"github.com/golang/mock/gomock"
)

func TestGetOAuthGoogleState(t *testing.T) {
	t.Run("should return a string of length 64", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		mockConfig := mocks.NewMockConfigInterface(ctrl)
		mockConfig.EXPECT().Config().Return(&config.ConfigStruct{}).AnyTimes()

		mockRedisDB := mocks.NewMockRedisInterface(ctrl)
		mockRedisDB.EXPECT().Get(gomock.Any()).Return("", nil).AnyTimes()
		mockRedisDB.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

		state, err := services.SaveOAuthGoogleState(mockRedisDB)
		if err != nil {
			t.Errorf("GetOAuthGoogleState() returned an error: %v", err)
		}
		if len(state) != 64 {
			t.Errorf("GetOAuthGoogleState() returned a string of length %d, expected 64", len(state))
		}
	})
}
