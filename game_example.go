package peano

const GameComponentTotal = 5

const (
	Position = iota
	Direction
	Speed
	Health
	Sprite //next
)

type PositionComponent struct {
	X, Y float64
}

type DirectionComponent struct {
	X, Y float64
}

type SpeedComponent int

type HealthComponent int

type SpriteComponent struct {
	tag  string
	W, H int
}

func (e *Entity) AddPosition(x float64, y float64) {
	var c *PositionComponent
	if comp, ok := e.Create(Position); ok {
		c = comp.(*PositionComponent)
	} else {
		c = new(PositionComponent)
	}
	c.X = x
	c.Y = y

	e.components[Position] = c
	e.onComponentAdd.Execute(e, Position, c)
}

func (e *Entity) ReplacePosition(x float64, y float64) {
	if e.components[Position] != nil {
		c := e.components[Position].(*PositionComponent)
		c.X = x
		c.Y = y

		e.onComponentReplace.Execute(e, Position, c)
	} else {
		e.AddPosition(x, y)
	}
}

func (e *Entity) OnPosition() {
	e.onComponentAdd.Execute(e, Position, e.components[Position])
}

func (e *Entity) OffPosition() {
	e.onComponentOff.Execute(e, Position, e.components[Position])
}

func (e *Entity) RemovePosition() {
	e.onComponentRemove.Execute(e, Position, e.components[Position])
}

func (e *Entity) GetPosition() *PositionComponent {
	return e.components[Position].(*PositionComponent)
}

func (e *Entity) AddDirection(y float64, x float64) {
	var c *DirectionComponent
	if comp, ok := e.Create(Direction); ok {
		c = comp.(*DirectionComponent)
	} else {
		c = new(DirectionComponent)
	}
	c.Y = y
	c.X = x

	e.components[Direction] = c
	e.onComponentAdd.Execute(e, Direction, c)
}

func (e *Entity) ReplaceDirection(y float64, x float64) {
	if e.components[Direction] != nil {
		c := e.components[Direction].(*DirectionComponent)
		c.Y = y
		c.X = x

		e.onComponentReplace.Execute(e, Direction, c)
	} else {
		e.AddDirection(y, x)
	}
}

func (e *Entity) OnDirection() {
	e.onComponentAdd.Execute(e, Direction, e.components[Direction])
}

func (e *Entity) OffDirection() {
	e.onComponentOff.Execute(e, Direction, e.components[Direction])
}

func (e *Entity) RemoveDirection() {
	e.onComponentRemove.Execute(e, Direction, e.components[Direction])
}

func (e *Entity) GetDirection() *DirectionComponent {
	return e.components[Direction].(*DirectionComponent)
}

func (e *Entity) AddSpeed(speed int) {
	var c *SpeedComponent
	if comp, ok := e.Create(Speed); ok {
		c = comp.(*SpeedComponent)
	} else {
		c = new(SpeedComponent)
	}
	c = (*SpeedComponent)(&speed)
	e.components[Speed] = c
	e.onComponentAdd.Execute(e, Speed, c)
}

func (e *Entity) ReplaceSpeed(speed int) {
	if e.components[Speed] != nil {
		c := e.components[Speed].(*SpeedComponent)
		c = (*SpeedComponent)(&speed)
		e.onComponentReplace.Execute(e, Speed, c)
	} else {
		e.AddSpeed(speed)
	}
}

func (e *Entity) OnSpeed() {
	e.onComponentAdd.Execute(e, Speed, e.components[Speed])
}

func (e *Entity) OffSpeed() {
	e.onComponentOff.Execute(e, Speed, e.components[Speed])
}

func (e *Entity) RemoveSpeed() {
	e.onComponentRemove.Execute(e, Speed, e.components[Speed])
}

func (e *Entity) GetSpeed() int {
	return e.components[Speed].(int)
}

func (e *Entity) AddHealth(health int) {
	var c *HealthComponent
	if comp, ok := e.Create(Health); ok {
		c = comp.(*HealthComponent)
	} else {
		c = new(HealthComponent)
	}
	c = (*HealthComponent)(&health)
	e.components[Health] = c
	e.onComponentAdd.Execute(e, Health, c)
}

func (e *Entity) ReplaceHealth(health int) {
	if e.components[Health] != nil {
		c := e.components[Health].(*HealthComponent)
		c = (*HealthComponent)(&health)
		e.onComponentReplace.Execute(e, Health, c)
	} else {
		e.AddHealth(health)
	}
}

func (e *Entity) OnHealth() {
	e.onComponentAdd.Execute(e, Health, e.components[Health])
}

func (e *Entity) OffHealth() {
	e.onComponentOff.Execute(e, Health, e.components[Health])
}

func (e *Entity) RemoveHealth() {
	e.onComponentRemove.Execute(e, Health, e.components[Health])
}

func (e *Entity) GetHealth() int {
	return e.components[Health].(int)
}

func (e *Entity) AddSprite(tag string, w int, h int) {
	var c *SpriteComponent
	if comp, ok := e.Create(Sprite); ok {
		c = comp.(*SpriteComponent)
	} else {
		c = new(SpriteComponent)
	}
	c.tag = tag
	c.W = w
	c.H = h

	e.components[Sprite] = c
	e.onComponentAdd.Execute(e, Sprite, c)
}

func (e *Entity) ReplaceSprite(tag string, w int, h int) {
	if e.components[Sprite] != nil {
		c := e.components[Sprite].(*SpriteComponent)
		c.tag = tag
		c.W = w
		c.H = h

		e.onComponentReplace.Execute(e, Sprite, c)
	} else {
		e.AddSprite(tag, w, h)
	}
}

func (e *Entity) OnSprite() {
	e.onComponentAdd.Execute(e, Sprite, e.components[Sprite])
}

func (e *Entity) OffSprite() {
	e.onComponentOff.Execute(e, Sprite, e.components[Sprite])
}

func (e *Entity) RemoveSprite() {
	e.onComponentRemove.Execute(e, Sprite, e.components[Sprite])
}

func (e *Entity) GetSprite() *SpriteComponent {
	return e.components[Sprite].(*SpriteComponent)
}
