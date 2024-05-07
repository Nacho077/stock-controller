import { ChangeEvent, useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { useAppSelector, useCreateNewMovement } from '../../hooks'

import UpdatableTableWithFilters from '../../components/updatableTableWithFilters'

import { MovementsFiltersFields, ProductFilters, ProductMovement, getDefaultFilters, getDefaultMovement, movementFormFields, movementHeaders } from './interfaces'

import styles from './movements.module.scss'
import { Link } from 'react-router-dom'
import { useGetProductsByCompanyId, useGetProductsMovementsFiltered, useUpdateMovement } from '../../hooks'

const Movements: React.FC = () => {
    const companyId = parseInt(useParams()["companyId"] || '0', 10)
    const useFilters = useGetProductsMovementsFiltered()
    useGetProductsByCompanyId(companyId)
    const {movements: rows, totalUnits, products} = useAppSelector(state => state.reducer)
    const [filters, setFilters] = useState<ProductFilters>(getDefaultFilters())
    const [movementForm, setMovementForm] = useState<ProductMovement>(getDefaultMovement(companyId))
    const createNewMovement = useCreateNewMovement()
    const updateMovement = useUpdateMovement()
    const [isLoading, setLoading] = useState<boolean>(false)

    const updateMovements = (filters: ProductFilters) => {
        setLoading(true)
        useFilters(companyId, filters)
        .then(() => setLoading(false))
    }

    useEffect(() => updateMovements(filters), [])


    const autoCompleteFields = <T extends ProductFilters | ProductMovement>(state: T, codeToFind: string): T => {
        const movement = rows.find(row => row.code?.toLowerCase() == codeToFind?.toLowerCase())
        const productFounded = products.find(product => product.name?.toLowerCase() == codeToFind?.toLowerCase())

        return {
            ...state,
            brand: movement?.brand || state.brand,
            name: movement?.name || state.name,
            detail: movement?.detail || (state as ProductMovement).detail,
            deposit: movement?.deposit || (state as ProductMovement).deposit,
            observations: movement?.observations || (state as ProductMovement).observations,
            productId: productFounded?.id,
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

    const clearFilters = () => {
      setFilters(getDefaultFilters())
      updateMovements(getDefaultFilters())
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
        if (movementForm.movementId === 0) {
            createNewMovement(companyId, movementForm)
        } else {
            updateMovement(companyId, movementForm.movementId, movementForm)
        }

        setMovementForm(getDefaultMovement(
            companyId,
            movementForm.date,
            movementForm.shippingCode,
            movementForm.code
        ))
    }

    const getAvailableProductsToAutocomplete = () => {
        if (movementForm.code === "" && movementForm.name === "" && movementForm.brand === "") return products
    
        return products.filter(product => 
            (movementForm.code !== "" && product.code.includes(movementForm.code)) ||
            (movementForm.name !== "" && product.name.includes(movementForm.name)) ||
            (movementForm.brand !== "" && product.brand.includes(movementForm.brand))
        )
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
                onSubmit: () => updateMovements(filters),
                onReset: clearFilters,
                refIndex: -1
            }}
            table={{
                headers: movementHeaders,
                rows: rows,
                handleDoubleClick: handleDoubleClick
            }}
            form={{
                title: movementForm.movementId === 0 ? "Nuevo Movimiento" : "Modificar Movimiento",
                fields: movementFormFields(movementForm.movementId !== 0, getAvailableProductsToAutocomplete()),
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
                    <span>{totalUnits}</span>
                </div>
                <Link to={`/company/${companyId}/products`} className={styles.button}>ver productos</Link>
            </div>
        </UpdatableTableWithFilters>
    )
}

export default Movements
