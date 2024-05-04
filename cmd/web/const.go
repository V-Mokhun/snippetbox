package main

type contextKey string

const SessionFlashKey = "flash"
const SessionUserIdKey = "authenticatedUserID"
const SessionRedirectUrlKey = "previousUrl"

const isAuthenticatedContextKey = contextKey("isAuthenticated")
