import { PayloadAction, createSlice } from '@reduxjs/toolkit'
import { ApiError } from '../api/apiError.ts'
import { ProductMovement } from '../../views/movements/interfaces.ts'
import { Product } from '../../views/products/interfaces.ts'

interface InitialState {
    error: string
    movements: ProductMovement[],
    actualCompany: string,
    products: Product[]
}

const initialState: InitialState = {
    error: '',
    movements: [],
    actualCompany: 'Movimientos',
    products: []
}

const slice = createSlice({
    name: 'slice',
    initialState,
    reducers: {
        setError: (state, action: PayloadAction<ApiError>) => {
            console.log(action.payload)
            state.error = `${action.payload.status} - ${action.payload.data ?? action.payload.error}`
        },
        deleteError: (state) => {
            state.error = ''
        },
        setActualCompany: (state, action: PayloadAction<string>) => {
          state.actualCompany = action.payload
        },
        setInitialMovements: (state, action: PayloadAction<ProductMovement[]>) => {
            state.movements = action.payload
        },
        addMovement: (state, action: PayloadAction<ProductMovement>) => {
            state.movements.unshift(action.payload)
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
        }
    }
})

export const { setError, deleteError, setActualCompany, setInitialMovements, addMovement, setInitialProducts, addProduct, updateProduct } = slice.actions

export default slice.reducer
