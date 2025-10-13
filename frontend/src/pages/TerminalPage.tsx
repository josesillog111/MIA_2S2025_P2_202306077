import React, { useState, useRef } from 'react'; // Importamos useRef
import { Terminal, Upload, Play, Trash2, ArrowDown } from 'lucide-react';

// Se mantiene el tipo de propiedades de ToolPageProps, aunque 'isReady' ya no se usa para bloquear.
interface ToolPageProps { 
    mountInfo?: any; 
    sessionInfo?: any; 
    onInitialize: () => void;
}

// Estructura de un script
interface ScriptFile {
    name: string;
    content: string;
}

// Datos de scripts de ejemplo
const DUMMY_SCRIPTS: ScriptFile[] = [
    {
        name: 'setup_env.sh',
        content: 'echo "Iniciando configuración de ambiente..."\n/usr/bin/python3 -m venv .venv\nsource .venv/bin/activate\npip install -r requirements.txt\necho "Ambiente listo!"\n',
    },
    {
        name: 'deploy.sh',
        content: 'echo "Preparando despliegue..."\nrsync -avz --delete ./build/ user@server:/var/www/app/\necho "Despliegue completado con éxito."\n',
    },
    {
        name: 'cleanup.sh',
        content: 'echo "Limpiando archivos temporales..."\nfind /tmp -type f -mtime +7 -delete\necho "Limpieza finalizada."\n',
    },
];


// =================================================================
// 1. TerminalPage Componente
// =================================================================

const TerminalPage: React.FC<ToolPageProps> = () => { 
    const [consoleOutput, setConsoleOutput] = useState<string>('Consola lista. Selecciona un script o ingresa un comando.');
    const [manualCommand, setManualCommand] = useState<string>('');
    const [selectedScript, setSelectedScript] = useState<ScriptFile | null>(DUMMY_SCRIPTS[0]);
    const [scriptContent, setScriptContent] = useState<string>(DUMMY_SCRIPTS[0].content);
    const [isLoading, setIsLoading] = useState(false);
    // Referencias para sincronizar el scroll del editor (estilo VS Code)
    const textareaRef = useRef<HTMLTextAreaElement>(null);
    const lineNumberRef = useRef<HTMLDivElement>(null);

    // Función para sincronizar elz scroll entre el textarea y los números de línea
    const handleScrollSync = () => {
        if (textareaRef.current && lineNumberRef.current) {
            lineNumberRef.current.scrollTop = textareaRef.current.scrollTop;
        }
    };

    // Simula la ejecución de código (basado en el snippet del usuario)
    const handleRunCode = async (codeToExecute: string, isManual: boolean) => {
        setIsLoading(true);
        
        const commandName = isManual ? 'Comando Manual' : selectedScript?.name || 'Script Personalizado';
        
        // Actualizamos la salida inmediatamente para indicar que la ejecución ha comenzado
        setConsoleOutput(prev => prev + `\n\n> Ejecutando: ${commandName}...`);

        try {
            const API_URL = "/execute";

            const response = await fetch(API_URL, {
            method: "POST",
            headers: { "Content-Type": "text/plain" },
            body: codeToExecute, // código crudo, sin escapado
            });

            const data = await response.json();

            if (data.success) {
            setConsoleOutput(prev => prev + `\n\n${data.output}`);
            } else {
            setConsoleOutput(prev => prev + `\n\n[ERROR] ${data.error}`);
            }
        } catch (error) {
            setConsoleOutput(
                prev => prev + `\n\n[ERROR] Error de conexión: ${error instanceof Error ? error.message : String(error)}`
            );
        } finally {
            setIsLoading(false);
        }
    };

    const handleRunScript = () => {
        if (scriptContent.trim() && !isLoading) {
            handleRunCode(scriptContent, false);
        }
    };

    const handleRunManualCommand = (e: React.FormEvent) => {
        e.preventDefault();
        if (manualCommand.trim() && !isLoading) {
            handleRunCode(manualCommand.trim(), true); 
            setManualCommand('');
        }
    };

    // Actualizar contenido al seleccionar un script
    const handleScriptChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const selectedName = e.target.value;
        const script = DUMMY_SCRIPTS.find(s => s.name === selectedName) || null;
        setSelectedScript(script);
        if (script) {
            setScriptContent(script.content);
        } else {
            setScriptContent('');
        }
        setConsoleOutput(`Script "${selectedName}" cargado en el editor.`);
    };
    
    // Calcula la cantidad de líneas para los números de línea
    const lineCount = scriptContent.split('\n').length;
    const lineNumbers = Array.from({ length: lineCount }, (_, i) => i + 1);


    return (
        <div className="p-6 h-full bg-[#1e1e1e] text-gray-200 flex flex-col">
            <h2 className="text-3xl font-light text-green-400 mb-4 flex items-center">
                <Terminal className="w-7 h-7 mr-2" /> Terminal de Scripts y Comandos (Acceso Libre)
            </h2>
            
            <div className="flex-1 min-h-0 grid grid-cols-1 lg:grid-cols-3 gap-6">
                
                {/* Columna de Script Editor (2/3 de ancho en desktop) */}
                <div className="lg:col-span-2 flex flex-col bg-[#252526] rounded-xl shadow-lg border border-[#3a3a3a] p-4">
                    <div className="flex justify-between items-center mb-3">
                        <div className="flex items-center space-x-3">
                            <label className="text-gray-400 font-semibold flex items-center"><Upload className="w-4 h-4 mr-2" /> Cargar Script:</label>
                            <select
                                onChange={handleScriptChange}
                                value={selectedScript?.name || ''}
                                className="bg-[#1e1e1e] text-gray-200 p-2 rounded-lg border border-[#3a3a3a] focus:ring-green-500 focus:border-green-500"
                            >
                                {DUMMY_SCRIPTS.map(script => (
                                    <option key={script.name} value={script.name}>{script.name}</option>
                                ))}
                                <option value="">--- (Script Personalizado) ---</option>
                            </select>
                        </div>
                        <button
                            onClick={handleRunScript}
                            disabled={isLoading || !scriptContent.trim()}
                            // Estilo de botón de Ejecutar Script (Verde)
                            className={`flex items-center px-6 py-2 rounded-full transition-all font-semibold shadow-md ${
                                isLoading 
                                ? 'bg-gray-600 text-gray-400 cursor-not-allowed' 
                                : 'bg-gradient-to-r from-green-500 to-emerald-600 hover:from-green-600 hover:to-emerald-700 shadow-green-900 text-white'
                            }`}
                        >
                            {isLoading ? 'Ejecutando...' : 'Ejecutar Script'}
                            <Play className="w-5 h-5 ml-2" />
                        </button>
                    </div>
                    
                    {/* INICIO: Simulación de Editor VS Code */}
                    <div className="flex-1 min-h-[150px] flex overflow-hidden rounded-lg border border-[#3a3a3a] bg-[#1e1e1e]">
                        {/* Columna de Números de Línea */}
                        <div
                            ref={lineNumberRef}
                            className="text-right text-sm text-gray-500 bg-[#252526] py-4 pr-3 select-none overflow-hidden"
                            style={{ 
                                width: '40px', 
                                lineHeight: '20px', // Coincidir con el line-height del textarea
                                flexShrink: 0
                            }}
                        >
                            {lineNumbers.map((num) => (
                                <div key={num} style={{ height: '20px' }}>{num}</div>
                            ))}
                        </div>

                        {/* Textarea (Área de Código) */}
                        <textarea
                            ref={textareaRef}
                            value={scriptContent}
                            onChange={(e) => setScriptContent(e.target.value)}
                            onScroll={handleScrollSync} // Sincroniza el scroll
                            placeholder="# Escribe tu script aquí (ej: Bash, Python, etc.)"
                            className="flex-1 p-4 bg-transparent text-gray-100 font-mono text-sm resize-none focus:outline-none focus:ring-0 focus:border-0 overflow-auto"
                            style={{ 
                                lineHeight: '20px', // Set consistent line height for sync
                                minHeight: '100%' 
                            }}
                        />
                    </div>
                    {/* FIN: Simulación de Editor VS Code */}
                </div>

                {/* Columna de Consola (1/3 de ancho en desktop) */}
                <div className="lg:col-span-1 flex flex-col space-y-4">
                    
                    {/* Área de Comandos Manuales */}
                    <form onSubmit={handleRunManualCommand} className="flex flex-col bg-[#252526] p-4 rounded-xl shadow-lg border border-[#3a3a3a]">
                        <label className="text-lg font-semibold mb-2 text-green-400 flex items-center"><ArrowDown className="w-4 h-4 mr-2"/> Ingrese un Comando Manual</label>
                        <input
                            type="text"
                            value={manualCommand}
                            onChange={(e) => setManualCommand(e.target.value)}
                            placeholder="Ej: ls -l /home/user/"
                            className="w-full p-2 bg-[#1e1e1e] border border-[#3a3a3a] rounded-md text-gray-100 font-mono mb-3 focus:outline-none focus:ring-2 focus:ring-green-500"
                            disabled={isLoading}
                        />
                        <button
                            type="submit"
                            disabled={isLoading || !manualCommand.trim()}
                            // Estilo de botón de Ejecutar Comando (Amarillo)
                            className={`w-full py-2 rounded-full transition-all font-semibold shadow-md text-white ${
                                isLoading 
                                ? 'bg-gray-600 text-gray-400 cursor-not-allowed' 
                                : 'bg-gradient-to-r from-yellow-500 to-orange-500 hover:from-yellow-600 hover:to-orange-600 shadow-yellow-900'
                            }`}
                        >
                            Ejecutar Comando
                        </button>
                    </form>

                    {/* Área de Output de la Consola */}
                    <div className="flex-1 flex flex-col bg-[#000000] p-4 rounded-xl shadow-inner border border-[#3a3a3a] min-h-[250px] lg:min-h-0">
                        <div className="flex justify-between items-center mb-2">
                            <h3 className="text-lg font-semibold text-green-300">Output de Consola</h3>
                            <button
                                onClick={() => setConsoleOutput('Consola lista. Selecciona un script o ingresa un comando.')}
                                title="Limpiar Consola"
                                // Estilo de botón de Limpiar Consola (Icono)
                                className="text-gray-400 hover:text-red-500 bg-[#3a3a3a] hover:bg-red-900/50 p-2 rounded-lg transition-colors"
                            >
                                <Trash2 className="w-4 h-4" />
                            </button>
                        </div>
                        <pre className="flex-1 bg-transparent text-sm text-green-200 font-mono overflow-auto whitespace-pre-wrap">
                            {consoleOutput}
                        </pre>
                    </div>

                </div>
            </div>
        </div>
    );
};

export default TerminalPage;