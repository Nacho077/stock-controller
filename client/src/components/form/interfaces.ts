import { ChangeEvent, FormEvent } from "react"

export interface Button {
    title: string
    type?: "submit" | "reset" | "button"
    onClick?: any
}

export interface FormFields {
    [key: string]: string | string[] | boolean
    label: string
    type: "text" | "date" | "number"
    name: string
    datalist: string[]
}

export interface FormProps<T> {
    title?: string
    fields: FormFields[]
    buttons: Button[]
    formValues: T
    handleChange: (e: ChangeEvent<HTMLInputElement>) => void
    onSubmit: (e: FormEvent) => void
    onReset?:() => void
    refIndex?: number
}