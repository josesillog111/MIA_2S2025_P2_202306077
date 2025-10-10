package main

import (
	dmanager "backend/disk_manager"
	parser "backend/parser"
	util "backend/util"

	"encoding/json"
	"io"
	"net/http"

	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

var globalVisitor *util.Visitor
var requestManager *util.RequestManager
var diskManager = dmanager.NewDiskManager()
var listMountedPartitions = dmanager.NewMountedPartitionList()
var idMountedAndLogued = new(string)

func main() {
	fmt.Println("=== Servidor MIA Proyecto ===")

	// Inicializar visitor global
	globalVisitor = util.NewVisitor(diskManager, &listMountedPartitions, idMountedAndLogued)
	requestManager = util.NewRequestManager(diskManager, &listMountedPartitions, idMountedAndLogued)

	// Configura el manejador para la ruta /execute
	http.HandleFunc("/execute", handleExecute)
	http.HandleFunc("/api/fs/list", requestManager.ListHandler)
	http.HandleFunc("/api/fs/read", requestManager.ReadHandler)
	http.HandleFunc("/api/fs/mkdir", requestManager.MkdirHandler)
	http.HandleFunc("/api/fs/create", requestManager.MkfileHandler)
	http.HandleFunc("/api/fs/delete", requestManager.DeleteHandler)

	fmt.Println("Servidor escuchando en http://localhost:8080/execute")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}

func handleExecute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Método no soportado. Usa POST.", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	code := string(body)

	// Ejecutar el parser con el código recibido

	runParser(code)

	var response map[string]interface{}
	if globalVisitor.Errors != "" {
		response = map[string]interface{}{
			"success": false,
			"error":   globalVisitor.Errors,
		}
	} else {
		response = map[string]interface{}{
			"success": true,
			"output":  globalVisitor.Console,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, _ := json.Marshal(response)
	w.Write(jsonResponse)
}

func runParser(code string) {
	input := antlr.NewInputStream(code)

	lexer := parser.NewGoDiskLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	par := parser.NewGoDiskGrammar(tokens)
	par.RemoveErrorListeners()
	par.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	globalVisitor.Console = ""
	globalVisitor.Errors = ""

	tree := par.Prog()

	_ = globalVisitor.Visit(tree)
}
