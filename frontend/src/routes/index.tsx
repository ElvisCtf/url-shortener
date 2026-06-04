import { Button } from '#/components/ui/button'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '#/components/ui/card'
import {
  InputGroup,
  InputGroupAddon,
  InputGroupInput,
} from '@/components/ui/input-group'
import { Globe, Link2 } from 'lucide-react'
import { createFileRoute } from '@tanstack/react-router'
import { Separator } from '@/components/ui/separator'
import { useState } from 'react'

export const Route = createFileRoute('/')({ component: Home })

function Home() {
  const [url, setUrl] = useState('')

  return (
    <main className="flex flex-col gap-8 items-stretch justify-start min-h-screen px-6 py-8 w-full max-w-2xl mx-auto">
      <section>
        <div className="flex items-center gap-4 mb-4">
          <div className="flex h-10 w-10 items-center justify-center rounded-full bg-primary">
            <Link2 className="h-5 w-5" />
          </div>
          <h1 className="text-3xl font-bold">SnapLink</h1>
        </div>
        <Separator />
      </section>

      <Card>
        <CardHeader>
          <CardTitle>Shorten a loooong URL</CardTitle>
          <CardDescription>
            Paste your destination link below to instantly generate a branded
            mini-link.
          </CardDescription>
        </CardHeader>

        <CardContent className="flex flex-col sm:flex-row items-center gap-2">
          <InputGroup className="py-6">
            <InputGroupInput placeholder="https://example.com/very-long-destination-url-path" />
            <InputGroupAddon>
              <Globe />
            </InputGroupAddon>
          </InputGroup>

          <Button type="submit" className="px-6 py-6 w-full sm:w-auto">
            Shorten →
          </Button>
        </CardContent>
      </Card>

      <div className="flex items-center justify-center rounded-2xl border border border-border bg-secondary w-full min-h-[100px]">
        <span className="text-sm text-muted-foreground">
          No links created yet. Paste a link above to get started!
        </span>
      </div>
    </main>
  )
}
