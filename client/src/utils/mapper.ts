import { ProductMovement } from "../views/movements/interfaces"
import { Product } from "../views/products/interfaces"

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