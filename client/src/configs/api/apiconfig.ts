import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import { loadingStarted, loadingFinished } from '../redux/slice'

export const api = createApi({
    baseQuery: fetchBaseQuery({baseUrl: 'http://localhost:8080', timeout: 1000}),
    endpoints: builder => ({
        ping: builder.query({
            // method, body, headers, params(reibe un obj)
            query: () => ({url: '/ping'}),
            async onQueryStarted({}, {dispatch, queryFulfilled}) {
                dispatch(loadingStarted())
                try {
                    await queryFulfilled
                    dispatch(loadingFinished())
                } catch (err) {
                    dispatch(loadingFinished())
                }
            }
        })
    }),
})

export const { usePingQuery } = api
