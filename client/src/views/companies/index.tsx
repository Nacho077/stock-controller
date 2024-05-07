import { Link } from 'react-router-dom'
import Footer from '../../components/footer'

import { useGetCompanies, useCreateNewCompany, useUpdateCompany, useAppSelector } from '../../hooks'

import styles from './companies.module.scss'
import { ChangeEvent, useState } from 'react'
import { Company, defaultCompany, formButtons, formFields } from './interfaces'

import Form from '../../components/form'
import EditLogo from '../../assets/edit.svg'

const Companies: React.FC = () => {
	useGetCompanies()
    const {companies} = useAppSelector(state => state.reducer)
	const createNewCompany = useCreateNewCompany()
	const updateCompany = useUpdateCompany()
	const [companyForm, setCompanyForm] = useState<Company>(defaultCompany)
	const [isModalVisible, setModalVisible] = useState<boolean>(false)

	const openForm = (companyId: number = 0) => {
		setCompanyForm({
			id: companyId,
			name: ""
		})

		setModalVisible(true)
	}

	const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
		const {name, value} = e.target
		
		setCompanyForm({
			...companyForm,
			[name]: value
		})
	}

	const onSubmit = async () => {
		if (companyForm.id === 0) {
			createNewCompany(companyForm)
		} else {
			updateCompany(companyForm.id, companyForm)
		}

		setModalVisible(false)
	}

	const onReset = () => {
		setCompanyForm(defaultCompany)
		setModalVisible(false)
	}

    return (
        <div className={styles.containerMain}>
            <div className={styles.container}>
				{isModalVisible && <div className={styles.containerModal}>
					<Form 
						title={companyForm.id === 0 ? 'Crear Compania' : 'Modificar Compania'}
						fields={formFields}
						buttons={formButtons}
						formValues={companyForm}
						handleChange={handleChange}
						onSubmit={onSubmit}
						onReset={onReset}
					/>
				</div>}
				<button className={styles.company} id={styles.createButton} onClick={() => openForm()}>Crea una nueva empresa</button>
				{companies.map(company => (
					<span key={company.id} className={styles.company}>
						<Link to={`/company/${company.id}`}>{company.name}</Link>
						<img src={EditLogo} onClick={() => openForm(company.id)} />
					</span>
				))}
            </div>
            <Footer />
        </div>
    )
}

export default Companies
