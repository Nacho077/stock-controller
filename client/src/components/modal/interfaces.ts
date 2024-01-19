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