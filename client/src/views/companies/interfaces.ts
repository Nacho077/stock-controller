import { Button, FormFields } from "../../components/form/interfaces"

export interface Company {
    [key: string]: string | number
    readonly id: number
    name: string
}

export const formFields: FormFields[] = [
    {
        label: "nombre",
        type: "text",
        name: "name",
        datalist: [],
        autoComplete: "off"
    }
]

export const formButtons: Button[] = [
    {
        title: "Cancelar",
        type: "reset"
    },
    {
        title: "Aceptar",
        type: "submit"
    }
]

export const defaultCompany: Company = {
    id: 0,
    name: ""
}
