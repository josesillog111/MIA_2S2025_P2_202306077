import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    // Configuración del Proxy
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true, // Esto es importante para el virtual hosting
      },
      '/execute': { 
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
      '/disk': { 
        target: 'http://localhost:8080',
        changeOrigin: true,
      }
    }
  }
})
