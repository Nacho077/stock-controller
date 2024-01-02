export interface ApiError {
    data: string
    status: number
}

export interface Company {
    readonly id: number
    name: string
}

// Interfaces for modal
export interface Button {
    text: string
    onClick?: VoidFunction
}

export interface ModalData {
    text: string
    buttons: Button[]
}

export interface ModalInfo {
    data: ModalData
    onClose: VoidFunction
}
