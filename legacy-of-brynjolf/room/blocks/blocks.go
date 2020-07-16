package blocks

import (
	"fmt"
	"legacy-of-brynjolf/command"
	"legacy-of-brynjolf/position"
	"legacy-of-brynjolf/room/entities"
)

type Block struct {
	entity entities.Entity
	pos position.Position
}

func NewBlock(entity entities.Entity, r int, c int) Block {
	pos := position.NewPostion(r, c)
	return Block{
		entity: entity,
		pos: pos,
	}
}

func (b Block) Entity() entities.Entity {
	return b.entity
}

func (b Block) Pos() position.Position {
	return b.pos
}

func (b *Block) UpdateEntity(e entities.Entity) {
	b.entity = e
}

func (b *Block) UpdatePos(c command.Command) {
	b.pos = b.pos.Update(c)
}

func (b Block) SamePosition(otherBlock Block) bool {
	return b.Pos().Row() == otherBlock.Pos().Row() && b.Pos().Col() == otherBlock.Pos().Col()
}

func BuildRow(row []string, rowIndex int) ([]Block, error) {
	var blockRow []Block
	for colIndex, entity := range row {
		entity, err := entities.BuildEntity(entity)
		if err != nil {
			return nil, err
		}
		block := NewBlock(entity, rowIndex, colIndex)
		blockRow = append(blockRow, block)
	}
	return blockRow, nil
}

func (b Block) IsGuard() bool {
	return b.entity == entities.Guard
}

func (b Block) IsBrynjolf() bool {
	return b.entity == entities.Brynjolf
}

func (b Block) IsExit() bool {
	return b.entity == entities.Exit
}

func (b Block) Includes(entities []entities.Entity) bool {
	for _, e := range entities{
		if e == b.entity{
			return true
		}
	}
	return false
}

func (b Block) Display() {
	fmt.Print(b.entity + ",")
}
