package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"

	"container/list"
)

// NPMURL is package url for npm registry
const NPMURL = "http://registry.npmjs.org/%s/latest"

var (
	errQueueIsEmpty         = errors.New("Queue is empty")
	errDependenciesNotFound = errors.New("Dependency not found")
)

// Queue ...
type Queue struct {
	*list.List
}

// Push value in queue
func (q *Queue) Push(value interface{}) {
	q.PushFront(value)
}

// Pull removes the last element from the Queue and returns it's value
func (q *Queue) Pull() (interface{}, error) {
	if q.Len() == 0 {
		return nil, errQueueIsEmpty
	}

	return q.Remove(q.Back()), nil
}

// NewQueue returns a queue
func NewQueue() *Queue { return &Queue{list.New()} }

// Package is default package information
type Package struct {
	Name string

	URL string
}

// String returns the name of the package
func (pk *Package) String() string { return pk.Name }

// Dependencies returns the list of direct dependencies for the package
func (pk *Package) Dependencies() ([]Package, error) {
	var (
		packages []Package
		data     map[string]interface{}
	)

	fmt.Println("Getting URL: ", pk.URL)
	res, err := http.Get(pk.URL)
	if err != nil {
		return nil, errors.Wrap(err, "fetching data from npm failed")
	}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling data failed")
	}

	dependencies, ok := data["dependencies"].(map[string]interface{})
	if !ok {
		return nil, errDependenciesNotFound
	}

	for dk := range dependencies {
		pkg := NewPackage(dk)
		packages = append(packages, pkg)
	}
	return packages, nil
}

// NewPackage Returns a new Package
func NewPackage(name string) Package {
	return Package{
		Name: name,
		URL:  fmt.Sprintf(NPMURL, name),
	}
}

func main() {

	var (
		queue        *Queue
		pkg          Package
		dependencies map[string]bool
	)

	queue = NewQueue()
	dependencies = make(map[string]bool)
	pkg = NewPackage("forever")

	queue.Push(pkg)

	for queue.Len() != 0 {
		val, err := queue.Pull()
		if err != nil && err == errQueueIsEmpty {
			break
		}

		pack := val.(Package)

		deps, err := pack.Dependencies()
		if err != nil && err != errDependenciesNotFound {
			fmt.Println("Error in Fetching Dependency, Ignoring Package: ", pack.String())
			continue
		} else if err != nil && err == errDependenciesNotFound {
			fmt.Println("Error Getting Dependency for: ", pack.String())
		}

		for _, dep := range deps {
			if _, ok := dependencies[dep.String()]; !ok {
				queue.Push(dep)
				dependencies[dep.String()] = true
			}
		}
	}

	for k := range dependencies {
		fmt.Println(k)
	}
}
