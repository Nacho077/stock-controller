import { PayloadAction, createSlice } from '@reduxjs/toolkit'
import { ApiError } from '../api/apiError.ts'
import { ProductMovement } from '../../views/movements/interfaces.ts'
import { Product } from '../../views/products/interfaces.ts'
import { Company } from '../../views/companies/interfaces.ts'

interface InitialState {
    error: string,
    companies: Company[],
    movements: ProductMovement[],
    actualCompany: string,
    products: Product[],
    totalUnits: number
}

const initialState: InitialState = {
    error: '',
    companies: [],
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
            console.error(action.payload)
            state.error = `${action.payload.status ?? action.payload.name} - ${action.payload.data ?? action.payload.error ?? action.payload.message}`
        },
        clearError: (state) => {
            state.error = ''
        },
        setCompanies: (state, action: PayloadAction<Company[]>) => {
            state.companies = action.payload
        },
        addCompany: (state, action: PayloadAction<Company>) => {
            state.companies.unshift(action.payload)
        },
        updateCompany: (state, action: PayloadAction<Company>) => {
            for (let i = 0; i < state.companies.length; i++) {
                if (state.companies[i].id == action.payload.id) {
                    state.companies[i] = action.payload
                    break
                }
            }
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
        incrementTotalUnits: (state, action: PayloadAction<number>) => {
            state.totalUnits = state.totalUnits + action.payload
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

export const { setError, clearError, setCompanies, addCompany, updateCompany, setActualCompany, setInitialMovements, setTotalUnits, incrementTotalUnits, addMovement, updateMovement, setInitialProducts, addProduct, updateProduct } = slice.actions

export default slice.reducer
