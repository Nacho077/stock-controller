import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import { getCompanies, setError } from '../redux/slice'
import { ApiError } from '../../interfaces/interfaces'

export const api = createApi({
    baseQuery: fetchBaseQuery({baseUrl: 'http://localhost:8080/', timeout: 1000}),
    endpoints: builder => ({
        getCompanies: builder.query({
            // method, body, headers, params(reibe un obj)
            query: () => ({url: 'company/'}),
            async onQueryStarted({}, {dispatch, queryFulfilled}) {
                try {
                    const result = await queryFulfilled
                    dispatch(getCompanies(result.data))
                } catch (err: any) {
                    dispatch(setError(err.error as ApiError))
                }
            }
        })
    })
})
