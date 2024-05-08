import { api } from '../configs/api/apiconfig.ts'
import { Company } from '../views/companies/interfaces.ts'
import { ProductFilters, ProductMovement } from '../views/movements/interfaces.ts'
import { Product } from '../views/products/interfaces.ts'

const {
    useGetCompaniesQuery,
    useAddNewCompanyMutation,
    useUpdateCompanyMutation,
    useGetProductsMovementsFilteredMutation,
    useAddNewMovementMutation,
    useUpdateMovementMutation,
    useGetProductsByCompanyIdQuery,
    useAddNewProductMutation,
    useUpdateProductMutation,
} = api

interface UseGetCompaniesResult {
    data: Company[]
    isLoading: boolean
}

export const useGetCompanies: () => UseGetCompaniesResult = () => {
    const result = useGetCompaniesQuery({})

    return {
        data: result.data || [],
        isLoading: result.isLoading
    }
}

export const useCreateNewCompany = () => {
    const [mutate, { }] = useAddNewCompanyMutation()

    return (company: Company) => mutate(company)
}

export const useUpdateCompany = () => {
    const [mutate, { }] = useUpdateCompanyMutation()

    return (companyId: number, body: Company) => mutate({ companyId, body })
}

export const useGetProductsMovementsFiltered = () => {
    const [mutate, { }] = useGetProductsMovementsFilteredMutation()

    return (companyId: number, filters: ProductFilters) => mutate({ companyId, filters })
}

export const useCreateNewMovement = () => {
    const [mutate, { }] = useAddNewMovementMutation()

    return (companyId: number, newMovement: ProductMovement) => mutate({ companyId, newMovement })
}

export const useUpdateMovement = () => {
    const [mutate, { }] = useUpdateMovementMutation()

    return (companyId: number, movementId: number, body: ProductMovement) => mutate({ companyId, movementId, body })
}

interface UseGetProducts {
    isLoading: boolean,
    refetch: () => void
}

type UseGetProductsByCompanyId = (id: number) => UseGetProducts

export const useGetProductsByCompanyId: UseGetProductsByCompanyId = (id: number) => {
    const result = useGetProductsByCompanyIdQuery(id)

    return {
        isLoading: result.isLoading,
        refetch: result.refetch
    }
}

export const useCreateNewProduct = () => {
    const [mutate, { }] = useAddNewProductMutation()

    return (body: Product) => mutate(body)
}

export const useUpdateProduct = () => {
    const [mutate, { }] = useUpdateProductMutation()

    return (productId: number, body: Product) => mutate({ productId, body })
}
