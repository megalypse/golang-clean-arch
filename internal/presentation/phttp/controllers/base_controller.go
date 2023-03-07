package controllers

import "github.com/megalypse/golang-clean-arch/internal/presentation/phttp"

type Controller interface {
	GetHandlers() map[string]phttp.RouteDefinition
}
