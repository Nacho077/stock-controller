import { PayloadAction, createSlice } from '@reduxjs/toolkit'
import { ApiError } from '../api/apiError.ts'
import { ProductMovement } from '../../views/movements/interfaces.ts'

interface InitialState {
    error: string
    movements: ProductMovement[],
    actualCompany: string,
}

const initialState: InitialState = {
    error: '',
    movements: [],
    actualCompany: 'Movements'
}

const slice = createSlice({
    name: 'slice',
    initialState,
    reducers: {
        setError: (state, action: PayloadAction<ApiError>) => {
            console.log(action.payload)
            state.error = `${action.payload.status} - ${action.payload.data}`
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
        }
    }
})

export const { setError, deleteError, setActualCompany, setInitialMovements, addMovement } = slice.actions

export default slice.reducer
