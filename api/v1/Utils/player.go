package utils

import (
	pb "github.com/LiamCWest/ChatTest/api/v1"
)

type Player struct {
	id         string
	name       string
	pos        Vector2
	radius     float32
	gameObject *GameObject
}

func NewPlayer(id string, name string, pos Vector2, radius float32) *Player {
	return &Player{id: id, name: name, pos: pos, radius: radius}
}

func NewPlayerFromMessage(p *pb.Player) *Player {
	return &Player{id: p.Id.Id, name: p.Name, pos: NewVector2(p.X, p.Y), radius: p.Radius}
}

func (p *Player) GenGameObject() *GameObject {
	p.gameObject = NewGameObject(p.name, p.pos, [][4][3]float32{[4][3]float32{
		[3]float32{0, 50, 0},
		[3]float32{0, 0, 0},
		[3]float32{50, 0, 0},
		[3]float32{50, 50, 0},
	}})
	return p.gameObject
}

func (p *Player) GetGameObject() *GameObject {
	return p.gameObject
}

func (p *Player) ToMessage() *pb.Player {
	return &pb.Player{Id: &pb.PlayerID{Id: p.id}, Name: p.name, X: p.pos.X, Y: p.pos.Y, Radius: p.radius}
}

func (p *Player) GetID() string {
	return p.id
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetPos() Vector2 {
	return p.pos
}

func (p *Player) GetRadius() float32 {
	return p.radius
}

func (p *Player) SetPos(pos Vector2) {
	p.pos = pos
}

func (p *Player) SetRadius(radius float32) {
	p.radius = radius
}

func (p *Player) Move(pos Vector2) {
	p.pos = p.pos.Add(pos)
}
