package day6

import (
	"bufio"
	"embed"
	strEx "ibratoev/aoc22/utils/string"
	"strings"
)

//go:embed input.txt
var f embed.FS

type Node interface {
	Name() string
	Size() int
}

type File struct {
	name string
	size int
}

func (file *File) Name() string {
	return file.name
}

func (file *File) Size() int {
	return file.size
}

type Directory struct {
	name     string
	children []Node
	Parent   *Directory
}

func (dir *Directory) Name() string {
	return dir.name
}

func (dir *Directory) Size() int {
	sum := 0
	for _, child := range dir.children {
		sum += child.Size()
	}
	return sum
}

func (dir *Directory) GetChildDirByName(name string) *Directory {
	for _, c := range dir.children {
		cd, ok := c.(*Directory)
		if ok && cd.Name() == name {
			return cd
		}
	}
	panic("Not expected to search for non existing directory")
}

func (dir *Directory) Dirs() []*Directory {
	result := make([]*Directory, 0, 8)
	for _, c := range dir.children {
		d, ok := c.(*Directory)
		if ok {
			result = append(result, d)
		}
	}
	return result
}

func (dir *Directory) ParseChildren(lines []string) {
	dir.children = make([]Node, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		if parts[0] == "dir" {
			dir.children[i] = &Directory{name: parts[1], Parent: dir}
		} else {
			size := strEx.Atoi(parts[0])
			dir.children[i] = &File{name: parts[1], size: size}
		}
	}
}

func Day7Part1() int {
	file, _ := f.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	root := &Directory{name: "."}
	root.Parent = root
	currentDir := root

	scanner.Scan()

loop:
	for {
		cmdParts := strings.Split(scanner.Text(), " ")
		switch cmdParts[1] {
		case "cd":
			dir := cmdParts[2]
			if dir == "/" {
				currentDir = root
			} else if dir == ".." {
				currentDir = currentDir.Parent
			} else {
				currentDir = currentDir.GetChildDirByName(dir)
			}
			if !scanner.Scan() {
				break loop
			}
		case "ls":
			if currentDir.children != nil {
				panic("Not expected to ls twice")
			}
			lines := make([]string, 0, 8)
			hasMore := scanner.Scan()
			for {
				if !hasMore || scanner.Text()[0] == '$' {
					currentDir.ParseChildren(lines)
					break
				} else {
					lines = append(lines, scanner.Text())
					hasMore = scanner.Scan()
				}
			}
			if !hasMore {
				break loop
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sum := 0
	allDirs := make([]*Directory, 0, 64)
	allDirs = append(allDirs, root)
	for i := 0; i < len(allDirs); i++ {
		allDirs = append(allDirs, allDirs[i].Dirs()...)
		if allDirs[i].Size() <= 100000 {
			sum += allDirs[i].Size()
		}
	}

	return sum
}

func Day7Part2() int {
	file, _ := f.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	root := &Directory{name: "."}
	root.Parent = root
	currentDir := root

	scanner.Scan()

loop:
	for {
		cmdParts := strings.Split(scanner.Text(), " ")
		switch cmdParts[1] {
		case "cd":
			dir := cmdParts[2]
			if dir == "/" {
				currentDir = root
			} else if dir == ".." {
				currentDir = currentDir.Parent
			} else {
				currentDir = currentDir.GetChildDirByName(dir)
			}
			if !scanner.Scan() {
				break loop
			}
		case "ls":
			if currentDir.children != nil {
				panic("Not expected to ls twice")
			}
			lines := make([]string, 0, 8)
			hasMore := scanner.Scan()
			for {
				if !hasMore || scanner.Text()[0] == '$' {
					currentDir.ParseChildren(lines)
					break
				} else {
					lines = append(lines, scanner.Text())
					hasMore = scanner.Scan()
				}
			}
			if !hasMore {
				break loop
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	rootSize := root.Size()
	neededSpace := 30000000 - (70000000 - rootSize)

	bestSize := rootSize

	allDirs := make([]*Directory, 0, 64)
	allDirs = append(allDirs, root)
	for i := 0; i < len(allDirs); i++ {
		allDirs = append(allDirs, allDirs[i].Dirs()...)
		size := allDirs[i].Size()
		if size >= neededSpace && size < bestSize {
			bestSize = size
		}
	}

	return bestSize
}
