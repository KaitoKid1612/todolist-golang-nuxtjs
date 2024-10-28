import { createVuetify } from 'vuetify'
import 'vuetify/styles'

export default defineNuxtPlugin((nuxtApp) => {
    const vuetify = createVuetify()
    nuxtApp.vueApp.use(vuetify)
})
function defineNuxtPlugin(plugin: (nuxtApp: any) => void) {
    return (nuxtApp: any) => {
        plugin(nuxtApp)
    }
}
