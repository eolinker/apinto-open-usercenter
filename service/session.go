package service

import (
	"fmt"
	"github.com/eolinker/apinto-dashboard/cache"
	user_model "github.com/eolinker/apinto-open-usercenter/model"
	"time"
)

func sessionCacheKey(session string) string {
	return fmt.Sprintf("session:%s", session)
}

func newSessionCache() ISessionCache {
	return cache.CreateRedisCache[user_model.Session](time.Hour*24*7, sessionCacheKey)

}
