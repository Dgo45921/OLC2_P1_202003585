package controllers

import (
	"PY1/environment"
	"PY1/interfaces"
	"PY1/models"
	"PY1/parser"
	"encoding/json"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var lastGivencode = ""

type TreeShapeListener struct {
	*parser.BaseSwiftGrammarListener
	Code []interface{}
}

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello skrrr!")
	if err != nil {
		return
	}
}
func Parse(w http.ResponseWriter, r *http.Request) {
	// newCode is responsible to save the given input
	var newCode models.SourceCode
	// consoleResponse is responsible of returning all of the console logs
	var consoleResponse models.ConsoleResponse
	// getting the body from the request
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "ERROR")
	}

	// parsing the json
	json.Unmarshal(reqBody, &newCode)
	// printing the input
	fmt.Println(newCode.Code)
	lastGivencode = newCode.Code

	err = writeSourceCodeFile("source.txt", lastGivencode)
	if err != nil {
		fmt.Println("Error:", err)
	}

	//Entrada
	var code string = newCode.Code
	//Leyendo entrada
	input := antlr.NewInputStream(code)
	lexer := parser.NewSwiftLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	//creacion de parser
	p := parser.NewSwiftGrammarParser(tokens)
	p.BuildParseTrees = true
	tree := p.S()
	//listener
	var listener *TreeShapeListener = NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	Code := listener.Code
	//create ast
	var Ast environment.AST
	// creating env
	var newEnv = environment.NewEnvironment(nil, environment.GLOBAL)
	//ejecución
	for _, inst := range Code {
		inst.(interfaces.Instruction).Execute(&Ast, newEnv)
	}
	fmt.Println(Ast.GetPrint())

	consoleResponse.Console = Ast.GetPrint()
	json.NewEncoder(w).Encode(consoleResponse)
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitS(ctx *parser.SContext) {
	this.Code = ctx.GetCode()
}

func writeSourceCodeFile(filename, content string) error {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func GetCST(w http.ResponseWriter, r *http.Request) {

	exePath, errr := os.Executable()
	if errr != nil {
		log.Fatal(errr)
	}

	currentDir := filepath.Dir(exePath)
	cmd := exec.Command("antlr4-parse", "CG.g4", "s", "-gui", "source.txt")
	cmd.Dir = currentDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
