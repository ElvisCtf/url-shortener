import type { ColumnDef } from '@tanstack/react-table'
import { Copy, QrCode } from 'lucide-react'
import { Button } from '@/components/ui/button'

export type LinkRecord = {
  id: string
  originalUrl: string
  shortUrl: string
}

export const columns: ColumnDef<LinkRecord>[] = [
  {
    accessorKey: 'originalUrl',
    header: 'Original URL',
    cell: ({ row }) => (
      <div
        className="max-w-[300px] truncate"
        title={row.getValue('originalUrl')}
      >
        {row.getValue('originalUrl')}
      </div>
    ),
  },
  {
    accessorKey: 'shortUrl',
    header: 'Short URL',
    cell: ({ row }) => (
      <a
        href={row.getValue('shortUrl')}
        target="_blank"
        rel="noopener noreferrer"
        className="text-blue-600 hover:underline dark:text-blue-400"
      >
        {row.getValue('shortUrl')}
      </a>
    ),
  },
  {
    id: 'actions',
    header: () => <div className="text-right">Actions</div>,
    cell: ({ row }) => {
      const record = row.original

      const handleCopy = async () => {
        await navigator.clipboard.writeText(record.shortUrl)
      }

      const handleQrClick = () => {
        console.log('Generate QR for:', record.shortUrl)
      }

      return (
        <div className="flex items-center justify-end gap-2">
          <Button
            variant="ghost"
            size="icon"
            className="h-8 w-8 text-muted-foreground hover:text-foreground"
            onClick={handleCopy}
            title="Copy Link"
          >
            <Copy className="h-4 w-4" />
          </Button>
          <Button
            variant="ghost"
            size="icon"
            className="h-8 w-8 text-muted-foreground hover:text-foreground"
            onClick={handleQrClick}
            title="View QR Code"
          >
            <QrCode className="h-4 w-4" />
          </Button>
        </div>
      )
    },
  },
]
