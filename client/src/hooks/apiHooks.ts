import { api } from '../configs/api/apiconfig.ts'

const {
    useGetCompaniesQuery,
    useGetProductsMovementsByCompanyIdQuery,
    useAddNewMovementMutation,
    useGetProductsByCompanyIdQuery,
    useAddNewProductMutation,
    useUpdateProductMutation
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

interface UseGetProductMovements {
    isLoading: boolean
}

type UseGetProductsMovementsByCompanyId = (id: number) => UseGetProductMovements

export const useGetProductsMovementsByCompanyId: UseGetProductsMovementsByCompanyId = (id: number) => {
    const result = useGetProductsMovementsByCompanyIdQuery(id)

    return {
        isLoading: result.isLoading
    }
}

type Mutation = () => any

export const useCreateNewMovement: Mutation = () => {
    const [mutate, {}] = useAddNewMovementMutation();

    return mutate
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

export const useCreateNewProduct: Mutation = () => {
    const [mutate, {}] = useAddNewProductMutation()

    return mutate
}

export const useUpdateProduct: Mutation = () => {
    const [mutate, {}] = useUpdateProductMutation()

    return mutate
}
