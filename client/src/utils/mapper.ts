import { Company } from "../views/companies/interfaces"
import { ProductMovement } from "../views/movements/interfaces"
import { Product } from "../views/products/interfaces"

export const companyToDomain = (company: any): Company => {
    return {
        id: company.id,
        name: company.name
    }
}

export const productMovementToMovementTable = (movement: any): ProductMovement => {
    return {
        id: movement.movement_id,
        movementId: movement.movement_id,
        date: movement.date,
        shippingCode: movement.shipping_code || "",
        units: movement.units || 0,
        deposit: movement.deposit || "",
        observations: movement.observations || "",
        productId: movement.product_id,
        name: movement.name || "",
        code: movement.code || "",
        brand: movement.brand || "",
        detail: movement.detail || "",
        companyId: movement.company_id
    }
}

export const productToDomain = (product: any): Product => {
    return {
        id: product.id,
        name: product.name || "",
        code: product.code,
        brand: product.brand || "",
        detail: product.detail || "",
        company_id: 0
    }
}

export const movementToCreateRequest = ({ date, shippingCode, units, deposit, observations, productId }: ProductMovement) => ({
    movement: {
        date,
        shipping_code: shippingCode,
        units: parseInt(units.toString()),
        deposit,
        observations
    },
    product_id: productId
})

export const movementToUpdateRequest = ({date, shippingCode, units, deposit, observations}: ProductMovement) => ({
    date,
    shipping_code: shippingCode,
    units: parseInt(units.toString()),
    deposit,
    observations
})

export const updateResponseToMovement = ({id, date, shipping_code, units, deposit, observations}: any): ProductMovement => ({
    id,
    date,
    shippingCode: shipping_code,
    units,
    deposit,
    observations,
    movementId: 0,
    productId: 0,
    name: "",
    code: "",
    brand: "",
    detail: "",
    companyId: 0,
})