package game

import (
	"cmp"
	_map "github.com/RedginaldGodeau/GoDoom/src/core/map"
	_type "github.com/RedginaldGodeau/GoDoom/src/core/type"
	objects2 "github.com/RedginaldGodeau/GoDoom/src/objects"
	"slices"
)

type SceneInterface interface{}

type Scene struct {
	levelSize _type.Vector2
	objects   []objects2.ObjectInterface
}

func NewScene(levelPath string) *Scene {
	var s Scene
	objs, size, err := _map.MappingByFile(levelPath)
	if err != nil {
		return nil
	}

	s.AddObjects(objs...)
	s.levelSize = size

	return &s
}

func (s *Scene) AddObjects(obj ...objects2.ObjectInterface) {
	s.objects = append(s.objects, obj...)
	slices.SortStableFunc(s.objects, func(a, b objects2.ObjectInterface) int {
		return cmp.Compare(a.GetZIndex(), b.GetZIndex())
	})
}

func (s *Scene) Draw() {
	for _, object := range s.objects {
		object.Draw()
	}
}
