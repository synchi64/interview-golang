package main

import (
	"fmt"
	"os/exec"
	"testing"
)

// none of the tests here are working, basic framework I'll rework
func TestBearerToken(t *testing.T) {
	goodToken := "Bearer secret123"
	badToken := "NOTBearer NOTsecret123"
	commandOne := exec.Command("curl", "-H", "Authorization: "+goodToken, "http://localhost:8080/animals")
	commandTwo := exec.Command("curl", "-H", "Authorization: "+badToken, "http://localhost:8080/animals")

	output, err := commandOne.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:", string(output))

	output, err = commandTwo.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:", string(output))
}

func TestGet(t *testing.T) {
	command := exec.Command("curl", "-H", "Authorization: Bearer secret123", "http://localhost:8080/animals")

	output, err := command.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:", string(output))
}

func TestGetById(t *testing.T) {
	commandOne := exec.Command("curl", "-H", "Authorization: Bearer secret123", "http://localhost:8080/animals/1")
	commandTwo := exec.Command("curl", "-H", "Authorization: Bearer secret123", "http://localhost:8080/animals/2")

	output, err := commandOne.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:", string(output))

	output, err = commandTwo.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:", string(output))
}

func TestPost(t *testing.T) {
	command := exec.Command("curl", "-H", "Authorization: Bearer secret123", "http://localhost:8080/animals", "--include", "--header", "\"Content-Type: application/json\"", "--request", "\"POST\"", "--data", "'{\"id\": 3, \"name\": \"bobcat\", \"species\": \"Lynx rufus\"}'")

	output, err := command.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:", string(output))
}

func TestPatch(t *testing.T) {
	command := exec.Command("curl", "-H", "Authorization: Bearer secret123", "http://localhost:8080/animals/1", "--include", "--header", "\"Content-Type: application/json\"", "--request", "\"PATCH\"", "--data", "'{\"name\": \"bobcat\", \"species\": \"Lynx rufus\"}'")

	output, err := command.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:", string(output))
}

func TestDelete(t *testing.T) {
	command := exec.Command("curl", "-H", "Authorization: Bearer secret123", "http://localhost:8080/animals/1", "--include", "--header", "\"Content-Type: application/json\"", "--request", "\"DELETE\"")

	output, err := command.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Output:", string(output))
}
