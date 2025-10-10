import { useState } from "react";

export default function FileContent() {
  const [view, setView] = useState("grid"); // "grid" o "list"

  // Datos de ejemplo
  const items = [
    { name: "Documentos", type: "folder", size: "-", modified: "2025-10-09", favorite: false },
    { name: "Proyecto.go", type: "file", size: "14 KB", modified: "2025-10-07", favorite: true },
    { name: "Notas.txt", type: "file", size: "2 KB", modified: "2025-10-05", favorite: false },
    { name: "Imágenes", type: "folder", size: "-", modified: "2025-10-01", favorite: false },
  ];

  return (
    <div className="flex-1 bg-[#1e1e1e] text-gray-200 p-4 overflow-y-auto rounded-b-2xl">
      {/* === Encabezado === */}
      <div className="flex justify-between items-center mb-3">
        <h2 className="text-sm text-gray-400 font-medium">Contenido actual</h2>
        <div className="flex gap-2">
          <button
            onClick={() => setView("list")}
            className={`px-3 py-1 rounded-md text-xs transition-colors ${
              view === "list" ? "bg-[#3a3a3a]" : "hover:bg-[#2c2c2c]"
            }`}
          >
            📄 Lista
          </button>
          <button
            onClick={() => setView("grid")}
            className={`px-3 py-1 rounded-md text-xs transition-colors ${
              view === "grid" ? "bg-[#3a3a3a]" : "hover:bg-[#2c2c2c]"
            }`}
          >
            🧩 Cuadrícula
          </button>
        </div>
      </div>

      {/* === Vista del contenido === */}
      {view === "grid" ? (
        // 🟦 Vista en cuadrícula
        <div className="grid grid-cols-4 gap-4">
          {items.map((item, i) => (
            <div
              key={i}
              className="flex flex-col items-center justify-center py-6 rounded-xl bg-[#2a2a2a] hover:bg-[#383838] cursor-pointer transition-colors"
            >
              <div className="text-3xl mb-2">
                {item.type === "folder" ? "📁" : "📄"}
              </div>
              <div className="text-sm text-gray-300 truncate">{item.name}</div>
            </div>
          ))}
        </div>
      ) : (
        // 🟨 Vista en lista (tabla)
        <div className="overflow-x-auto rounded-lg border border-[#2a2a2a]">
          <table className="w-full text-sm text-gray-300">
            <thead className="bg-[#2a2a2a] text-gray-400 uppercase text-xs">
              <tr>
                <th className="px-4 py-3 text-left">Nombre</th>
                <th className="px-4 py-3 text-left">Tamaño</th>
                <th className="px-4 py-3 text-left">Modificación</th>
                <th className="px-4 py-3 text-center">Favorito</th>
              </tr>
            </thead>
            <tbody>
              {items.map((item, i) => (
                <tr
                  key={i}
                  className="hover:bg-[#383838] transition-colors cursor-pointer"
                >
                  <td className="px-4 py-2 flex items-center gap-2">
                    <span className="text-xl">
                      {item.type === "folder" ? "📁" : "📄"}
                    </span>
                    {item.name}
                  </td>
                  <td className="px-4 py-2">{item.size}</td>
                  <td className="px-4 py-2">{item.modified}</td>
                  <td className="px-4 py-2 text-center">
                    {item.favorite ? "⭐" : "☆"}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}
