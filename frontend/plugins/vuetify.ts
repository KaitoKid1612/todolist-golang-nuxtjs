import { defineNuxtPlugin } from 'nuxt/dist/app/nuxt'
import { createVuetify } from 'vuetify'
import 'vuetify/styles'

export default defineNuxtPlugin((nuxtApp) => {
    const vuetify = createVuetify()
    nuxtApp.vueApp.use(vuetify)
})
