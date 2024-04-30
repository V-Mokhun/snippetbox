package main

type contextKey string

const SessionFlashKey = "flash"
const SessionUserIdKey = "authenticatedUserID"

const isAuthenticatedContextKey = contextKey("isAuthenticated")
