import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import { setInitialMovements, setError, setActualCompany, addMovement, setInitialProducts, addProduct, updateProduct } from '../redux/slice'
import { ApiError } from './apiError'
import { productMovementToMovementTable, productToDomain } from '../../utils/mapper'
import { ProductMovement } from '../../views/movements/interfaces'

export const api = createApi({
    baseQuery: fetchBaseQuery({baseUrl: 'http://localhost:8080/', timeout: 1000}),
    endpoints: builder => ({
        getCompanies: builder.query({
            query: () => ({url: 'company/'}),
            async onQueryStarted({}, {dispatch, queryFulfilled}) {
                try {
                    await queryFulfilled
                } catch (err: any) {
                    dispatch(setError(err.error as ApiError))
                }
            }
        }),
        getProductsMovementsByCompanyId: builder.query({
            query: (companyId: number) => ({url: `/company/${companyId}/movements?page_size=50`}),
            transformResponse: (response: any) => ({
              companyName: response["company_name"],
              movements: response.movements.map((m: any) => productMovementToMovementTable(m))
            }),
            async onQueryStarted({}, {dispatch, queryFulfilled}) {
                try {
                    const result = await queryFulfilled
                    dispatch(setActualCompany(result.data.companyName))
                    dispatch(setInitialMovements(result.data.movements))
                } catch (err: any) {
                    dispatch(setError(err.error as ApiError))
                }
            }
        }),
        addNewMovement: builder.mutation({
            query: () => ({url: `/ping`}),
            async onQueryStarted(newMovement: ProductMovement, {dispatch, queryFulfilled}) {
                try {
                    await queryFulfilled
                    dispatch(addMovement(newMovement))
                } catch (err: any) {
                    dispatch(setError(err.error as ApiError))
                }
            }
        }),
        getProductsByCompanyId: builder.query({
            query: (companyId: number) => ({url: `/company/${companyId}/products`}),
            transformResponse: (response: any[]) => response.map((r: any) => productToDomain(r)),
            async onQueryStarted({}, {dispatch, queryFulfilled}) {
                try {
                    const result = await queryFulfilled
                    dispatch(setInitialProducts(result.data))
                } catch (err: any) {
                    dispatch(setError(err.error as ApiError))
                }
            }
        }),
        addNewProduct: builder.mutation({
            query: (body) => ({
                url: `/product/`,
                method: 'POST',
                body
            }),
            transformResponse: (response: any) => productToDomain(response),
            async onQueryStarted({}, {dispatch, queryFulfilled}) {
                try {
                    const result = await queryFulfilled
                    dispatch(addProduct(result.data))
                } catch (err: any) {
                    dispatch(setError(err.error as ApiError))
                }
            }
        }),
        updateProduct: builder.mutation({
            query: ({productId, body}) => ({
                url: `/product/${productId}`,
                method: 'PUT',
                body
            }),
            transformResponse: (response: any) => productToDomain(response),
            async onQueryStarted({}, {dispatch, queryFulfilled}) {
                try {
                    const result = await queryFulfilled
                    dispatch(updateProduct(result.data))
                } catch (err: any) {
                    dispatch(setError(err.error as ApiError))
                }
            }
        })
    })
})
