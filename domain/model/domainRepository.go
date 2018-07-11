package model

var DomainRegistry *domainRegistry

type domainRegistry struct {
	Repository        Repository
	QueryService      QueryService
}
