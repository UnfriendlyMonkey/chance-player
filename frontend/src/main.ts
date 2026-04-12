import '@mdi/font/css/materialdesignicons.css'
import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createVuetify } from 'vuetify'
import { md3 } from 'vuetify/blueprints'
import 'vuetify/styles'

import App from './App.vue'
import router from './router'
import { i18n } from './i18n'

const savedTheme = localStorage.getItem('theme') ?? 'light'

const vuetify = createVuetify({
  blueprint: md3,
  theme: {
    defaultTheme: savedTheme,
    themes: {
      light: {
        colors: {
          primary: '#6750A4',
          secondary: '#625B71',
          surface: '#FFFBFE',
          background: '#FFFBFE',
        },
      },
      dark: {
        colors: {
          primary: '#D0BCFF',
          secondary: '#CCC2DC',
          surface: '#1C1B1F',
          background: '#1C1B1F',
        },
      },
    },
  },
})

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(vuetify)
app.use(i18n)

app.mount('#app')
