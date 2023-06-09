import {defineConfig} from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [react(
        {
            include: [/\.(tsx|ts|jsx|js)?$/],
        }
    )],
    build: {
        outDir: 'web/dist',
    },
    resolve: {
        alias: {
            '@/': '/src/',
            '@kanban/': '/src/pages/task/kanban/'
        }
    }
})
