import { api } from '../configs/api/apiconfig.ts'

const {
    useGetCompaniesQuery,
    useGetProductsMovementsByCompanyIdQuery,
    useAddNewMovementMutation
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

type CreateNewMovement = () => any

export const useCreateNewMovement: CreateNewMovement = () => {
    const [mutate, { }] = useAddNewMovementMutation();

    return mutate
}
