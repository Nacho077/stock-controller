import { getDateNowString } from '../../utils/date';
import { FormFields } from '../../components/form/interfaces';
import { Header, Row } from '../../components/table/interfaces';
import { Product } from '../products/interfaces';

export interface ProductMovement extends Row {
    movementId: number
    date: string
    shippingCode: string
    units: number
    deposit: string
    observations: string
    productId: number
    name: string
    code: string
    brand: string
    detail: string
    companyId: number
}

export interface ProductFilters {
    code: string
    name: string
    brand: string
}

const removeDuplicates = (arr: any[]) => [...new Set(arr)]

export const movementsFiltersFields =  (products: Product[]): FormFields[] => [
    {
        label: "codigo",
        type: "text",
        name: "code",
        datalist: removeDuplicates(products.map(p => p.code.trim().toLowerCase())),
        autoComplete: "off"
    },
    {
        label: "aparato",
        type: "text",
        name: "name",
        datalist: removeDuplicates(products.map(p => p.name.trim().toLowerCase())),
        autoComplete: "off"
    },
    {
        label: "marca",
        type: "text",
        name: "brand",
        datalist: removeDuplicates(products.map(p => p.brand.trim().toLowerCase())),
        autoComplete: "off"
    },
]

export const movementHeaders: Header[] = [
    { key: "date", value: "fecha" },
    { key: "shippingCode", value: "remito" },
    { key: "units", value: "unidades" },
    { key: "code", value: "codigo" },
    { key: "name", value: "aparato" },
    { key: "brand", value: "marca" },
    { key: "detail", value: "detalle" },
    { key: "deposit", value: "deposito" },
    { key: "observations", value: "observación" }
]

export const movementFormFields = (isUpdate: boolean, products: Product[]): FormFields[] => [
    {
        label: "fecha",
        type: "date",
        name: "date",
        datalist: [],
        required: true,
        autoComplete: "off",
    },
    {
        label: "remito",
        type: "text",
        name: "shippingCode",
        datalist: [],
        required: true,
        autoComplete: "off"
    },
    {
        label: "codigo",
        type: "text",
        name: "code",
        datalist: removeDuplicates(products.map(p => p.code.trim().toLowerCase())),
        required: true,
        autoComplete: "off",
        disabled: isUpdate,
    },
    {
        label: "unidades",
        type: "number",
        name: "units",
        datalist: [],
        autoComplete: "off"
    },
    {
        label: "aparato",
        type: "text",
        name: "name",
        datalist: removeDuplicates(products.map(p => p.name.trim().toLowerCase())),
        autoComplete: "off",
        disabled: isUpdate,
    },
    {
        label: "marca",
        type: "text",
        name: "brand",
        datalist: removeDuplicates(products.map(p => p.brand.trim().toLowerCase())),
        autoComplete: "off",
        disabled: isUpdate,
    },
    {
        label: "detalle",
        type: "text",
        name: "detail",
        datalist: [],
        autoComplete: "off",
        disabled: isUpdate,
    },
    {
        label: "deposito",
        type: "text",
        name: "deposit",
        datalist: [],
        autoComplete: "off"
    },
    {
        label: "observación",
        name: "observations",
        datalist: [],
        type: "text",
        autoComplete: "off"
    }
]

export const getDefaultFilters = () => ({
    code: "",
    name: "",
    brand: "",
})

export const getDefaultMovement = (
    companyId: number,
    date = "",
    shippingCode = "",
    code = "",
    productId = 0
): ProductMovement => ({
    id: 0,
    movementId: 0,
    date: date || getDateNowString(),
    units: 0,
    deposit: "",
    observations: "",
    productId: productId,
    name: "",
    brand: "",
    detail: "",
    shippingCode,
    code,
    companyId
})