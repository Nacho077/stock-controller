import { api } from '../configs/api/apiconfig.ts'
import { ProductFilters, ProductMovement } from '../views/movements/interfaces.ts'
import { Product } from '../views/products/interfaces.ts'

const {
    useGetCompaniesQuery,
    useGetProductsMovementsFilteredMutation,
    useAddNewMovementMutation,
    useUpdateMovementMutation,
    useGetProductsByCompanyIdQuery,
    useAddNewProductMutation,
    useUpdateProductMutation,
} = api

interface Company {
    readonly id: number
    name: string
}

interface UseGetCompaniesResult {
    data: Company[]
    isLoading: boolean
}

export const useGetCompanies: () => UseGetCompaniesResult = () => {
    const result = useGetCompaniesQuery({})
    const defaultCompany = {
        id: "create",
        name: "Crea una nueva empresa"
    }

    return {
        data: [defaultCompany, ...result.data || []],
        isLoading: result.isLoading
    }
}

export const useGetProductsMovementsFiltered = () => {
    const [mutate, {}] = useGetProductsMovementsFilteredMutation()

    return (companyId: number, filters: ProductFilters) => mutate({companyId, filters})
}

export const useCreateNewMovement = () => {
    const [mutate, {}] = useAddNewMovementMutation()

    return (companyId: number, newMovement: ProductMovement) => mutate({companyId, newMovement})
}

export const useUpdateMovement = () => {
    const [mutate, {}] = useUpdateMovementMutation()

    return (companyId: number, movementId: number, body: ProductMovement) => mutate({companyId, movementId, body})
}

interface UseGetProducts {
    isLoading: boolean
}

type UseGetProductsByCompanyId = (id: number) => UseGetProducts

export const useGetProductsByCompanyId: UseGetProductsByCompanyId = (id: number) => {
    const result = useGetProductsByCompanyIdQuery(id)

    return {
        isLoading: result.isLoading
    }
}

export const useCreateNewProduct = () => {
    const [mutate, {}] = useAddNewProductMutation()

    return (body: Product) => mutate(body)
}

export const useUpdateProduct = () => {
    const [mutate, {}] = useUpdateProductMutation()

    return (productId: number, body: Product) => mutate({productId, body})
}
