import { createI18n } from 'vue-i18n'
import en from './locales/en.json'
import ru from './locales/ru.json'

export type MessageSchema = typeof en

export const i18n = createI18n<[MessageSchema], 'en' | 'ru'>({
  legacy: false,
  locale: (localStorage.getItem('locale') as 'en' | 'ru') ?? 'en',
  fallbackLocale: 'en',
  messages: { en, ru },
})
