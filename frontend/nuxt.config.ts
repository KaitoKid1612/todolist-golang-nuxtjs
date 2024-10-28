import { defineNuxtConfig } from "nuxt/config";

export default defineNuxtConfig({
    css: [
        'vuetify/styles',
        '@/assets/css/tailwind.css'
    ],
    build: {
        transpile: ['vuetify']
    },
    typescript: {
        strict: true
    }
})
