import { Row, SelectableTableProps } from './interfaces'

import styles from './table.module.scss'

const SelectableTable = <T extends Row>({
    headers,
    rows,
    handleDoubleClick = () => {},
    selectedRow = -1
}: SelectableTableProps<T>) => {
    // const EMPTYIMAGE = 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/wcAAwAB/epMzgAAAABJRU5ErkJggg=='

    // const getSelectedRows = (initialSelectedRowId:number, finalSelectedRowId: number) => {
    //     const initialIndex = rows.findIndex(row => row.id == initialSelectedRowId)
    //     const finalIndex = rows.findIndex(row => row.id == finalSelectedRowId)

    //     const startIndex = Math.min(initialIndex, finalIndex)
    //     const endIndex = Math.max(initialIndex, finalIndex)

    //     const newSelectedRows = rows.slice(startIndex, endIndex + 1)

    //     setSelectedRows(newSelectedRows.map((row) => row.id))
    // }

    // const handleDragStart = (e: React.DragEvent, id: number) => {
    //     setInitialSelectedRow(id)

    //     const dragImage = new Image()
    //     dragImage.src = EMPTYIMAGE
    //     e.dataTransfer.setDragImage(dragImage, 0, 0)
    // }

    // const handleDragOver = (e: React.DragEvent) => {
    //     e.preventDefault()
    // }

    // const handleDragEnter = (e: React.DragEvent, id: number) => {
    //     e.preventDefault()

    //     getSelectedRows(initialSelectedRow, id)
    // }

    // const handleDragLeave = (e: React.DragEvent) => {
    //     e.preventDefault()
    // }

    // const handleDrop = (e: React.DragEvent, id: number) => {
    //     e.preventDefault()
        
    //     getSelectedRows(initialSelectedRow, id)
    // }

    return (
        <table className={styles.containerTable}>
            <thead>
                <tr>
                    {headers.map(header => <th key={header.key}>{header.value}</th>)}
                </tr>
            </thead>
            <tbody>
                {rows.map(row => (
                    <tr
                        key={row.id}
                        // draggable
                        // onDragStart={(e) => handleDragStart(e, row.id)}
                        // onDragOver={(e) => handleDragOver(e)}
                        // onDragEnter={(e) => handleDragEnter(e, row.id)}
                        // onDragLeave={(e) => handleDragLeave(e)}
                        // onDrop={(e) => handleDrop(e, row.id)}
                        className={selectedRow == row.id ? styles.selected : ''}
                        onDoubleClick={() => {handleDoubleClick(row)}}
                    >
                        {headers.map(header => (
                            <td key={`${header.key}-${row.id}`}>
                                {row[header.key] || ''}
                            </td>
                        ))}
                    </tr>
                ))}
            </tbody>
        </table>
    )
}

export default SelectableTable
