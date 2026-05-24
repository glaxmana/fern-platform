// Package graphql provides GraphQL API for the fern-reporter service
package graphql

import (
	"time"

	analyticsApp "github.com/guidewire-oss/fern-platform/internal/domains/analytics/application"
	"github.com/guidewire-oss/fern-platform/internal/domains/integrations"
	projectsApp "github.com/guidewire-oss/fern-platform/internal/domains/projects/application"
	tagsApp "github.com/guidewire-oss/fern-platform/internal/domains/tags/application"
	testingApp "github.com/guidewire-oss/fern-platform/internal/domains/testing/application"
	"github.com/guidewire-oss/fern-platform/pkg/logging"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the root GraphQL resolver.
//
// Dataloaders are intentionally NOT stored on the resolver — they are
// constructed per HTTP request in graphqlHandler so their per-key cache
// is request-scoped. A previous version kept a process-wide Loaders here,
// which meant project stats returned at server startup (potentially with
// missing test_runs while a seed was still in flight) were cached forever
// and the UI showed stale zeros.
type Resolver struct {
	testingService        *testingApp.TestRunService
	projectService        *projectsApp.ProjectService
	tagService            *tagsApp.TagService
	flakyDetectionService *analyticsApp.FlakyDetectionService
	jiraConnectionService *integrations.JiraConnectionService
	db                    *gorm.DB
	logger                *logging.Logger
	treemap               TreemapCache
}

// SetTreemapCache swaps in a custom backing cache (e.g. Redis-backed
// for multi-replica deploys). Call once after NewResolver. If unset,
// the resolver falls back to the in-process map created by NewResolver.
func (r *Resolver) SetTreemapCache(cache TreemapCache) {
	r.treemap = cache
}

// NewResolver creates a new GraphQL resolver
func NewResolver(
	testingService *testingApp.TestRunService,
	projectService *projectsApp.ProjectService,
	tagService *tagsApp.TagService,
	flakyDetectionService *analyticsApp.FlakyDetectionService,
	jiraConnectionService *integrations.JiraConnectionService,
	db *gorm.DB,
	logger *logging.Logger,
) *Resolver {
	return &Resolver{
		testingService:        testingService,
		projectService:        projectService,
		tagService:            tagService,
		flakyDetectionService: flakyDetectionService,
		jiraConnectionService: jiraConnectionService,
		db:                    db,
		logger:                logger,
		treemap:               newInMemoryTreemapCache(60 * time.Second),
	}
}
