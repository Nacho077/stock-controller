import { configureStore } from '@reduxjs/toolkit'
import { api } from '../api/apiconfig'
import reducer from './slice'

const store = configureStore({
    reducer: {
        [api.reducerPath]: api.reducer,
        reducer
    },
    middleware: getDefaultMiddleware => getDefaultMiddleware().concat(api.middleware),
})

export type RootState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch

export default store
