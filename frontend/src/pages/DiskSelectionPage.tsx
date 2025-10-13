// src/pages/DiskSelectionPage.tsx

import React, { useState, useEffect } from 'react';
import { HardDrive, ArrowRight } from 'lucide-react';
import type { WizardPageProps } from '../WizardPageProps';

// TypeScript interface for disk records (optional, but good practice)
interface DiskRecord {
    id: string; // The backend doesn't return an id, we'll create one.
    name: string;
    path: string;
    size: string; // Updated to string since the Go backend returns it as a string
    date: string;
}

// B1. DiskSelectionPage
const DiskSelectionPage: React.FC<WizardPageProps> = ({ onSelect }) => {
    // State to hold the fetched disk data
    const [disks, setDisks] = useState<DiskRecord[]>([]);
    // State for the currently selected disk ID
    const [selectedDisk, setSelectedDisk] = useState<string | null>(null);
    // State for loading status
    const [isLoading, setIsLoading] = useState(true);
    // State for error handling
    const [error, setError] = useState<string | null>(null);

    // useEffect hook to fetch data when the component mounts
    useEffect(() => {
    const fetchDisks = async () => {
        try {
            const API_URL = "/disk/list";

            const response = await fetch(API_URL);
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            const data = await response.json();

            // 1. Validar la estructura esperada del JSON: 
            //    Ahora se espera 'data.disks' (en minúsculas)
            if (!data || !Array.isArray(data.disks)) {
                throw new Error("Formato inválido recibido desde el servidor. Se esperaba un objeto con la propiedad 'disks'.");
            }

            const formatSize = (sizeInBytes: number): string => {
                if (sizeInBytes >= 1e12) return (sizeInBytes / 1e12).toFixed(2) + ' TB';
                if (sizeInBytes >= 1e9) return (sizeInBytes / 1e9).toFixed(2) + ' GB';
                if (sizeInBytes >= 1e6) return (sizeInBytes / 1e6).toFixed(2) + ' MB';
                if (sizeInBytes >= 1e3) return (sizeInBytes / 1e3).toFixed(2) + ' KB';
                return sizeInBytes + ' B';
            }

            // 2. Mapear y asignar ID, usando claves en minúsculas (disks, path, name, size, date)
            const disksWithId = data.disks.map((disk: any) => ({
                id: disk.path, // clave única, usando 'path'
                name: disk.name,
                path: disk.path,
                size: formatSize(disk.size), // Convertir a número y formatear
                date: disk.date,
            }));

            // Actualizar el estado
            setDisks(disksWithId);
            setError(null);
            setIsLoading(false);

            // Si no hay disco seleccionado o el anterior ya no existe, selecciona el primero
            if (disksWithId.length > 0) {
                setSelectedDisk(prev => {
                    // La verificación 'stillExists' debería usar la interfaz DiskRecord
                    const stillExists = disksWithId.some((d: any) => d.id === prev); // Usamos 'any' si DiskRecord no está actualizada
                    return stillExists ? prev : disksWithId[0].id;
                });
            } else {
                 setSelectedDisk(null);
            }
        } catch (e: any) {
            setError(`No se pudieron cargar los discos: ${e.message}. Verifica el servidor o el archivo disks.json.`);
            setIsLoading(false);
        }
    };

    fetchDisks();
    }, []); // Empty dependency array ensures this runs only once on mount

    const handleNext = (e: React.FormEvent) => {
        e.preventDefault();
        console.log("Disco seleccionado ID:", selectedDisk);
        const disk = disks.find(d => d.id === selectedDisk);
        if (disk && onSelect) {
            onSelect({ id: disk.id, name: disk.name });
        }
    };
    
    // Conditional rendering based on state
    if (isLoading) {
        return (
            <div className="flex w-full h-full justify-center items-center bg-[#1e1e1e] text-white">
                <p>Cargando discos...</p>
            </div>
        );
    }
    
    if (error) {
        return (
            <div className="flex w-full h-full justify-center items-center bg-[#1e1e1e] text-red-400">
                <p>{error}</p>
            </div>
        );
    }
    
    // Main component rendering logic
    return (
        <div className="flex w-full h-full bg-[#1e1e1e] text-gray-200 shadow-xl rounded-lg overflow-hidden">
            <div className="flex-1 flex flex-col justify-center items-center p-12 bg-gradient-to-br from-[#2c2c2c] via-[#1a1a1a] to-[#0d0d0d] text-center">
                <HardDrive className="w-16 h-16 text-blue-400 mb-4" />
                <h1 className="text-4xl font-bold mb-4 text-white">Paso 1: Selección de Disco</h1>
                <p className="text-lg text-gray-300 max-w-sm">Elige el disco principal que deseas montar para acceder a su contenido.</p>
            </div>
            <div className="flex-1 flex flex-col justify-center items-center p-12 bg-[#2a2a2a]">
                <div className="w-full max-w-md">
                    <h2 className="text-2xl font-semibold mb-6 text-gray-100 text-center">SELECCIONA UNA UNIDAD</h2>
                    <form onSubmit={handleNext} className="space-y-6">
                        {disks.length > 0 ? (
                            disks.map(disk => (
                                <div
                                    key={disk.id}
                                    onClick={() => setSelectedDisk(disk.id)}
                                    className={`flex items-center justify-between p-4 cursor-pointer bg-[#3a3a3a] border rounded-md transition-all ${selectedDisk === disk.id ? 'border-blue-500 ring-2 ring-blue-500' : 'border-[#4a4a4a] hover:border-gray-500'}`}
                                >
                                    <div className="flex items-center gap-3">
                                        <HardDrive className="w-5 h-5 text-gray-400" />
                                        <div>
                                            <p className="font-medium">{disk.name}</p>
                                            <p className="text-xs text-gray-400">{disk.size}</p>
                                        </div>
                                    </div>
                                    {selectedDisk === disk.id && <span className="text-blue-400">Seleccionado</span>}
                                </div>
                            ))
                        ) : (
                            <p className="text-gray-400 text-center">No hay discos disponibles.</p>
                        )}
                        <button
                            type="submit"
                            disabled={!selectedDisk}
                            className={`w-full py-3 px-4 text-white font-semibold rounded-md transition-all ${selectedDisk ? 'bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700' : 'bg-gray-500 cursor-not-allowed'}`}
                        >
                            Siguiente Paso <ArrowRight className="w-4 h-4 inline ml-2" />
                        </button>
                    </form>
                </div>
            </div>
        </div>
    );
};

export default DiskSelectionPage;