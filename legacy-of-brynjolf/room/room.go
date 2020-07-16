package room

import (
	"fmt"
	"legacy-of-brynjolf/command"
	_blocks "legacy-of-brynjolf/room/blocks"
	_entities "legacy-of-brynjolf/room/entities"
	"strings"
)

type Room struct {
	blocks [][]_blocks.Block
}


func NewRoom(data string) (Room, error) {
	var blocks [][]_blocks.Block
	rows := strings.Split(strings.TrimSpace(data), "\n")
	for index, row := range rows {
		blocksRow, err := _blocks.BuildRow(strings.Split(strings.TrimSpace(row), ","), index)
		if err != nil{
			return Room{}, err
		}
		blocks = append(blocks, blocksRow)
	}
	return Room{blocks: blocks}, nil
}

func (r Room) Size() int {
	return len(r.blocks)
}

func (r Room) duplicateRoomBlocks() [][]_blocks.Block {
	duplicate := make([][]_blocks.Block, len(r.blocks))
	for index := range r.blocks {
		duplicate[index] = make([]_blocks.Block, len(r.blocks[index]))
		copy(duplicate[index], r.blocks[index])
	}
	return duplicate
}

func (r Room) MoveEntities(movingBlocks []_blocks.Block, command command.Command) Room {
	newBlocks := r.duplicateRoomBlocks()
	for index, block := range movingBlocks {
		blockingEntities := block.Entity().GetBlockingEntities()
		newBlocks[block.Pos().Row()][block.Pos().Col()].UpdateEntity(_entities.EmptySpace)
		r.moveEntity(&block, blockingEntities, newBlocks, command)
		r.updateNewBlocks(block, newBlocks)
		movingBlocks[index] = block
	}
	return Room{newBlocks}
}

func (r Room) updateNewBlocks(block _blocks.Block, newBlocks [][]_blocks.Block) {
	oldBlock := r.blocks[block.Pos().Row()][block.Pos().Col()]
	newBlock := &newBlocks[block.Pos().Row()][block.Pos().Col()]
	if !oldBlock.IsExit() && !(block.IsBrynjolf() && newBlock.IsGuard()) {
		newBlock.UpdateEntity(block.Entity())
	}
}

func (r Room) moveEntity(block *_blocks.Block, blockingEntities []_entities.Entity, newBlocks [][]_blocks.Block, command command.Command) {
	for r.NotAtEdgeOrBlocked(*block, blockingEntities, command) {
		if block.IsBrynjolf() && (newBlocks[block.Pos().Row()][block.Pos().Col()].IsGuard() || newBlocks[block.Pos().Row()][block.Pos().Col()].IsExit()) {
			break
		}
		block.UpdatePos(command)
	}
}

func (r Room) notBlocked(row int, col int, blockingEntities []_entities.Entity) bool {
	return !r.blocks[row][col].Includes(blockingEntities)
}

func (r Room) NotAtEdgeOrBlocked(b _blocks.Block, blockingEntities []_entities.Entity, c command.Command) bool {
	switch c {
	case command.Up:
		return b.Pos().Row() > 0 && r.notBlocked(b.Pos().Row() - 1, b.Pos().Col(), blockingEntities)
	case command.Down:
		return b.Pos().Row() < len(r.blocks) - 1 && r.notBlocked(b.Pos().Row() + 1, b.Pos().Col(), blockingEntities)
	case command.Left:
		return b.Pos().Col() > 0 && r.notBlocked(b.Pos().Row(), b.Pos().Col() - 1, blockingEntities)
	case command.Right:
		return b.Pos().Col() < len(r.blocks[0]) - 1 && r.notBlocked(b.Pos().Row(), b.Pos().Col() + 1, blockingEntities)
	}
	return false
}

func (r Room) FindBlocks(entities []_entities.Entity) []_blocks.Block {
	var blocks []_blocks.Block
	for _, blocksRow := range r.blocks {
		for _, block := range blocksRow {
			if block.Includes(entities) {
				if block.IsGuard() {
					blocks = append([]_blocks.Block{block}, blocks...)
				}else {
					blocks = append(blocks, block)
				}
			}
		}
	}
	return blocks
}

func (r Room) Display(mssg string) {
	fmt.Println(mssg)
	for _, row := range r.blocks {
		for _, block := range row {
			block.Display()
		}
		fmt.Print("\n")
	}
}