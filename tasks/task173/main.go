// echo | go run tasks/task173/main.go
package main

import (
	"errors"
	"fmt"
)

// label - уникальное наименование
type label string

// command - команда, которую можно выполнять в игре
type command label

// список доступных команд
var (
	eat  = command("eat")
	take = command("take")
	talk = command("talk to")
)

// thing - объект, который существует в игре
type thing struct {
	name    label
	actions map[command]string
}

// supports() возвращает true, если объект
// поддерживает команду action
func (t thing) supports(action command) bool {
	_, ok := t.actions[action]
	return ok
}

var (
	apple = thing{"apple", map[command]string{
		eat:  "mmm, delicious!",
		take: "you have an apple now",
	}}
	bob = thing{"bob", map[command]string{
		talk: "Bob says hello",
	}}
	coin = thing{"coin", map[command]string{
		take: "you have a coin now",
	}}
	mirror = thing{"mirror", map[command]string{
		take: "you have a mirror now",
		talk: "mirror does not answer",
	}}
	mushroom = thing{"mushroom", map[command]string{
		eat:  "tastes funny",
		take: "you have a mushroom now",
	}}
)

// step описывает шаг игры: сочетание команды и объекта
type step struct {
	cmd command
	obj thing
}

// isValid() возвращает true, если объект
// совместим с командой
func (s step) isValid() bool {
	return s.obj.supports(s.cmd)
}

// начало решения

// invalidStepError - ошибка, которая возникает,
// когда команда шага не совместима с объектом
type invalidStepError struct {
	st step
}

func (e invalidStepError) Error() string {
	return fmt.Sprintf("%s", e.st)
}

// notEnoughObjectsError - ошибка, которая возникает,
// когда в игре закончились объекты определенного типа
type notEnoughObjectsError struct {
	th thing
}

func (e notEnoughObjectsError) Error() string {
	return fmt.Sprintf("there are no %ss left", e.th)
}

// commandLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на выполнение команды
type commandLimitExceededError struct {
	nEaten   int
	nDialogs int
	obj      thing
	objLimit int
}

func (e commandLimitExceededError) Error() string {
	switch {
	case e.nEaten > 1:
		return "you don't want to eat anymore"
	case e.nDialogs > 1:
		return "you don't want to talk anymore"
	default:
		return fmt.Sprintf("you already have a %s", e.obj)
	}
}

// objectLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на количество объектов
// определенного типа в инвентаре
type objectLimitExceededError struct {
	obj   thing
	limit int
}

func (e objectLimitExceededError) Error() string {
	return fmt.Sprintf("there are no %ss left", e.obj)
}

// gameOverError - ошибка, которая произошла в игре
type gameOverError struct {
	// количество шагов, успешно выполненных
	// до того, как произошла ошибка
	nSteps int
	err    error
	// ...
}

func (e gameOverError) Error() string {
	return fmt.Sprintf("%s language error: %v", e.nSteps, e.err)
}

func (e gameOverError) Unwrap() error {
	return e.err
}

// player - игрок
type player struct {
	// количество съеденного
	nEaten int
	// количество диалогов
	nDialogs int
	// инвентарь
	inventory []thing
}

// has() возвращает true, если у игрока
// в инвентаре есть предмет obj
func (p *player) has(obj thing) bool {
	for _, got := range p.inventory {
		if got.name == obj.name {
			return true
		}
	}
	return false
}

// do() выполняет команду cmd над объектом obj
// от имени игрока
func (p *player) do(cmd command, obj thing) error {
	// действуем в соответствии с командой
	switch cmd {
	case eat:
		if p.nEaten > 1 {
			return commandLimitExceededError{nEaten: p.nEaten}
		}
		p.nEaten++
	case take:
		if p.has(obj) {
			return objectLimitExceededError{obj: obj, limit: 1}
		}
		p.inventory = append(p.inventory, obj)
	case talk:
		if p.nDialogs > 0 {
			return commandLimitExceededError{nDialogs: p.nDialogs}
		}
		p.nDialogs++
	}
	return nil
}

// newPlayer создает нового игрока
func newPlayer() *player {
	return &player{inventory: []thing{}}
}

// game описывает игру
type game struct {
	// игрок
	player *player
	// объекты игрового мира
	things map[label]int
	// количество успешно выполненных шагов
	nSteps int
}

// has() проверяет, остались ли в игровом мире указанные предметы
func (g *game) has(obj thing) bool {
	count := g.things[obj.name]
	return count > 0
}

// execute() выполняет шаг step
func (g *game) execute(st step) error {
	// проверяем совместимость команды и объекта
	if !st.isValid() {
		return gameOverError{nSteps: g.nSteps, err: invalidStepError{st}}
	}

	// когда игрок берет или съедает предмет,
	// тот пропадает из игрового мира
	if st.cmd == take || st.cmd == eat {
		if !g.has(st.obj) {
			return gameOverError{nSteps: g.nSteps, err: notEnoughObjectsError{st.obj}}
		}
		g.things[st.obj.name]--
	}

	// выполняем команду от имени игрока
	if err := g.player.do(st.cmd, st.obj); err != nil {
		return gameOverError{nSteps: g.nSteps, err: err}
	}

	g.nSteps++
	return nil
}

// newGame() создает новую игру
func newGame() *game {
	p := newPlayer()
	things := map[label]int{
		apple.name:    2,
		coin.name:     3,
		mirror.name:   1,
		mushroom.name: 1,
	}
	return &game{p, things, 0}
}

// giveAdvice() возвращает совет, который
// поможет игроку избежать ошибки err в будущем
func giveAdvice(err error) string {
	var ise invalidStepError
	if errors.As(err, &ise) {
		return fmt.Sprintf("things like '%s %s' are impossible", ise.st.cmd, ise.st.obj)
	}

	var neoe notEnoughObjectsError
	if errors.As(err, &neoe) {
		return fmt.Sprintf("be careful with scarce %ss", neoe.th.name)
	}

	var olee objectLimitExceededError
	if errors.As(err, &olee) {
		return fmt.Sprintf("don't be greedy, %d %s is enough", olee.limit, olee.obj)
	}

	var clee commandLimitExceededError
	if errors.As(err, &clee) {
		switch {
		case clee.nEaten > 0:
			return "eat less"
		case clee.nDialogs > 0:
			return "talk to less"
		}
	}

	return ""
}

// конец решения
