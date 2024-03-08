export interface Header {
    key: string
    value: string
}

export interface Row {
    readonly id: number
    [key: string]: string | number
}

export interface SelectableTableProps<T extends Row> {
    headers: Header[]
    rows: T[]
    handleDoubleClick?: (obj: T) => void
    selectedRow?: number
}