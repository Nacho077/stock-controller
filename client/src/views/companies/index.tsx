import { Link } from 'react-router-dom'
import Loader from '../../components/loader'
import Footer from '../../components/footer'

import { useGetCompanies } from '../../hooks'

import styles from './companies.module.scss'

const Companies: React.FC = () => {
    const { data, isLoading } = useGetCompanies()

    return (
        <div className={styles.containerMain}>
            <div className={styles.container}>
                <Loader isLoading={isLoading}>
                    {data.map(company => (
                        <Link to={`/company/${company.id}`} key={company.id} className={styles.company}>
                            {company.name}
                        </Link>
                    ))}
                </Loader>
            </div>
            <Footer />
        </div>
    )
}

export default Companies
