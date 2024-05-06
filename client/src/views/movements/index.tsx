import { ChangeEvent, useState } from 'react'
import { useParams } from 'react-router-dom'
import { useGetProductsMovementsByCompanyId, useAppSelector, useCreateNewMovement } from '../../hooks'

import UpdatableTableWithFilters from '../../components/updatableTableWithFilters'

import { MovementsFiltersFields, ProductFilters, ProductMovement, getDefaultFilters, getDefaultMovement, movementFormFields, movementHeaders } from './interfaces'

import styles from './movements.module.scss'
import { Link } from 'react-router-dom'

const Movements: React.FC = () => {
    const companyId = parseInt(useParams()["companyId"] || '0', 10)
    const { isLoading } = useGetProductsMovementsByCompanyId(companyId)
    const rows = useAppSelector(state => state.reducer.movements)
    const [filters, setFilters] = useState<ProductFilters>(getDefaultFilters())
    const [movementForm, setMovementForm] = useState<ProductMovement>(getDefaultMovement(companyId))
    const createNewMovement = useCreateNewMovement()

    const autoCompleteFields = <T extends ProductFilters | ProductMovement>(state: T, codeToFind: string): T => {
        const movement = rows.find(row => row.code?.toLowerCase() == codeToFind?.toLowerCase())

        return {
            ...state,
            brand: movement?.brand || state.brand,
            name: movement?.name || state.name
        }
    }

    const handleFilters = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target

        let newState = {
            ...filters,
            [name]: value
        }

        if (name == "code") {
            newState = autoCompleteFields(newState, value)
        }

        setFilters(newState)
    }

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target

        let newState = {
            ...movementForm,
            [name]: value
        }

        if (name == "code") {
            newState = autoCompleteFields(newState, value)
        }

        setMovementForm(newState)
    }

    const handleDoubleClick = (movement: ProductMovement) => {
        setMovementForm(movement)
    }

    const resetForm = () => {
        setMovementForm(getDefaultMovement(companyId))
    }

    const handleSubmit = () => {
        createNewMovement({ companyId, newMovement: movementForm })

        setMovementForm(getDefaultMovement(
            companyId,
            movementForm.date,
            movementForm.shippingCode,
            movementForm.code
        ))
    }

    return (
        <UpdatableTableWithFilters
            className={styles.containerMain}
            isLoading={isLoading}
            filters={{
                fields: MovementsFiltersFields,
                buttons: [{
                    title: "Vaciar Filtros",
                    type: "reset"
                }, {
                    title: "Buscar",
                    type: "submit"
                }],
                formValues: filters,
                handleChange: handleFilters,
                onSubmit: () => console.log("Filtrando...."),
                onReset: () => { setFilters(getDefaultFilters()); console.log("mostrar de nuevo las rows") },
                refIndex: -1
            }}
            table={{
                headers: movementHeaders,
                rows: rows,
                handleDoubleClick: handleDoubleClick
            }}
            form={{
                title: "Nuevo Movimiento",
                fields: movementFormFields,
                buttons: [{
                    title: "Limpiar",
                    type: "reset"
                }, {
                    title: "Aceptar",
                    type: "submit"
                }],
                formValues: movementForm,
                handleChange: handleChange,
                onSubmit: handleSubmit,
                onReset: resetForm,
                refIndex: 2
            }}
        >
            <div className={styles.containerInfo}>
                <div>
                    <span>Cant Total</span>
                    <span>{rows.reduce(((acc, row) => acc += Number(row.units)), 0)}
                    </span>
                </div>
                <Link to={`/company/${companyId}/products`} className={styles.button}>ver productos</Link>
            </div>
        </UpdatableTableWithFilters>
    )
}

export default Movements
