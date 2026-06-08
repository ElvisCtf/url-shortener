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
import { Separator } from '@/components/ui/separator'
import { useEffect, useState } from 'react'
import { columns, type LinkRecord } from './columns'
import { DataTable } from './data-table'

interface ShortenedListProps {
  shortenList: LinkRecord[]
  setShortenList: React.Dispatch<React.SetStateAction<LinkRecord[]>>
}

interface ShortenFormProps {
  url: string
  setUrl: React.Dispatch<React.SetStateAction<string>>
}

export default function ShortenPage() {
  const [url, setUrl] = useState('')
  // todo: replace with actual data from backend
  const [shortenList, setShortenList] = useState<LinkRecord[]>([])

  useEffect(() => {
    setShortenList([
      {
        id: '1',
        originalUrl: 'https://example.com/very-long-destination-url-path',
        shortUrl: 'https://snap.link/abc123',
      },
    ])
  }, [])

  return (
    <main className="flex flex-col gap-8 items-stretch justify-start min-h-screen px-6 py-8 w-full max-w-2xl mx-auto">
      <Header />
      <ShortenForm url={url} setUrl={setUrl} />
      <ShortenedList
        shortenList={shortenList}
        setShortenList={setShortenList}
      />
    </main>
  )
}

function Header() {
  return (
    <section>
      <div className="flex items-center gap-4 mb-4">
        <div className="flex h-10 w-10 items-center justify-center rounded-full bg-primary">
          <Link2 className="h-5 w-5" />
        </div>
        <h1 className="text-3xl font-bold">SnapLink</h1>
      </div>
      <Separator />
    </section>
  )
}

function ShortenForm({ url, setUrl }: ShortenFormProps) {
  const submit = () => {
    // todo: call backend to shorten the URL and update the shortenedList state
  }

  return (
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
          <InputGroupInput
            placeholder="https://example.com/very-long-destination-url-path"
            value={url}
            onChange={(e) => setUrl(e.target.value)}
          />
          <InputGroupAddon>
            <Globe />
          </InputGroupAddon>
        </InputGroup>

        <Button
          type="button"
          className="px-6 py-6 w-full sm:w-auto"
          onClick={submit}
        >
          Shorten →
        </Button>
      </CardContent>
    </Card>
  )
}

function ShortenedList({ shortenList, setShortenList }: ShortenedListProps) {
  const clearAll = () => {
    // todo: clear local storage and update state
    setShortenList([])
  }

  if (shortenList.length === 0) {
    return (
      <div className="flex items-center justify-center rounded-2xl border border-border bg-secondary w-full min-h-[100px]">
        <span className="text-sm text-muted-foreground">
          No links created yet. Paste a link above to get started!
        </span>
      </div>
    )
  } else {
    return (
      <DataTable columns={columns} data={shortenList} onClearAll={clearAll} />
    )
  }
}
