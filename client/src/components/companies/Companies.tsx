import Loader from '../loader/Loader'

import { useGetCompanies } from '../../hooks/apiHooks'

const Companies: React.FC = () => {
    const { data, isLoading } = useGetCompanies()


    return (
        <>
            <Loader isLoading={true}>
                {data.map(company => (
                    <div key={company.id}>{company.name}</div>
                ))}
            </Loader>
        </>
    )
}

export default Companies
