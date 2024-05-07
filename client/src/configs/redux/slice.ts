import { PayloadAction, createSlice } from '@reduxjs/toolkit'
import { ApiError } from '../api/apiError.ts'
import { ProductMovement } from '../../views/movements/interfaces.ts'
import { Product } from '../../views/products/interfaces.ts'

interface InitialState {
    error: string
    movements: ProductMovement[],
    actualCompany: string,
    products: Product[],
    totalUnits: number
}

const initialState: InitialState = {
    error: '',
    movements: [],
    actualCompany: 'Movimientos',
    products: [],
    totalUnits: 0
}

const slice = createSlice({
    name: 'slice',
    initialState,
    reducers: {
        setError: (state, action: PayloadAction<ApiError>) => {
            console.log(action.payload)
            state.error = `${action.payload.status} - ${action.payload.data ?? action.payload.error}`
        },
        clearError: (state) => {
            state.error = ''
        },
        setActualCompany: (state, action: PayloadAction<string>) => {
          state.actualCompany = action.payload
        },
        setInitialMovements: (state, action: PayloadAction<ProductMovement[]>) => {
            state.movements = action.payload
        },
        setTotalUnits: (state, action: PayloadAction<number>) => {
            state.totalUnits = action.payload
        },
        addMovement: (state, action: PayloadAction<ProductMovement>) => {
            state.movements.unshift(action.payload)
        },
        updateMovement: (state, action: PayloadAction<ProductMovement>) => {
            for (let i = 0; i < state.movements.length; i++) {
                if (state.movements[i].id == action.payload.id) {
                    state.movements[i].date = action.payload.date
                    state.movements[i].shippingCode = action.payload.shippingCode
                    state.movements[i].units = action.payload.units
                    state.movements[i].deposit = action.payload.deposit
                    state.movements[i].observations = action.payload.observations
                    break
                }
            }
        },
        setInitialProducts: (state, action: PayloadAction<Product[]>) => {
            state.products = action.payload
        },
        addProduct: (state, action: PayloadAction<Product>) => {
            state.products.unshift(action.payload)
        },
        updateProduct: (state, action: PayloadAction<Product>) => {
            for (let i = 0; i < state.products.length; i++) {
                if (state.products[i].id == action.payload.id) {
                    state.products[i] = action.payload
                    break
                }
            }
        },
    }
})

export const { setError, clearError, setActualCompany, setInitialMovements, setTotalUnits, addMovement, updateMovement, setInitialProducts, addProduct, updateProduct} = slice.actions

export default slice.reducer
