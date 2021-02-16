package hw10_program_optimization //nolint:golint,stylecheck

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int
type users []User

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r, domain)
	if err != nil {
		return nil, fmt.Errorf("get users error: %s", err)
	}

	return countDomains(u, domain), nil
}

func getUsers(r io.Reader, domain string) (result users, err error) {
	var user User

	lines := getLines(r, domain)

	result = make(users, len(lines))

	for i, line := range lines {
		err := json.Unmarshal([]byte(line), &user)
		if err != nil {
			return nil, fmt.Errorf("unmarshal error: %w", err)
		}

		result[i] = user
	}

	return result, nil
}

func countDomains(u users, domain string) DomainStat {
	result := make(DomainStat)

	for _, user := range u {
		if checkEmail(user.Email, domain) {
			domain := strings.SplitN(strings.ToLower(user.Email), "@", 2)[1]
			count := result[domain]
			count++
			result[domain] = count
		}
	}
	return result
}

func getLines(r io.Reader, domain string) []string {
	var lines []string

	newScanner := bufio.NewScanner(r)
	for newScanner.Scan() {
		str := newScanner.Text()
		if strings.Contains(str, domain) {
			lines = append(lines, newScanner.Text())
		}
	}

	return lines
}

func checkEmail(email string, domain string) bool {
	return strings.HasSuffix(email, "."+domain) && strings.Contains(email, "@")
}
