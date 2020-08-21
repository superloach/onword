package onword

import (
	"fmt"
	"strings"

	"github.com/superloach/onword/value"
)

type Ses struct {
	ID     uint
	Onword *Onword
	Queue  *Queue
	Input  chan string
	Output chan string
}

func (ses *Ses) Loop() {
	user := (*UserDef)(nil)
	hint := []string(nil)

	for {
		fmt.Println(*ses.Queue)

		line := <-ses.Input

		if len(line) == 0 {
			continue
		}

		for _, word := range strings.Split(line, " ") {
			d := Def(nil)

			if hint == nil {
				if word == "(" {
					hint = []string{}
					continue
				}
			}

			if hint != nil {
				if word == ")" {
					d = HintDef(strings.Join(hint, " "))
					hint = nil
				} else {
					hint = append(hint, word)
					continue
				}
			}

			if d == nil && hint == nil && word[0] == ';' {
				if user != nil {
					ses.Onword.Defs[user.Name] = user
					user = nil
				}

				if len(word) > 1 {
					user = &UserDef{
						Name: word[1:],
					}
				}

				continue
			}

			if d == nil && hint == nil && word[0] == '\'' {
				d = PushDef{value.String(word[1:])}
			}

			if d == nil && hint == nil {
				od, ok := ses.Onword.Defs[word]
				if !ok {
					ses.Output <- word + "?"
					continue
				}

				d = od
			}

			if user != nil {
				user.Defs = append(user.Defs, d)
				continue
			}

			err := d.Call(ses)
			if err != nil {
				err = fmt.Errorf("call %q: %w", word, err)
				ses.Output <- fmt.Sprintf("error: %s", err)
			}
		}
	}
}

func (ses *Ses) Close() {
	delete(ses.Onword.Sess, ses.ID)

	close(ses.Input)

	ses.Output <- "session closed"
	close(ses.Output)
}
