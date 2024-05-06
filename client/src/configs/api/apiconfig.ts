import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import { setInitialMovements, setError, setActualCompany, addMovement } from '../redux/slice'
import { ApiError } from './apiError'
import { movementToRequest, productMovementArrToMovementTable } from '../../utils/mapper/movement'

export const api = createApi({
    baseQuery: fetchBaseQuery({ baseUrl: 'http://localhost:8080/', timeout: 1000 }),
    endpoints: builder => ({
        getCompanies: builder.query({
            query: () => ({ url: 'company/' }),
            async onQueryStarted({ }, { dispatch, queryFulfilled }) {
                try {
                    await queryFulfilled
                } catch (err: any) {
                    dispatch(setError(err.error as ApiError))
                }
            }
        }),
        getProductsMovementsByCompanyId: builder.query({
            query: (companyId: number) => ({ url: `/company/${companyId}/movements?page_size=20` }),
            async onQueryStarted({ }, { dispatch, queryFulfilled }) {
                try {
                    const result = await queryFulfilled
                    const mappedData = productMovementArrToMovementTable(result.data.movements)
                    dispatch(setActualCompany(result.data["company_name"]))
                    dispatch(setInitialMovements(mappedData))
                } catch (err: any) {
                    dispatch(setError(err.error as ApiError))
                }
            }
        }),
        addNewMovement: builder.mutation({
            query: ({ companyId, newMovement }) => {
                return {
                    url: `/company/${companyId}/movements`,
                    method: 'POST',
                    body: movementToRequest(newMovement)
                }
            },
            async onQueryStarted({ }, { dispatch, queryFulfilled }) {
                try {
                    const result = await queryFulfilled
                    dispatch(addMovement(result.data))
                } catch (err: any) {
                    dispatch(setError(err.error as ApiError))
                }
            }
        })
    })
})
