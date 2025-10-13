package main

import (
	dmanager "backend/disk_manager"
	parser "backend/parser"
	util "backend/util"
	"io"

	"encoding/json"
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
	http.HandleFunc("/disk/list", requestManager.ListDisksHandler)
	http.HandleFunc("/disk/partitions", requestManager.ListPartitionsHandler)
	http.HandleFunc("/disk/checkfs", requestManager.CheckFSHandler)
	http.HandleFunc("/api/fs/list", requestManager.ListHandler)
	http.HandleFunc("/api/fs/read", requestManager.ReadHandler)
	http.HandleFunc("/api/fs/mkdir", requestManager.MkdirHandler)
	http.HandleFunc("/api/fs/create", requestManager.MkfileHandler)
	http.HandleFunc("/api/fs/delete", requestManager.DeleteHandler)

	fmt.Println("Servidor escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}

func handleExecute(w http.ResponseWriter, r *http.Request) {
	// ==== CORS ====
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Método no soportado. Usa POST.", http.StatusMethodNotAllowed)
		return
	}

	// ==== Leer cuerpo como texto plano ====
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	code := string(body)

	// ==== Ejecutar el parser ====
	runParser(code)

	// ==== Construir respuesta ====
	var response struct {
		Success bool   `json:"success"`
		Output  string `json:"output,omitempty"`
		Error   string `json:"error,omitempty"`
	}

	if globalVisitor.Errors != "" {
		response.Success = false
		response.Error = globalVisitor.Errors
	} else {
		response.Success = true
		response.Output = globalVisitor.Console
	}

	// ==== Enviar JSON ====
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error al codificar respuesta JSON", http.StatusInternalServerError)
	}
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
