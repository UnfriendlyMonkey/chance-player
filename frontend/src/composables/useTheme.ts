import { ref } from 'vue'
import { useTheme as useVuetifyTheme } from 'vuetify'

export function useTheme() {
  const vuetifyTheme = useVuetifyTheme()
  const isDark = ref(vuetifyTheme.global.name.value === 'dark')

  function toggle() {
    isDark.value = !isDark.value
    vuetifyTheme.global.name.value = isDark.value ? 'dark' : 'light'
    localStorage.setItem('theme', isDark.value ? 'dark' : 'light')
  }

  return { isDark, toggle }
}
