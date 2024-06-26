package lang

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"strconv"
	"strings"

	"github.com/roman-mazur/architecture-lab-3/painter"
)

type Parser struct {
	state State
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	p.state.ResetOperations()

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		cmdl := scanner.Text()

		err := p.parse(cmdl)
		if err != nil {
			return nil, err
		}
	}
	res := p.state.GetOperations()

	return res, nil
}
func (p *Parser) parse(cmdl string) error {
	words := strings.Split(cmdl, " ")
	command := words[0]

	switch command {
	case "white":
		if len(words) != 1 {
			return fmt.Errorf("Error! Too many or none visible arguments for white command")
		}
		p.state.WhiteBackground()
	case "green":
		if len(words) != 1 {
			return fmt.Errorf("Error! Too many or none visible arguments for green command")
		}
		p.state.GreenBackground()
	case "update":
		if len(words) != 1 {
			return fmt.Errorf("Error! Too many or none visible arguments for update command")
		}
		p.state.SetUpdateOperation()
	case "bgrect":
		parameters, err := checkForErrorsInParameters(words, 5)
		if err != nil {
			return err
		}
		p.state.BackgroundRectangle(image.Point{X: parameters[0], Y: parameters[1]}, image.Point{X: parameters[2], Y: parameters[3]})
	case "figure":
		parameters, err := checkForErrorsInParameters(words, 3)
		if err != nil {
			return err
		}
		p.state.AddFigure(image.Point{X: parameters[0], Y: parameters[1]})
	case "move":
		parameters, err := checkForErrorsInParameters(words, 3)
		if err != nil {
			return err
		}
		p.state.AddMoveOperation(parameters[0], parameters[1])
	case "reset":
		if len(words) != 1 {
			return fmt.Errorf("Error! Too many or none visible arguments for reset command")
		}
		p.state.ResetStateAndBackground()
	default:
		return fmt.Errorf("Error! Invalid %v command ", words[0])
	}
	return nil
}

func checkForErrorsInParameters(words []string, expected int) ([]int, error) {
	if len(words) != expected {
		return nil, fmt.Errorf("Error! Too many or none visible arguments for %v command", words[0])
	}
	var command = words[0]
	var params []int
	for _, param := range words[1:] {
		p, err := parseInt(param)
		if err != nil {
			return nil, fmt.Errorf("Error! Invalid parameter for %s command.Command %s is not a number", command, param)
		}
		params = append(params, p)
	}
	return params, nil
}

func parseInt(s string) (int, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("Error! Can't parse float %s", s)
	}
	return int(f * 400), nil
}
