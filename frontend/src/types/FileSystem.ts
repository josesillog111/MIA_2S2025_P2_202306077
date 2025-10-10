// Define los tipos de nodo (archivo o carpeta)
export type TipoNodo = 'archivo' | 'carpeta';

// Estructura base para cualquier elemento en la lista principal
export interface ElementoArchivo {
  nombre: string;
  tipo: TipoNodo;
  tamano: string; // Ejemplo: '40 elementos' o '9.1. M.'
  modificacion: string; // Fecha de modificación
  ruta: string; // Ruta completa
  esFavorito: boolean;
}

// Estructura para los elementos del sidebar
export interface ElementoSidebar {
  icono: string; // Nombre del icono (ej: 'home', 'download', 'trash')
  nombre: string;
  ruta: string;
  esUbicacionPrincipal: boolean;
}