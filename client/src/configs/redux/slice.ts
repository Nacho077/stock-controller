import { createSlice } from '@reduxjs/toolkit'

const initialState = {
    loading: false
}

const slice = createSlice({
    name: 'slice',
    initialState,
    reducers: {
        loadingStarted: state => {
            state.loading = true
        },
        loadingFinished: state => {
            state.loading = false
        }
    }
})

export const { loadingStarted, loadingFinished } = slice.actions

export default slice.reducer
