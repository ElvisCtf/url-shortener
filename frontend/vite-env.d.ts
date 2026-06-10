interface ImportMetaEnv {
  readonly VITE_APP_ENV: string
  readonly VITE_APP_TITLE: string
  readonly VITE_API_CLIENT_URL: string
  readonly VITE_API_CLIENT_TIMEOUT: string
  readonly VITE_API_ADMIN_URL: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
