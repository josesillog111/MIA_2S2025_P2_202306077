import React, { useState, useEffect, useRef } from 'react';
import type { Dispatch, SetStateAction } from "react";
import { Folder, ChevronLeft, Home, FileText, Star, Zap } from 'lucide-react';

// =================================================================
// 1. Tipos de Datos Compartidos
// =================================================================

interface FileItem {
    name: string;
    type: "folder" | "file";
    size: string;
    modified: string;
    favorite: boolean;
}

interface APIFileItem {
    name: string;
    owner: number;
    perm: string;
    size: number;
    time: number; // timestamp
    type: "dir" | "file"; // 'dir' o 'file'
}

interface FileContentProps {
    currentPath: string;
    setCurrentPath: Dispatch<SetStateAction<string>>; 
}


// =================================================================
// 2. Funciones de Ayuda y Simulación
// =================================================================

// Simulación de API para obtener la lista de archivos
const simulateFetch = (path: string): Promise<APIFileItem[]> => {
    console.log(`Simulando fetch para: ${path}`);
    const now = Date.now() / 1000;
    
    // Simulación de datos basada en la ruta
    if (path.includes('src')) {
        return Promise.resolve([
            { name: '.', owner: 1001, perm: 'rwxr-xr-x', size: 4096, time: now, type: 'dir' },
            { name: '..', owner: 1001, perm: 'rwxr-xr-x', size: 4096, time: now, type: 'dir' },
            { name: 'App.tsx', owner: 1001, perm: 'rw-r--r--', size: 10240, time: now - 3600, type: 'file' },
            { name: 'TerminalPage.tsx', owner: 1001, perm: 'rw-r--r--', size: 15360, time: now - 1800, type: 'file' },
            { name: 'components', owner: 1001, perm: 'rwxr-xr-x', size: 4096, time: now - 60, type: 'dir' },
        ]);
    }

    if (path.includes('components')) {
        return Promise.resolve([
            { name: '.', owner: 1001, perm: 'rwxr-xr-x', size: 4096, time: now, type: 'dir' },
            { name: '..', owner: 1001, perm: 'rwxr-xr-x', size: 4096, time: now, type: 'dir' },
            { name: 'Button.tsx', owner: 1001, perm: 'rw-r--r--', size: 2560, time: now - 300, type: 'file' },
            { name: 'Icon.tsx', owner: 1001, perm: 'rw-r--r--', size: 1500, time: now - 120, type: 'file' },
        ]);
    }
    
    // Raíz o ruta por defecto
    return Promise.resolve([
        { name: '.', owner: 1001, perm: 'rwxr-xr-x', size: 4096, time: now, type: 'dir' },
        { name: '..', owner: 1001, perm: 'rwxr-xr-x', size: 4096, time: now, type: 'dir' },
        { name: 'src', owner: 1001, perm: 'rwxr-xr-x', size: 4096, time: now - 7200, type: 'dir' },
        { name: 'node_modules', owner: 1001, perm: 'rwxr-xr-x', size: 81920, time: now - 18000, type: 'dir' },
        { name: 'README.md', owner: 1001, perm: 'rw-r--r--', size: 512, time: now - 3600, type: 'file' },
        { name: 'package.json', owner: 1001, perm: 'rw-r--r--', size: 1280, time: now - 3600, type: 'file' },
        { name: 'config.yaml', owner: 1001, perm: 'rw-r--r--', size: 800, time: now - 120, type: 'file' },
    ]);
};

// Función para obtener el ícono según el tipo/extensión
const getFileIcon = (item: FileItem) => {
    if (item.type === "folder") {
        return <Folder className="w-5 h-5 text-yellow-500" />;
    }
    const ext = item.name.split('.').pop()?.toLowerCase();
    switch (ext) {
        case 'tsx':
        case 'jsx':
            return <Zap className="w-5 h-5 text-blue-400" />;
        case 'sh':
        case 'yaml':
        case 'md':
            return <FileText className="w-5 h-5 text-gray-400" />;
        case 'json':
            return <FileText className="w-5 h-5 text-green-400" />;
        default:
            return <FileText className="w-5 h-5 text-gray-500" />;
    }
}


// =================================================================
// 3. FileContent Componente (Lista y Cuadrícula)
// =================================================================

function FileContent({ currentPath, setCurrentPath }: FileContentProps) {
    const [view, setView] = useState("list"); // "grid" o "list"
    const [items, setItems] = useState<FileItem[]>([]);
    const tableRef = useRef<HTMLTableElement>(null);
    const [favorites, setFavorites] = useState<string[]>([]);


  useEffect(() => {
    const fetchItems = async () => {
      console.log(`Fetching: ${currentPath}`);

      try {
        const apiData: APIFileItem[] = await simulateFetch(currentPath);

        const formattedItems: FileItem[] = apiData
          .filter(item => item.name !== "." && item.name !== "..")
          .map(item => {
            const date = new Date(item.time * 1000).toLocaleDateString("es-ES", {
              year: 'numeric', month: '2-digit', day: '2-digit'
            });
            
            const sizeString = item.type === "dir" ? "-" : 
                item.size < 1024 ? `${item.size} B` : 
                item.size < 1024 * 1024 ? `${(item.size / 1024).toFixed(1)} KB` : 
                `${(item.size / (1024 * 1024)).toFixed(1)} MB`;
            
            return {
              name: item.name,
              type: item.type === "dir" ? "folder" : "file",
              size: sizeString,
              modified: date,
              favorite: favorites.includes(item.name), 
            };
          });

        setItems(formattedItems);

      } catch (error) {
        console.error("Error fetching file list:", error);
        setItems([]);
      }
    };

    fetchItems();
  }, [currentPath, favorites]);

  // Función: Alternar estado favorito
  const toggleFavorite = (nameToToggle: string) => {
    setFavorites(prevFavorites => {
        if (prevFavorites.includes(nameToToggle)) {
            return prevFavorites.filter(name => name !== nameToToggle);
        } else {
            return [...prevFavorites, nameToToggle];
        }
    });
  };

  // Función para manejar la navegación o apertura de archivos
  const handleItemClick = (item: FileItem) => {
    if (item.type === 'folder') {
        setCurrentPath(prev => {
            const path = prev.endsWith('/') ? prev.slice(0, -1) : prev;
            return `${path}/${item.name}`;
        });
    } else {
        console.log(`Abriendo archivo: ${item.name}`);
        // TO DO: Abrir archivo en el editor de código.
    }
  };

  // Lógica para el redimensionamiento de columnas
  useEffect(() => {
    if (view !== "list" || !tableRef.current) {
      return;
    }

    const tableElement = tableRef.current;
    const cols = tableElement.querySelectorAll("th") as NodeListOf<HTMLElement>;

    const startResize = (e: MouseEvent) => {
        e.preventDefault();
        const col = (e.currentTarget as HTMLElement).parentElement as HTMLElement;
        const startX = e.pageX;
        const startWidth = col.offsetWidth;

        const resize = (e: MouseEvent) => {
          const newWidth = startWidth + (e.pageX - startX);
          if (newWidth > 50) {
              col.style.width = `${newWidth}px`;
              col.style.minWidth = `${newWidth}px`;
          }
        };

        const stopResize = () => {
          window.removeEventListener("mousemove", resize);
          window.removeEventListener("mouseup", stopResize);
        };

        window.addEventListener("mousemove", resize);
        window.addEventListener("mouseup", stopResize);
      };

    // Aplicar resizers
    cols.forEach((col) => {
      if (!col.querySelector(".resizer")) {
        const resizer = document.createElement("div");
        resizer.classList.add("resizer");
        col.appendChild(resizer);
        resizer.addEventListener("mousedown", startResize);
      }
    });


    // Limpieza: Esta función de limpieza es necesaria si el componente se desmonta.
    // Aunque no está completa en el original, se mantiene la estructura.
    return () => {
        const currentResizers = tableElement.querySelectorAll(".resizer") as NodeListOf<HTMLElement>;
        currentResizers.forEach(r => r.removeEventListener("mousedown", startResize));
    };

  }, [view, items.length]); 

  return (
    <div className="flex-1 bg-[#252526] text-gray-200 p-4 overflow-hidden rounded-b-xl flex flex-col border border-[#3a3a3a]"> 
      
      {/* === Controles de Vista === */}
      <div className="flex justify-between items-center mb-4 border-b border-[#3a3a3a] pb-3">
        <h2 className="text-sm text-gray-400 font-medium">Archivos ({items.length})</h2>
        <div className="flex gap-2 bg-[#1e1e1e] p-1 rounded-lg">
          <button
            onClick={() => setView("list")}
            title="Vista de Lista"
            className={`p-2 rounded-md transition-colors ${
              view === "list" ? "bg-blue-600 text-white shadow-md" : "hover:bg-[#3a3a3a] text-gray-400"
            }`}
          >
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><line x1="8" y1="6" x2="21" y2="6"></line><line x1="8" y1="12" x2="21" y2="12"></line><line x1="8" y1="18" x2="21" y2="18"></line><line x1="3" y1="6" x2="3.01" y2="6"></line><line x1="3" y1="12" x2="3.01" y2="12"></line><line x1="3" y1="18" x2="3.01" y2="18"></line></svg>
          </button>
          <button
            onClick={() => setView("grid")}
            title="Vista de Cuadrícula"
            className={`p-2 rounded-md transition-colors ${
              view === "grid" ? "bg-blue-600 text-white shadow-md" : "hover:bg-[#3a3a3a] text-gray-400"
            }`}
          >
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><rect x="3" y="3" width="7" height="7"></rect><rect x="14" y="3" width="7" height="7"></rect><rect x="14" y="14" width="7" height="7"></rect><rect x="3" y="14" width="7" height="7"></rect></svg>
          </button>
        </div>
      </div>

      {/* === Vista del contenido === */}
      {view === "grid" ? (
        <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 gap-x-4 gap-y-6 overflow-y-auto p-2">
          {items.map((item, i) => (
            <div
              key={i}
              className="flex flex-col items-center cursor-pointer p-4 hover:bg-[#383838] rounded-xl transition-colors group"
              onClick={() => handleItemClick(item)}
            >
              <div className={`text-5xl mb-3 transition-transform group-hover:scale-105 ${item.type === 'folder' ? 'text-yellow-500' : 'text-gray-400'}`}>
                {getFileIcon(item)}
              </div>
              <div className="text-sm font-medium text-gray-200 w-full text-center truncate">{item.name}</div>
              <div className="text-xs text-gray-500 mt-1">{item.size}</div>
            </div>
          ))}
        </div>
      ) : (
        // 📄 Vista en tabla/lista
        <div className="relative flex-1 min-h-0">
          <div className="overflow-y-auto h-full rounded-lg">
            <table
              ref={tableRef}
              className="w-full text-sm text-gray-300"
            >
              {/* Encabezado fijo con sticky */}
              <thead className="bg-[#1e1e1e] text-blue-300 uppercase text-xs sticky top-0 shadow-md z-10 select-none">
                <tr className="border-b border-[#3a3a3a]">
                  <th className="px-4 py-3 text-left relative min-w-[200px]" style={{ width: '45%' }}>Nombre</th>
                  <th className="px-4 py-3 text-right relative min-w-[100px]" style={{ width: '20%' }}>Tamaño</th>
                  <th className="px-4 py-3 text-right relative min-w-[150px]" style={{ width: '25%' }}>Modificación</th>
                  <th className="px-4 py-3 text-center relative min-w-[50px]" style={{ width: '10%' }}>Favorito</th>
                </tr>
              </thead>
              <tbody>
                {items.map((item, i) => (
                  <tr
                    key={i}
                    className={`transition-colors border-t border-[#3a3a3a] ${
                        i % 2 === 0 ? 'bg-[#2a2a2a]' : 'bg-[#252526]'
                    } hover:bg-[#383838] cursor-pointer`}
                    onDoubleClick={() => handleItemClick(item)}
                  >
                    <td className="px-4 py-3 flex items-center gap-3 font-medium">
                      {getFileIcon(item)}
                      <span className="truncate">{item.name}</span>
                    </td>
                    <td className="px-4 py-3 text-right font-mono text-gray-400">{item.size}</td>
                    <td className="px-4 py-3 text-right text-gray-400">{item.modified}</td>
                    <td 
                        className="px-4 py-3 text-center text-xl hover:scale-110 transition-transform"
                        onClick={() => toggleFavorite(item.name)}
                    >
                      {item.favorite ? <Star className="w-5 h-5 text-yellow-400 fill-yellow-400 inline-block"/> : <Star className="w-5 h-5 text-gray-600 inline-block"/>}
                    </td>
                  </tr>
                ))}
                {items.length === 0 && (
                    <tr>
                        <td colSpan={4} className="text-center py-8 text-gray-500 italic">
                            Esta carpeta está vacía.
                        </td>
                    </tr>
                )}
              </tbody>
            </table>
          </div>

          {/* Estilo para el resizer (permite redimensionamiento de columnas) */}
          <style>{`
            .resizer {
              position: absolute;
              right: -2px; 
              top: 0;
              width: 5px;
              height: 100%;
              cursor: col-resize;
              user-select: none;
              transition: background-color 0.2s;
            }
            th {
              position: relative;
            }
            th .resizer:hover {
              background-color: #007bff; 
            }
          `}</style>
        </div>
      )}
    </div>
  );
}


// =================================================================
// 4. FileExplorer Componente (Ruta y Navegación) - Contenedor Principal
// =================================================================

const FileExplorerPage: React.FC = () => {
    // Estado de la ruta actual, simulando el directorio raíz.
    const [currentPath, setCurrentPath] = useState<string>('/home/user');

    console.log("FileExplorerPage renderizado. currentPath:", currentPath);

    const navigateToPath = (path: string) => {
        // Simple sanitización de ruta
        const cleanPath = path.startsWith('/') ? path : `/${path}`;
        setCurrentPath(cleanPath);
    };

    const navigateUp = () => {
        if (currentPath === '/') return;
        
        const pathSegments = currentPath.split('/').filter(Boolean);
        if (pathSegments.length > 0) {
            pathSegments.pop();
            navigateToPath('/' + pathSegments.join('/'));
        } else {
            navigateToPath('/');
        }
    };
    
    // Divide la ruta para crear migas de pan interactivas
    const pathSegments = currentPath.split('/').filter(Boolean);
    let pathAccumulator = '';


    return (
        <div className="p-6 h-full bg-[#1e1e1e] text-gray-200 flex flex-col">
            <h2 className="text-3xl font-light text-blue-400 mb-4 flex items-center">
                <Folder className="w-7 h-7 mr-2" /> Explorador de Archivos
            </h2>

            {/* BARRA DE RUTA DE NAVEGACIÓN */}
            <div className="bg-[#252526] p-3 rounded-t-xl shadow-lg border-b border-[#3a3a3a] flex items-center space-x-2">
                
                {/* Botón de Inicio */}
                <button 
                    onClick={() => navigateToPath('/home/user')}
                    className="p-1 rounded-full text-blue-400 hover:bg-[#383838] transition-colors"
                    title="Ir a la raíz de usuario"
                >
                    <Home className="w-5 h-5" />
                </button>
                
                {/* Botón de Retroceso */}
                <button 
                    onClick={navigateUp}
                    // Deshabilitar si ya estamos en la "raíz" simulada (/home/user)
                    disabled={currentPath === '/home/user'}
                    className="p-1 rounded-full text-gray-400 disabled:opacity-50 hover:bg-[#383838] transition-colors"
                    title="Carpeta anterior"
                >
                    <ChevronLeft className="w-5 h-5" />
                </button>

                {/* Migas de Pan (Breadcrumbs) */}
                <span className="text-sm font-mono text-gray-400 flex items-center overflow-x-auto whitespace-nowrap">
                    {/* Inicio de ruta para la raíz de usuario */}
                    <span className="text-gray-500">~</span> 
                    
                    {pathSegments.map((segment, index) => {
                        pathAccumulator += '/' + segment;
                        const fullPath = pathAccumulator;
                        const isLast = index === pathSegments.length - 1;
                        
                        return (
                            <React.Fragment key={index}>
                                <span className="mx-1 text-gray-600">/</span>
                                <button
                                    onClick={() => navigateToPath(fullPath)}
                                    className={`transition-colors ${isLast ? 'text-gray-300 font-semibold' : 'text-blue-400 hover:text-blue-300'}`}
                                >
                                    {segment}
                                </button>
                            </React.Fragment>
                        );
                    })}
                </span>
            </div>

            {/* CONTENIDO DEL EXPLORADOR */}
            <FileContent currentPath={currentPath} setCurrentPath={setCurrentPath} />
        </div>
    );
}

export default FileExplorerPage;