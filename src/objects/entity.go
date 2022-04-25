package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
    HERO_ENTITY = iota
)


type Entity struct {
	Image    *ebiten.Image   // Tile image
	n         int            // Entity number of usage on single level
    i         int            // Entity index
    max       int            // maxium number of entities on single level (validation cheks)
    min       int            // minimum number of entities on single level (validation cheks)
    xs        []int          // positions of entities
    ys        []int          // position of entities
}

type EntityStack struct {
	Stack       []Entity     // collection of Entities
    current     int
}

// creates new Entity
func NewEntity(path string, i int, min int, max int) Entity {
	var e Entity

    e.Image  = loadImage(path)
	e.min    = min
	e.max    = max
    e.xs = append(e.xs, -1)
    e.ys = append(e.ys, -1)

	return e
}

func (es *EntityStack) AppentToEntityStack(e Entity) {
    es.Stack = append(es.Stack, e)
}

func (e Entity) EntityCanBeDrawn() (bool) {
    if e.n +1 < e.max {
        return true
    } else {
        return false
    }
}

func CreateNewEntityStack() EntityStack {
    var es EntityStack

    hero_entity := NewEntity("assets/hero_entity_icon.png", HERO_ENTITY, 1, 1)
    es.AppentToEntityStack(hero_entity)

    return es
}

