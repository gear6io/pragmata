import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import tailwindcss from '@tailwindcss/vite';

export default defineConfig({
  plugins: [react(), tailwindcss()],
  base: '/',
  define: {
    __API_BASE_URL__: JSON.stringify(process.env.VITE_API_URL ?? ''),
  },
  build: { outDir: 'dist' },
  server: {
    proxy: {
      '/api': 'http://localhost:7181',
    },
  },
});
