import { PayloadAction, createSlice } from '@reduxjs/toolkit'
import { ApiError } from '../api/apiError.ts'
import { ProductMovement } from '../../views/movements/interfaces.ts'

interface InitialState {
    error: string
    movements: ProductMovement[],
}

const initialState: InitialState = {
    error: '',
    movements: []
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
        setInitialMovements: (state, action: PayloadAction<ProductMovement[]>) => {
            state.movements = action.payload
        },
        addMovement: (state, action: PayloadAction<ProductMovement>) => {
            state.movements.unshift(action.payload)
        }
    }
})

export const { setError, deleteError, setInitialMovements, addMovement } = slice.actions

export default slice.reducer
