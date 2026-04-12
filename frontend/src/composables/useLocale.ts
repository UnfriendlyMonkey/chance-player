import { useI18n } from 'vue-i18n'

export function useLocale() {
  const { locale } = useI18n()

  const locales = [
    { code: 'en', label: 'EN' },
    { code: 'ru', label: 'RU' },
  ]

  function setLocale(code: string) {
    locale.value = code
    localStorage.setItem('locale', code)
  }

  return { locale, locales, setLocale }
}
