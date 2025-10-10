import { Button } from "../ui/button";
import { ArrowLeft, ArrowRight, X, Minus, Maximize } from "lucide-react";

export default function Toolbar() {
  return (
    <div className="flex items-center justify-between px-3 py-1 bg-[#2c2c2c] text-white shadow-lg select-none border-b border-[#3a3a3a]">
      {/* === Flechas de navegación === */}
      <div className="flex items-center gap-2">
        <Button
          variant="ghost"
          size="icon"
          className="hover:bg-[#3a3a3a] text-gray-300 transition-colors"
        >
          <ArrowLeft className="w-4 h-4" />
        </Button>
        <Button
          variant="ghost"
          size="icon"
          className="hover:bg-[#3a3a3a] text-gray-300 transition-colors"
        >
          <ArrowRight className="w-4 h-4" />
        </Button>
      </div>

      {/* === Título central === */}
      <div className="text-sm font-medium text-gray-200 tracking-wide">
        Explorador de Proyecto
      </div>

      {/* === Botones de ventana === */}
      <div className="flex items-center gap-1">
        <Button
          variant="ghost"
          size="icon"
          className="hover:bg-[#3a3a3a] text-gray-300 transition-colors"
        >
          <Minus className="w-4 h-4" />
        </Button>
        <Button
          variant="ghost"
          size="icon"
          className="hover:bg-[#3a3a3a] text-gray-300 transition-colors"
        >
          <Maximize className="w-4 h-4" />
        </Button>
        <Button
          variant="ghost"
          size="icon"
          className="hover:bg-red-600 text-white transition-colors"
        >
          <X className="w-4 h-4" />
        </Button>
      </div>
    </div>
  );
}

