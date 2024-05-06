import { FormFields } from "../../components/form/interfaces";
import { Header, Row } from "../../components/table/interfaces";

export interface Product extends Row {
    name: string,
    code: string,
    brand: string,
    detail: string,
    company_id: number
}

export const productsHeaders: Header[] = [
    {key: "code", value: "codigo"},
    {key: "name", value: "nombre"},
    {key: "brand", value: "marca"},
    {key: "detail", value: "detalle"}
]

export const getDefaultProduct = (companyId: number): Product => {
    return {
        id: 0,
        name: "",
        code: "",
        brand: "",
        detail: "",
        company_id: companyId
    }
}

export const getProduct = (product: Product, companyId: number): Product => {
    return {
        ...product,
        company_id: companyId
    }
}

export const productFormFields: FormFields[] = [
    {
        label: "nombre",
        type: "text",
        name: "name",
        datalist: [],
        required: true,
        autoComplete: "off"
    },
    {
        label: "codigo",
        type: "text",
        name: "code",
        datalist: [],
        required: true,
        autoComplete: "off"
    },
    {
        label: "marca",
        type: "text",
        name: "brand",
        datalist: [],
        required: false,
        autoComplete: "off"
    },
    {
        label: "detalle",
        type: "text",
        name: "detail",
        datalist: [],
        required: false,
        autoComplete: "off"
    }
]