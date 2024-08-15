package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/CongTran1999/todo-app/utils"
	"github.com/alexeyco/simpletable"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todo []item

func (t *Todo) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todo) Complete(index int) error {
	list := *t
	if index < 0 || index > len(list) {
		return errors.New("invalid index")
	}

	list[index-1].Done = true
	list[index-1].CompletedAt = time.Now()

	return nil
}

func (t *Todo) Delete(index int) error {
	list := *t
	if index < 0 || index > len(list) {
		return errors.New("invalid index")
	}

	*t = append(list[:index-1], list[index:]...)

	return nil
}

func (t *Todo) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return errors.New("file have no content")
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todo) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (t *Todo) Print() {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		task := utils.Blue(item.Task)
		done := utils.Blue("no")
		if item.Done {
			task = utils.Green(fmt.Sprintf("\u2705 %s", item.Task))
			done = utils.Green("yes")
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})

		table.Body = &simpletable.Body{Cells: cells}

		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: utils.Red(fmt.Sprintf("You have %d pending todos", t.CountPending()))},
		}}

		table.SetStyle(simpletable.StyleUnicode)

		table.Println()
	}
}

func (t *Todo) CountPending() int {
	cnt := 0
	for _, item := range *t {
		if !item.Done {
			cnt++
		}
	}

	return cnt
}
