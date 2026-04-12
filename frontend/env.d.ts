/// <reference types="vite/client" />

// Vuetify styles have no TS declarations — declare as side-effect module
declare module 'vuetify/styles' {
  const styles: unknown
  export default styles
}
