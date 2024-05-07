import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import { setInitialMovements, setError, setActualCompany, addMovement, setInitialProducts, addProduct, updateProduct, setTotalUnits, updateMovement, setCompanies, addCompany, updateCompany } from '../redux/slice'
import { ApiError } from './apiError'
import { companyToDomain, movementToCreateRequest, movementToUpdateRequest, productMovementToMovementTable, productToDomain, updateResponseToMovement } from '../../utils/mapper'

export const api = createApi({
    baseQuery: fetchBaseQuery({ baseUrl: 'http://localhost:8080/', timeout: 1000 }),
    endpoints: builder => ({
        getCompanies: builder.query({
            query: () => ({ url: '/company/' }),
            transformResponse: (response: any[]) => response?.map(companyToDomain) || [],
            async onQueryStarted({}, { dispatch, queryFulfilled }) {
                try {
                    const result = await queryFulfilled
                    dispatch(setCompanies(result.data))
                } catch (err: any) {
                    dispatch(setError(err.error as ApiError))
                }
            }
        }),
        addNewCompany: builder.mutation({
            query: (body) => ({
                url: `/company/`,
                method: 'POST',
                body: body
            }),
            transformResponse: companyToDomain,
            async onQueryStarted({}, {dispatch, queryFulfilled}) {
                try {
                    const result = await queryFulfilled
                    dispatch(addCompany(result.data))
                } catch (err: any) {
                    dispatch(setError(err.error as ApiError))
                }
            }
        }),
        updateCompany: builder.mutation({
            query: ({companyId, body}) => ({
                url: `/company/${companyId}`,
                method: 'PUT',
                body: body
            }),
            transformResponse: companyToDomain,
            async onQueryStarted({}, {dispatch, queryFulfilled}) {
                try {
                    const result = await queryFulfilled
                    dispatch(updateCompany(result.data))
                } catch (err: any) {
                    dispatch(setError(err.error as ApiError))
                }
            }
        }),
        getProductsMovementsFiltered: builder.mutation({
            query: ({ companyId, filters }) => ({
                url: `/company/${companyId}/movements?page_size=50&name=${filters.name}&brand=${filters.brand}&code=${filters.code}`
            }),
            transformResponse: (response: any) => ({
              companyName: response["company_name"],
              totalUnits: response["total_units"],
              movements: response.movements?.map(productMovementToMovementTable) || []
            }),
            async onQueryStarted({}, {dispatch, queryFulfilled}) {
                try {
                    const result = await queryFulfilled
                    dispatch(setActualCompany(result.data.companyName))
                    dispatch(setInitialMovements(result.data.movements))
                    dispatch(setTotalUnits(result.data.totalUnits))
                } catch (err: any) {
                    dispatch(setError(err.error as ApiError))
                }
            }
        }),
        addNewMovement: builder.mutation({
            query: ({ companyId, newMovement }) => ({
                url: `/company/${companyId}/movements`,
                method: 'POST',
                body: movementToCreateRequest(newMovement)
            }),
            transformResponse: productMovementToMovementTable,
            async onQueryStarted({}, {dispatch, queryFulfilled}) {
                try {
                    const result = await queryFulfilled
                    dispatch(addMovement(result.data))
                } catch (err: any) {
                    dispatch(setError(err.error as ApiError))
                }
            }
        }),
        updateMovement: builder.mutation({
            query: ({companyId, movementId, body}) => ({
                url: `company/${companyId}/movements/${movementId}`,
                method: 'PUT',
                body: movementToUpdateRequest(body)
            }),
            transformResponse: updateResponseToMovement,
            async onQueryStarted({}, {dispatch, queryFulfilled}) {
                try {
                    const result = await queryFulfilled
                    dispatch(updateMovement(result.data))
                } catch (err: any) {
                    dispatch(setError(err.error as ApiError))
                }
            }
        }),
        getProductsByCompanyId: builder.query({
            query: (companyId: number) => ({url: `/company/${companyId}/products`}),
            transformResponse: (response: any[]) => response?.map(productToDomain) || [],
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
            transformResponse: productToDomain,
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
