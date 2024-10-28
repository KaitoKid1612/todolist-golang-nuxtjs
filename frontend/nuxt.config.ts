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

function defineNuxtConfig(config: { css: string[]; build: { transpile: string[]; }; typescript: { strict: boolean; }; }) {
    return config;
}

