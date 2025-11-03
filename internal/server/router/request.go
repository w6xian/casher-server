package router

import "casher-server/internal/muxhttp"

func GetRequestIdKey() muxhttp.ContextKey {
	return muxhttp.ContextKey("request_id")
}

func GetLanguageKey() muxhttp.ContextKey {
	return muxhttp.ContextKey("language")
}
