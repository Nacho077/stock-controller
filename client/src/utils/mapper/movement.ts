import { ProductMovement } from "../../views/movements/interfaces"

const productMovementToMovementTable = (movement: any): ProductMovement => {
    return {
        id: movement.movement_id,
        movementId: movement.movement_id,
        date: movement.date,
        shippingCode: movement.shipping_code,
        units: movement.units,
        deposit: movement.deposit,
        observations: movement.observations,
        productId: movement.product_id,
        name: movement.name,
        code: movement.code,
        brand: movement.brand,
        detail: movement.detail,
        companyId: movement.company_id
    }
}

export const productMovementArrToMovementTable = (movements: any[]): ProductMovement[] => {
    return movements.map(m => productMovementToMovementTable(m))
}

export const movementToRequest = ({ date, shippingCode, units, deposit, observations, productId }: ProductMovement) => ({
    movement: {
        date,
        shipping_code: shippingCode,
        inits: parseInt(units.toString()),
        deposit,
        observations
    },
    product_id: productId
})