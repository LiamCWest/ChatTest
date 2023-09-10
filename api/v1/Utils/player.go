package utils

import (
	pb "github.com/LiamCWest/ChatTest/api/v1"
)

type player struct {
	id     string
	name   string
	pos    Vector2
	radius float32
}

func NewPlayer(id string, name string, pos Vector2, radius float32) player {
	return player{id: id, name: name, pos: pos, radius: radius}
}

func NewPlayerFromMessage(p *pb.Player) player {
	return player{id: p.Id.Id, name: p.Name, pos: NewVector2(p.X, p.Y), radius: p.Radius}
}

func (p player) ToMessage() *pb.Player {
	return &pb.Player{Id: &pb.PlayerID{Id: p.id}, Name: p.name, X: p.pos.X, Y: p.pos.Y, Radius: p.radius}
}

func (p player) GetID() string {
	return p.id
}

func (p player) GetName() string {
	return p.name
}

func (p player) GetPos() Vector2 {
	return p.pos
}

func (p player) GetRadius() float32 {
	return p.radius
}

func (p *player) SetPos(pos Vector2) {
	p.pos = pos
}

func (p *player) SetRadius(radius float32) {
	p.radius = radius
}
