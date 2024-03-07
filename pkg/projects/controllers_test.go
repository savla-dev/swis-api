package projects

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.savla.dev/swis/v5/pkg/core"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// https://circleci.com/blog/gin-gonic-testing/
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	projectsRouter := router.Group("/projects")
	Routes(projectsRouter)
	return router
}

func setupCache() {
	if Cache == nil {
		Cache = &core.Cache{}
	}
}

/*
 * POST /projects
 * GET  /projects
 * GET  /projects/:id
 */

func TestPostNewProjectByKey(t *testing.T) {
	setupCache()
	r := setupRouter()

	var project Project = Project{
		ID:          "test_project",
		Name:        "Test Project",
		Description: "Description for a test project",
		DocsLink:    "https://savla.dev",
		Manager:     "random",
		Published:   false,
		Backuped:    true,
		URL:         "http://savla.dev",
	}

	jsonValue, _ := json.Marshal(project)
	req, _ := http.NewRequest("POST", "/projects/test_project", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetProjects(t *testing.T) {
	setupCache()
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/projects/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	//var projects = make(map[string]Project)
	var projects = struct {
		Projects map[string]Project `json:"items"`
	}{}
	json.Unmarshal(w.Body.Bytes(), &projects)

	// (krusty Oct 22, 2022) this is not right: cannot get any POST'd project, even when there is time.Sleep()
	// function used, even if the order is POST then GET, both functions even points to the same memory address,
	// garbage collector (GC) does nothing between the two requests, so maybe it is the testing package
	// who cleans the memory after each request (does not even work when the two Test functions are merged)...
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, projects.Projects)
}

func TestGetProjectByKey(t *testing.T) {
	setupCache()
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/projects/test_project", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var project = struct {
		Project Project `json:"item"`
	}{}
	json.Unmarshal(w.Body.Bytes(), &project)

	// this should return StatusOK and requested project by ID
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, project.Project)
}

func TestUpdateProjectByKey(t *testing.T) {
	setupCache()
	r := setupRouter()

	var project Project = Project{
		ID:          "test_project",
		Name:        "Test Project",
		Description: "Description for a test project",
		DocsLink:    "https://savla.dev",
		Manager:     "genuine person",
		Published:   true,
		Backuped:    true,
		URL:         "http://savla.dev",
	}

	jsonValue, _ := json.Marshal(project)
	req, _ := http.NewRequest("PUT", "/projects/test_project", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var item = struct {
		Project Project `json:"item"`
	}{}
	json.Unmarshal(w.Body.Bytes(), &item)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, item.Project.Published, true)
}

func TestDeleteProjectByKey(t *testing.T) {
	setupCache()
	r := setupRouter()

	req, _ := http.NewRequest("DELETE", "/projects/test_project", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var ret = struct {
		Key string `json:"key"`
	}{}
	json.Unmarshal(w.Body.Bytes(), &ret)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, ret.Key, "test_project")
}

