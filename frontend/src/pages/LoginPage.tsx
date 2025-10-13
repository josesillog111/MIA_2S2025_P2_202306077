// src/pages/LoginPage.tsx
import React, { useState } from 'react';
import { User, Lock, ArrowLeft } from 'lucide-react'; // Para los iconos del formulario
import type { WizardPageProps } from '../WizardPageProps';

const LoginPage: React.FC<WizardPageProps> = ({ onLogin, onBack }) => {
    const [username, setUsername] = useState('admin');
    const [password, setPassword] = useState('admin');
    const [error, setError] = useState('');
    const handleLogin = (e: React.FormEvent) => {
        e.preventDefault();
        if (username === 'admin' && password === 'admin') {
            if (onLogin) onLogin({ user: username, token: 'fake-jwt-token' });
        } else {
            setError('Credenciales incorrectas. Usa "admin" / "admin".');
        }
    };
    return (
        <div className="flex w-full h-full bg-[#1e1e1e] text-gray-200">
            <div className="relative flex-1 flex flex-col justify-center items-center p-12 bg-gradient-to-br from-[#2c2c2c] via-[#1a1a1a] to-[#0d0d0d] overflow-hidden rounded-l-lg shadow-lg">
                <div className="relative z-10 text-center max-w-lg"><h1 className="text-4xl lg:text-5xl font-bold mb-4 text-white">Bienvenido de Vuelta</h1><p className="text-lg text-gray-300">Accede a tu espacio de desarrollo personal.</p></div>
            </div>
            <div className="flex-1 flex flex-col justify-center items-center p-12 bg-[#2a2a2a] rounded-r-lg shadow-lg">
                <div className="w-full max-w-md">
                    <h2 className="text-2xl font-semibold mb-6 text-gray-100 text-center">ACCESO DE USUARIO</h2>
                    <form onSubmit={handleLogin} className="space-y-6">
                        <div className="relative"><User className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" /><input type="text" placeholder="Nombre de usuario" value={username} onChange={(e) => setUsername(e.target.value)} className="w-full pl-10 pr-4 py-2 bg-[#3a3a3a] border border-[#4a4a4a] rounded-md text-gray-200 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent" required /></div>
                        <div className="relative"><Lock className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5" /><input type="password" placeholder="Contraseña" value={password} onChange={(e) => setPassword(e.target.value)} className="w-full pl-10 pr-4 py-2 bg-[#3a3a3a] border border-[#4a4a4a] rounded-md text-gray-200 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent" required /></div>
                        {error && (<p className="text-red-400 text-center text-sm">{error}</p>)}
                        <div className="flex gap-4">
                            <button type="button" onClick={onBack} className="w-1/3 py-2 px-4 bg-gray-500 text-white font-semibold rounded-md hover:bg-gray-600 transition-all flex items-center justify-center">
                                <ArrowLeft className="w-4 h-4 inline mr-2" /> Atrás
                            </button>
                            <button type="submit" className="w-2/3 py-2 px-4 bg-gradient-to-r from-blue-600 to-purple-600 text-white font-semibold rounded-md hover:from-blue-700 hover:to-purple-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 focus:ring-offset-[#2a2a2a] transition-all">LOGIN</button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    );
};

export default LoginPage;