package main

import (
	"bufio" 		// read from engine
	"os"			// read from engine
	"strings"		// split and replace function for strings
	"strconv"		// convert string to int
	"math/rand"		// play randomly
	"fmt"			// output
	"time"			// for rand.seed(stuff)
)

type BotStarter struct {
	name string		// bot name
	id, round int	// bot id and round #
	field Field
}

func (bs BotStarter) Init() BotStarter{ // field initialization
	bs.field = Field{}
	bs.field = bs.field.init()
	return bs
}
func (bs BotStarter) Parse() { // parse engine messages
	reader := bufio.NewReader(os.Stdin)
	for true {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		message := strings.Split(text, " ")
		if strings.Compare(message[0], "settings") == 0 {
			if strings.Compare(message[1], "your_bot") == 0 {
				bs.name = message[2]
			} else if strings.Compare(message[1], "your_botid") == 0 {
				bs.id, _ = strconv.Atoi(message[2])
			}
		} else if strings.Compare(message[0], "update") == 0 {
			if strings.Compare(message[2], "round") == 0 {
				bs.round, _ = strconv.Atoi(message[3])
			} else if strings.Compare(message[2], "field") == 0 {
				bs.field = bs.field.parse(message[3])
			}
		} else if strings.Compare(message[0], "action") == 0 {
			bs.Play()
		}
	}
}
func (bs BotStarter) Run(){
	bs.Parse()
}
func (bs BotStarter) Play(){
	rand.Seed(time.Now().UTC().UnixNano()) // not perfect but it works
	col_to_play := rand.Intn(bs.field.width)
	for bs.field.board[0][col_to_play] != 0 { // column has at least one free space
		col_to_play = rand.Intn(bs.field.width)
	}
	fmt.Println("place_disc", col_to_play)
}

type Field struct {
	height, width int
	board [][]int
}

func (f Field) init() Field{
	f.height = 6
	f.width = 7
	f.board = make([][]int, f.height)
	for row:= 0; row < f.height; row++{
		f.board[row] = make([]int, f.width)
		for col:= 0; col < f.width; col++{
			f.board[row][col] = 0
		}
	}
	return f
}
func (f Field) parse(s string) Field{ // parse engine field's update
	rows := strings.Split(s, ";")
	for row:=0; row < f.height; row++ {
		values := strings.Split(rows[row], ",")
		for col:=0; col < f.width; col++ {
			f.board[row][col], _ = strconv.Atoi(values[col])
		}
	}
	return f
}


func main() {
	starter := BotStarter{}
	starter = starter.Init()
	starter.Run()
}
/*
settings timebank 10000
settings time_per_move 500
settings player_names player1,player2
settings your_bot player1
settings your_botid 1
settings field_columns 7
settings field_rows 6
Round 1
update game round 1
update game field 0,0,0,0,0,0,0;0,0,0,0,0,0,0;0,0,0,0,0,0,0;0,0,0,0,0,0,0;0,0,0,0,0,0,0;0,0,0,0,0,0,0
action move 10000
 */