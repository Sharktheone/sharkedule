import {defineConfig} from "vite"
import react from "@vitejs/plugin-react-swc"

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [react()],
    build: {
        outDir: "web/dist",
    },
    resolve: {
        alias: {
            "@/": "/src/",
            "@kanban/": "/src/pages/task/kanban/"
        }
    }
})
