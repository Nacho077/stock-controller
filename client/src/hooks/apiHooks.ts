import { api } from '../configs/api/apiconfig.ts'
import { Company } from '../interfaces/interfaces.ts'

const { useGetCompaniesQuery } = api

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
