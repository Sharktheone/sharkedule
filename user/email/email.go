package email

import "errors"

type EMail struct {
	mail string
}

func (e *EMail) GetMail() (string, error) {
	if e.mail == "" {
		return "", errors.New("user has no email")
	}
	return e.mail, nil
}

func (e *EMail) SetMail(mail string) {
	e.mail = mail
}
