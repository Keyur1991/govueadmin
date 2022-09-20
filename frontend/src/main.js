import { createApp } from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
import router from "./router/index"
import urls from './plugins/urls'
import { $axios } from './plugins/axios'
import store from './store'

loadFonts()

const app = createApp(App)

app.provide('axios', $axios)
app.provide('urls', urls)
app.use(vuetify)
app.use(router)
app.use(store)
app.mount('#app')