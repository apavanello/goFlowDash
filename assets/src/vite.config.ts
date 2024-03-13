import { resolve } from 'path'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vuetify from 'vite-plugin-vuetify'
export default defineConfig({
    resolve: {
        alias: {
            '~': resolve('src'),
        },
        dedupe: ['vue'],
        extensions: ['.ts', '.vue'],
    },
    plugins: [
        vue(),
        vuetify({ autoImport: true }), // Enabled by default
    ],
})