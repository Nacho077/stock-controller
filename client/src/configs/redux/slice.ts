import { PayloadAction, createSlice } from '@reduxjs/toolkit'
import { ApiError, Company } from '../../interfaces/interfaces.ts'

interface InitialState {
    error: string
    companies: Company[]
}

const initialState: InitialState = {
    error: '',
    companies: []
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
        getCompanies: (state, action: PayloadAction<Company[]>) => {
            state.companies.push(...action.payload)
        }
    }
})

export const { getCompanies, setError, deleteError } = slice.actions

export default slice.reducer
