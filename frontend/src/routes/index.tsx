import { createFileRoute } from '@tanstack/react-router'
import ShortenPage from '#/features/shorten/components/page'

export const Route = createFileRoute('/')({ component: ShortenPage })
