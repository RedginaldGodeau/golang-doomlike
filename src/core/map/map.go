package _map

import (
	"bufio"
	g_type "github.com/RedginaldGodeau/GoDoom/src/core/type"
	"github.com/RedginaldGodeau/GoDoom/src/objects"
	"os"
)

func MappingByFile(filePath string) ([]objects.ObjectInterface, g_type.Vector2, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, g_type.NewVector2(0, 0), err
	}

	var y int32 = 0
	var xsize int32 = 0
	var mapping []objects.ObjectInterface
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for x, letter := range scanner.Text() {
			xsize = int32(x)
			if letter == 'X' {
				mapping = append(mapping, NewWall(g_type.NewVector2(xsize, y)))
			} else if letter == '0' {
				mapping = append(mapping, NewAir(g_type.NewVector2(xsize, y)))
			}
		}
		y++
	}
	return mapping, g_type.NewVector2(xsize, y), nil
}
